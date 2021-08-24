package repo

import "ova-serial-api/internal/model"

type Repo interface {
	AddEntities(entities []model.Serial) error
	ListEntities(limit, offset uint64) ([]model.Serial, error)
	GetEntity(entityId int64) (model.Serial, error)
}
