package main

import (
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"net"
	"os"
	server "ova-serial-api/internal/api"
	"ova-serial-api/internal/config"
	api "ova-serial-api/pkg/ova-serial-api"
	"time"
)

const (
	configPath         = "config/test_config.json"
	confUpdIntervalSec = 10
	grpcPort           = ":82"
)

func main() {
	output := zerolog.ConsoleWriter{Out: os.Stdout}
	log := zerolog.New(output).With().Timestamp().Logger()

	var cfg config.Config
	go func() {
		for {
			err := config.UpdateConfig(configPath, &cfg)
			if err != nil {
				log.Error().Msgf("Error while reading config: %s\n", nil)
			} else {
				log.Debug().Msgf("Config '%s' updated: %+v\n", configPath, cfg)
			}

			time.Sleep(confUpdIntervalSec * time.Second)
		}
	}()

	if err := run(&log); err != nil {
		log.Fatal().Msgf("Error while starting server: %v", err)
	}
}

func run(logger *zerolog.Logger) error {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		logger.Fatal().Msgf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	api.RegisterOvaSerialV1Server(s, server.NewSerialAPI(logger))

	if err := s.Serve(listen); err != nil {
		logger.Fatal().Msgf("failed to serve: %v", err)
	}

	return nil
}
