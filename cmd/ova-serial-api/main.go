package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
	"os"
	server "ova-serial-api/internal/api"
	"ova-serial-api/internal/config"
	"ova-serial-api/internal/repo"
	api "ova-serial-api/pkg/ova-serial-api"
	"time"
)

const (
	configPath         = "config/test_config.yaml"
	confUpdIntervalSec = 10
	grpcPort           = ":82"
)

func main() {
	cfg := config.NewConfig(configPath)
	go func() {
		for {
			err := cfg.Update()
			if err != nil {
				log.Error().Msgf("Error while reading config: %s\n", err)
			} else {
				//log.Debug().Msgf("Config '%s' updated: %+v\n", configPath, cfg.GetData())
			}

			time.Sleep(confUpdIntervalSec * time.Second)
		}
	}()

	if err := startGRPCServer(); err != nil {
		log.Fatal().Msgf("Error while starting server: %v", err)
	}
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

	srv := server.NewSerialAPI(repo.NewSerialRepo(db))
	_, _ = srv.UpdateSerialV1(nil, &api.UpdateSerialRequestV1{
		Serial: &api.SerialV1{
			Id:      21,
			UserId:  1,
			Title:   "DEADBEEF",
			Genre:   "DEADBEEF",
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

	return nil
}
