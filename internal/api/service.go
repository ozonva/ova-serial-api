package api

import (
	"context"
	"github.com/rs/zerolog"
	"google.golang.org/protobuf/types/known/emptypb"
	api "ova-serial-api/pkg/ova-serial-api"
)

type OvaSerialAPI struct {
	api.UnimplementedOvaSerialV1Server
	logger *zerolog.Logger
}

func NewSerialAPI(logger *zerolog.Logger) api.OvaSerialV1Server {
	return &OvaSerialAPI{
		UnimplementedOvaSerialV1Server: api.UnimplementedOvaSerialV1Server{},
		logger:                         logger,
	}
}

func (a *OvaSerialAPI) CreateSerialV1(ctx context.Context, req *api.CreateSerialRequestV1) (*api.CreateSerialResponseV1, error) {
	a.logger.Info().Msgf("CreateSerial request: %v", req)
	return a.UnimplementedOvaSerialV1Server.CreateSerial(ctx, req)
}

func (a *OvaSerialAPI) GetSerialV1(ctx context.Context, req *api.GetSerialRequestV1) (*api.GetSerialResponseV1, error) {
	a.logger.Info().Msgf("CreateSerial request: %v", req)
	return a.UnimplementedOvaSerialV1Server.GetSerial(ctx, req)
}

func (a *OvaSerialAPI) ListSerialsV1(ctx context.Context, empty *emptypb.Empty) (*api.ListSerialsResponseV1, error) {
	a.logger.Info().Msgf("CreateSerial request: %v", empty)
	return a.UnimplementedOvaSerialV1Server.ListSerials(ctx, empty)
}

func (a *OvaSerialAPI) RemoveSerialV1(ctx context.Context, req *api.RemoveSerialRequestV1) (*emptypb.Empty, error) {
	a.logger.Info().Msgf("CreateSerial request: %v", req)
	return a.UnimplementedOvaSerialV1Server.RemoveSerial(ctx, req)
}
