package repo

import (
	"ova-serial-api/internal/model"
)

type Repo interface {
	AddEntity(entity model.Serial) (entityId int64, err error)
	AddEntities(entities []model.Serial) error
	ListEntities(limit, offset uint64) ([]model.Serial, error)
	GetEntity(entityId int64) (*model.Serial, error)
	RemoveEntity(entityId int64) error
	UpdateEntity(entity model.Serial) error
}
