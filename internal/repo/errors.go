package repo

import "fmt"

type NotFound struct {
	Operation string
	Table     string
	Id        int64
}

func (m *NotFound) Error() string {
	return fmt.Sprintf("row with id %d not found in table %s while %s", m.Id, m.Table, m.Operation)
}
