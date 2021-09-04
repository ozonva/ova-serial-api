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

	log.Debug().
		Int64("UserId", req.UserId).
		Str("Title", req.Title).
		Str("Genre", req.Genre).
		Uint32("Year", req.Year).
		Uint32("Seasons", req.Seasons).
		Msg("Create serial")

	serial := model.Serial{
		UserID:  req.UserId,
		Title:   req.Title,
		Genre:   req.Genre,
		Year:    req.Year,
		Seasons: req.Seasons,
	}

	id, err := a.repo.AddEntity(serial)
	if err != nil {
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

	log.Debug().
		Int64("Id", req.Id).
		Msg("Get serial")

	serial, err := a.repo.GetEntity(req.Id)
	if err != nil {
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
	log.Debug().
		Uint64("Limit", req.Limit).
		Uint64("Offset", req.Offset).
		Msg("List serials")

	fetched, err := a.repo.ListEntities(req.Limit, req.Offset)
	if err != nil {
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

	log.Debug().
		Int64("Id", req.Id).
		Msg("Remove serial")

	err := a.repo.RemoveEntity(req.Id)
	if err != nil {
		return nil, getErrorText(err)
	}

	return &emptypb.Empty{}, nil
}

func getErrorText(err error) error {
	if errors.Is(err, &repo.NotFound{}) {
		return status.Error(codes.NotFound, "not found")
	}
	return status.Error(codes.Internal, "internal error")
}
