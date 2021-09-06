package api

import (
	"context"
	"errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"ova-serial-api/internal/model"
	"ova-serial-api/internal/repo"
	api "ova-serial-api/pkg/ova-serial-api"
)

type OvaSerialAPI struct {
	api.UnimplementedOvaSerialServer
	repo repo.Repo
}

func NewSerialAPI(repo repo.Repo) api.OvaSerialServer {
	return &OvaSerialAPI{
		repo: repo,
	}
}

func (a *OvaSerialAPI) CreateSerialV1(ctx context.Context, req *api.CreateSerialRequestV1) (*api.CreateSerialResponseV1, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	serial := model.Serial{
		UserID:  req.UserId,
		Title:   req.Title,
		Genre:   req.Genre,
		Year:    req.Year,
		Seasons: req.Seasons,
	}

	log.Debug().Msgf("Create serial: %+v", serial)

	id, err := a.repo.AddEntity(serial)
	if err != nil {
		log.Error().Msgf("Error occurred while creating serial: %+v", err)
		return nil, getErrorText(err)
	}

	return &api.CreateSerialResponseV1{
		Id: id,
	}, nil
}

func (a *OvaSerialAPI) GetSerialV1(ctx context.Context, req *api.GetSerialRequestV1) (*api.GetSerialResponseV1, error) {
	if err := req.Validate(); err != nil {
		return nil, getErrorText(err)
	}

	log.Debug().Msgf("Get serial with Id %d", req.Id)

	serial, err := a.repo.GetEntity(req.Id)
	if err != nil {
		log.Error().Msgf("Error occurred while getting serial with Id %d: %+v", req.Id, err)
		return nil, getErrorText(err)
	}

	return &api.GetSerialResponseV1{
		Serial: &api.SerialV1{
			Id:      serial.ID,
			UserId:  serial.UserID,
			Title:   serial.Title,
			Genre:   serial.Genre,
			Year:    serial.Year,
			Seasons: serial.Seasons,
		},
	}, nil
}

func (a *OvaSerialAPI) ListSerialsV1(ctx context.Context, req *api.ListSerialsRequestV1) (*api.ListSerialsResponseV1, error) {
	log.Debug().Msgf("List serials with limit %d and offset %d", req.Limit, req.Offset)

	fetched, err := a.repo.ListEntities(req.Limit, req.Offset)
	if err != nil {
		log.Error().Msgf("Error occurred while getting serials with limit %d and offset %d: %+v", req.Limit, req.Offset, err)
		return nil, getErrorText(err)
	}

	serials := make([]*api.SerialV1, len(fetched))

	for index := range fetched {
		serials[index] = &api.SerialV1{
			Id:      fetched[index].ID,
			UserId:  fetched[index].UserID,
			Title:   fetched[index].Title,
			Genre:   fetched[index].Genre,
			Year:    fetched[index].Year,
			Seasons: fetched[index].Seasons,
		}
	}

	return &api.ListSerialsResponseV1{
		Serials: serials,
	}, nil
}

func (a *OvaSerialAPI) RemoveSerialV1(ctx context.Context, req *api.RemoveSerialRequestV1) (*emptypb.Empty, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Debug().Msgf("Remove serial with Id: %d", req.Id)

	err := a.repo.RemoveEntity(req.Id)
	if err != nil {
		log.Error().Msgf("Error occurred while removing serial with Id %d: %+v", req.Id, err)
		return nil, getErrorText(err)
	}

	return &emptypb.Empty{}, nil
}

func (a *OvaSerialAPI) UpdateSerialV1(ctx context.Context, req *api.UpdateSerialRequestV1) (*emptypb.Empty, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	serial := model.Serial{
		ID:      req.Serial.Id,
		UserID:  req.Serial.UserId,
		Title:   req.Serial.Title,
		Genre:   req.Serial.Genre,
		Year:    req.Serial.Year,
		Seasons: req.Serial.Seasons,
	}

	log.Debug().Msgf("Update serial: %+v", serial)

	err := a.repo.UpdateEntity(serial)
	if err != nil {
		return nil, getErrorText(err)
	}

	return &emptypb.Empty{}, nil
}

func getErrorText(err error) error {
	if errors.Is(err, &repo.NotFound{}) {
		return status.Error(codes.NotFound, "not found")
	}
	return status.Errorf(codes.Internal, "internal error: %+v", err)
}
