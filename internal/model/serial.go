package model

import "fmt"

type Serial struct {
	ID      int64  `db:"id"`
	UserID  int64  `db:"user_id"`
	Title   string `db:"title"`
	Genre   string `db:"genre"`
	Year    uint32 `db:"year"`
	Seasons uint32 `db:"seasons"`
}

func (s Serial) String() string {
	return fmt.Sprintf("Serial '%s', %d, %s, %d seasons", s.Title, s.Year, s.Genre, s.Seasons)
}
