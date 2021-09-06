package main

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
	"google.golang.org/grpc"
	"io"
	"net"
	"net/http"
	"os"
	server "ova-serial-api/internal/api"
	"ova-serial-api/internal/config"
	kafka_client "ova-serial-api/internal/kafka"
	"ova-serial-api/internal/repo"
	api "ova-serial-api/pkg/ova-serial-api"
	"time"
)

const (
	configPath         = "config/test_config.yaml"
	confUpdIntervalSec = 10
	grpcPort           = ":82"
	prometheusPort     = ":8081"
	serviceName        = "ova-serial-api"
	kafkaTopic         = "serial-CUD-events"
)

func main() {
	cfg := config.NewConfig(configPath)
	go func() {
		for {
			err := cfg.Update()
			if err != nil {
				log.Error().Msgf("Error while reading config: %s\n", err)
			}
			time.Sleep(confUpdIntervalSec * time.Second)
		}
	}()

	go runPrometheusMetrics()

	tracer, closer := initTracer()
	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()

	if err := startGRPCServer(); err != nil {
		log.Fatal().Msgf("Error while starting server: %v", err)
	}
}

func runPrometheusMetrics() {
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(prometheusPort, nil); err != nil {
		log.Fatal().Msgf("Failed to start listen to metric requests, error %s", err)
	}
}

func initTracer() (opentracing.Tracer, io.Closer) {
	cfg := jaegercfg.Configuration{
		ServiceName: serviceName,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}

	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory

	tracer, closer, err := cfg.NewTracer(
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)
	if err != nil {
		log.Fatal().Msgf("Failed to create tracer, %s", err)
	}
	return tracer, closer
}

func startGRPCServer() error {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatal().Msgf("Failed to listen: %v", err)
		return err
	}

	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal().Msgf("Error while loading config", err)
		return err
	}

	db, err := sqlx.Connect(os.Getenv("DB_DRIVER"), os.Getenv("DB_STRING"))
	if err != nil {
		log.Fatal().Msgf("Error while establishing sql connection", err)
		return err
	}

	kafkaClient := kafka_client.NewKafkaClient()
	kafkaDsn := fmt.Sprintf("%s:%s", os.Getenv("KAFKA_HOST"), os.Getenv("KAFKA_PORT"))
	if kafkaConnErr := kafkaClient.Connect(context.Background(), kafkaDsn, kafkaTopic, 0); kafkaConnErr != nil {
		log.Fatal().Msgf("Error connecting to kafka, %s", kafkaConnErr)
	}

	srv := server.NewSerialAPI(repo.NewSerialRepo(db), kafkaClient)
	s := grpc.NewServer()

	api.RegisterOvaSerialServer(s, srv)

	if err := s.Serve(listen); err != nil {
		log.Fatal().Msgf("Failed to serve: %v", err)
		return err
	}

	return nil
}
