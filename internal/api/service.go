package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/opentracing/opentracing-go"
	opentracingLog "github.com/opentracing/opentracing-go/log"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	kafka_client "ova-serial-api/internal/kafka"
	"ova-serial-api/internal/model"
	"ova-serial-api/internal/repo"
	"ova-serial-api/internal/utils"
	api "ova-serial-api/pkg/ova-serial-api"
)

type OvaSerialAPI struct {
	api.UnimplementedOvaSerialServer
	repo        repo.Repo
	kafkaClient kafka_client.Client
	metrics     Metrics
}

const BATCH_SIZE = 2

func NewSerialAPI(repo repo.Repo, kafkaClient kafka_client.Client) api.OvaSerialServer {
	return &OvaSerialAPI{
		repo:        repo,
		kafkaClient: kafkaClient,
		metrics:     newApiMetrics(),
	}
}

func (s *OvaSerialAPI) sendKafkaCUDEvent(detail string) error {
	return s.kafkaClient.SendMessage(detail)
}

func (s *OvaSerialAPI) sendKafkaCreateEvent(detail string) error {
	return s.sendKafkaCUDEvent(detail)
}

func (s *OvaSerialAPI) sendKafkaUpdateEvent(detail string) error {
	return s.sendKafkaCUDEvent(detail)
}

func (s *OvaSerialAPI) sendKafkaRemoveEvent(detail string) error {
	return s.sendKafkaCUDEvent(detail)
}

func (a *OvaSerialAPI) CreateSerialV1(ctx context.Context, req *api.CreateSerialRequestV1) (res *api.CreateSerialResponseV1, err error) {
	defer func() {
		if err != nil {
			a.metrics.incFailCreateSerialCounter()
		} else {
			a.metrics.incSuccessCreateSerialCounter()
		}
	}()

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

	logMessage := fmt.Sprintf("Create serial: %+v", serial)

	if sendError := a.sendKafkaCreateEvent(logMessage); sendError != nil {
		log.Error().Msgf("Error occurred while sending create event to kafka, error: %s", sendError)
	}

	log.Info().Msg(logMessage)

	id, err := a.repo.AddEntity(serial)
	if err != nil {
		log.Error().Msgf("Error occurred while creating serial: %+v", err)
		return nil, getErrorText(err)
	}

	return &api.CreateSerialResponseV1{
		Id: id,
	}, nil
}

func (a *OvaSerialAPI) MultiCreateSerialV1(ctx context.Context, req *api.MultiCreateSerialRequestV1) (*emptypb.Empty, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Info().Msgf("Multi create %d serials", len(req.Serials))

	serials := make([]model.Serial, 0, len(req.Serials))
	for _, s := range req.Serials {
		serials = append(serials, model.Serial{
			ID:      s.Id,
			UserID:  s.UserId,
			Title:   s.Title,
			Genre:   s.Genre,
			Year:    s.Year,
			Seasons: s.Seasons,
		})
	}

	span := opentracing.StartSpan("MultiCreateSerials")
	span.LogFields(opentracingLog.Int("Total serials count", len(serials)))
	defer span.Finish()

	for _, serialsSlice := range utils.SplitSerialSlice(serials, BATCH_SIZE) {
		if err := a.batchCreate(span, serialsSlice); err != nil {
			return nil, err
		}
	}

	return &emptypb.Empty{}, nil
}

func (a *OvaSerialAPI) batchCreate(parentSpan opentracing.Span, serialsSlice []model.Serial) error {
	span := opentracing.StartSpan("MultiCreateSerialsBatch", opentracing.ChildOf(parentSpan.Context()))
	span.LogFields(opentracingLog.Int("Serials count", len(serialsSlice)))
	defer span.Finish()

	if err := a.repo.AddEntities(serialsSlice); err != nil {
		log.Error().Msgf("Error occurred while creating serials: %+v", err)
		return err
	}
	return nil
}

func (a *OvaSerialAPI) GetSerialV1(ctx context.Context, req *api.GetSerialRequestV1) (*api.GetSerialResponseV1, error) {
	if err := req.Validate(); err != nil {
		return nil, getErrorText(err)
	}

	log.Info().Msgf("Get serial with Id %d", req.Id)

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
	log.Info().Msgf("List serials with limit %d and offset %d", req.Limit, req.Offset)

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

func (a *OvaSerialAPI) RemoveSerialV1(ctx context.Context, req *api.RemoveSerialRequestV1) (e *emptypb.Empty, err error) {
	defer func() {
		if err != nil {
			a.metrics.incFailRemoveSerialCounter()
		} else {
			a.metrics.incSuccessRemoveSerialCounter()
		}
	}()

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	logMessage := fmt.Sprintf("Remove serial with Id: %d", req.Id)
	if err = a.sendKafkaRemoveEvent(logMessage); err != nil {
		log.Error().Msgf("Error sending remove event to kafka, error: %s", err)
	}

	log.Info().Msg(logMessage)

	err = a.repo.RemoveEntity(req.Id)
	if err != nil {
		log.Error().Msgf("Error occurred while removing serial with Id %d: %+v", req.Id, err)
		return nil, getErrorText(err)
	}

	return &emptypb.Empty{}, nil
}

func (a *OvaSerialAPI) UpdateSerialV1(ctx context.Context, req *api.UpdateSerialRequestV1) (e *emptypb.Empty, err error) {
	defer func() {
		if err != nil {
			a.metrics.incFailUpdateSerialCounter()
		} else {
			a.metrics.incSuccessUpdateSerialCounter()
		}
	}()

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

	logMessage := fmt.Sprintf("Update serial: %+v", serial)

	if err = a.sendKafkaUpdateEvent(logMessage); err != nil {
		log.Error().Msgf("Error sending update event to kafka, error: %s", err)
	}

	log.Info().Msg(logMessage)

	err = a.repo.UpdateEntity(serial)
	if err != nil {
		log.Error().Msgf("Error occurred while updating serial with Id %d: %+v", req.Serial.Id, err)
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
