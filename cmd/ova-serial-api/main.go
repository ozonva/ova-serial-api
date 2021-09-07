package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
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
	httpPort           = ":8085"
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

	go startGRPCServer(grpcPort)

	startHttpServer(context.Background(), grpcPort, httpPort)
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

func startHttpServer(ctx context.Context, grpcPort, httpPort string) error {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := api.RegisterOvaSerialHandlerFromEndpoint(ctx, mux, "localhost"+grpcPort, opts); err != nil {
		log.Fatal().Msgf("Failed to start HTTP gateway: %v", err)
	}

	srv := &http.Server{
		Addr:    httpPort,
		Handler: mux,
	}

	return srv.ListenAndServe()
}

func startGRPCServer(grpcPort string) error {
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

	srv := server.NewSerialAPI(repo.NewSerialRepo(db), kafkaClient, server.NewApiMetrics())
	_, _ = srv.MultiCreateSerialV1(context.TODO(), &api.MultiCreateSerialRequestV1{
		Serials: []*api.SerialV1{
			{
				Id:      11,
				UserId:  11,
				Title:   "11abc",
				Genre:   "DEADBEEF",
				Year:    2011,
				Seasons: 11,
			},
			{
				Id:      12,
				UserId:  12,
				Title:   "12abc",
				Genre:   "DEADBEEF",
				Year:    2012,
				Seasons: 12,
			},
			{
				Id:      13,
				UserId:  13,
				Title:   "13abc",
				Genre:   "DEADBEEF",
				Year:    2013,
				Seasons: 13,
			},
		},
	})

	srv.CreateSerialV1(nil, &api.CreateSerialRequestV1{
		UserId:  0,
		Title:   "",
		Genre:   "",
		Year:    0,
		Seasons: 0,
	})
	_, _ = srv.UpdateSerialV1(nil, &api.UpdateSerialRequestV1{
		Serial: &api.SerialV1{
			Id:      79,
			UserId:  1,
			Title:   "test123",
			Genre:   "test123",
			Year:    2021,
			Seasons: 10,
		},
	})
	s := grpc.NewServer()

	api.RegisterOvaSerialServer(s, srv)

	if err := s.Serve(listen); err != nil {
		log.Fatal().Msgf("Failed to serve: %v", err)
		return err
	}
	s.Serve(listen)

	return nil
}
