package flusher

import (
	"ova-serial-api/internal/model"
	"ova-serial-api/internal/repo"
	"ova-serial-api/internal/utils"
)

type Flusher interface {
	Flush(entities []model.Serial) []model.Serial
}

func NewFlusher(
	chunkSize uint,
	serialRepo repo.Repo,
) Flusher {
	return &flusher{
		chunkSize:  chunkSize,
		serialRepo: serialRepo,
	}
}

type flusher struct {
	chunkSize  uint
	serialRepo repo.Repo
}

func (f flusher) Flush(serials []model.Serial) []model.Serial {
	var notSaved []model.Serial
	for _, chunk := range utils.SplitSerialSlice(serials, f.chunkSize) {
		err := f.serialRepo.AddEntities(chunk)
		if err != nil {
			notSaved = append(notSaved, chunk...)
		}
	}
	return notSaved
}
