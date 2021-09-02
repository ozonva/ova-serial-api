package main

import (
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
	server "ova-serial-api/internal/api"
	"ova-serial-api/internal/config"
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
				log.Debug().Msgf("Config '%s' updated: %+v\n", configPath, cfg.GetData())
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
		log.Fatal().Msgf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	api.RegisterOvaSerialServer(s, server.NewSerialAPI())

	if err := s.Serve(listen); err != nil {
		log.Fatal().Msgf("failed to serve: %v", err)
	}

	return nil
}
