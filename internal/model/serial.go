package model

import "fmt"

type Serial struct {
	ID      uint64
	UserID  uint64
	Title   string
	Genre   string
	Year    uint32
	Seasons uint32
}

func (s Serial) String() string {
	return fmt.Sprintf("Serial '%s', %d, %s, %d seasons", s.Title, s.Year, s.Genre, s.Seasons)
}
