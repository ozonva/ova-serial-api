package saver

import (
	"fmt"
	"ova-serial-api/internal/flusher"
	"ova-serial-api/internal/model"
	"time"
)

type Saver interface {
	Save(serial model.Serial)
	Init()
	Close()
}

type saver struct {
	capacity  uint
	flusher   flusher.Flusher
	storage   []model.Serial
	ticker    *time.Ticker
	saveChan  chan model.Serial
	closeChan chan interface{}
}

func NewSaver(
	capacity uint,
	flusher flusher.Flusher,
	timeoutSec uint,
) Saver {
	slice := make([]model.Serial, 0, capacity)

	saveChan := make(chan model.Serial)
	closeChan := make(chan interface{})

	ticker := time.NewTicker(time.Duration(timeoutSec) * time.Second)

	s := saver{
		capacity:  capacity,
		flusher:   flusher,
		storage:   slice,
		ticker:    ticker,
		saveChan:  saveChan,
		closeChan: closeChan,
	}

	return &s
}

func (s *saver) Init() {
	go func() {
		for {
			select {
			case <-s.ticker.C:
				s.flush()

			case v, ok := <-s.saveChan:
				if !ok {
					fmt.Println("Error while reading save chan")
					continue
				}
				s.storage = append(s.storage, v)

			case _, ok := <-s.closeChan:
				if !ok {
					fmt.Println("Error while reading close chan")
					continue
				}

				s.flush()
				close(s.saveChan)
				close(s.closeChan)
				s.ticker.Stop()

				return
			}
		}
	}()
}

func (s *saver) Save(serial model.Serial) {
	s.saveChan <- serial
}

func (s *saver) Close() {
	s.closeChan <- struct{}{}
}

func (s *saver) flush() {
	if len(s.storage) > 0 {
		s.storage = s.flusher.Flush(s.storage)
	}
}
