package saver

import (
	"errors"
	"ova-serial-api/internal/flusher"
	"ova-serial-api/internal/model"
	"sync"
	"time"
)

type Saver interface {
	Save(serial model.Serial) error
	Close()
}

type saver struct {
	capacity uint
	flusher  flusher.Flusher
	storage  []model.Serial
	mutex    sync.Mutex
}

func NewSaver(
	capacity uint,
	flusher flusher.Flusher,
	timeoutSec uint,
) Saver {
	slice := make([]model.Serial, 0, capacity)

	s := saver{
		capacity: capacity,
		flusher:  flusher,
		storage:  slice,
	}

	go func() {
		for {
			time.Sleep(time.Duration(timeoutSec) * time.Second)
			s.flush()
		}
	}()

	return &s
}

func (s *saver) Save(serial model.Serial) error {
	if len(s.storage) >= int(s.capacity) {
		s.flush()

		if len(s.storage) >= int(s.capacity) {
			return errors.New("no capacity in storage")
		}
	}
	s.storage = append(s.storage, serial)
	return nil
}

func (s *saver) Close() {
	s.flush()
}

func (s *saver) flush() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if len(s.storage) > 0 {
		s.storage = s.flusher.Flush(s.storage)
	}
}
