package api

import (
	"context"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"
	api "ova-serial-api/pkg/ova-serial-api"
)

type OvaSerialAPI struct {
	api.UnimplementedOvaSerialServer
}

func NewSerialAPI() api.OvaSerialServer {
	return &OvaSerialAPI{
		UnimplementedOvaSerialServer: api.UnimplementedOvaSerialServer{},
	}
}

func (a *OvaSerialAPI) CreateSerialV1(ctx context.Context, req *api.CreateSerialRequestV1) (*api.CreateSerialResponseV1, error) {
	log.Info().Msgf("CreateSerial request: %v", req)
	return nil, nil
}

func (a *OvaSerialAPI) GetSerialV1(ctx context.Context, req *api.GetSerialRequestV1) (*api.GetSerialResponseV1, error) {
	log.Info().Msgf("CreateSerial request: %v", req)
	return nil, nil
}

func (a *OvaSerialAPI) ListSerialsV1(ctx context.Context, empty *emptypb.Empty) (*api.ListSerialsResponseV1, error) {
	log.Info().Msgf("CreateSerial request: %v", empty)
	return nil, nil
}

func (a *OvaSerialAPI) RemoveSerialV1(ctx context.Context, req *api.RemoveSerialRequestV1) (*emptypb.Empty, error) {
	log.Info().Msgf("CreateSerial request: %v", req)
	return nil, nil
}
