package repo

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"ova-serial-api/internal/model"
)

type serial_repo struct {
	db sqlx.DB
}

const TABLE_NAME = "serial"

func NewSerialRepo(db *sqlx.DB) Repo {
	return &serial_repo{db: *db}
}

func (r *serial_repo) AddEntity(entity model.Serial) (int64, error) {
	var id int64
	query, err := r.db.PrepareNamed(fmt.Sprintf(
		"INSERT INTO %s (user_id,title,genre,year,seasons) VALUES (:user_id,:title,:genre,:year,:seasons) RETURNING id", TABLE_NAME,
	))

	if err != nil {
		return 0, err
	}

	err = query.Get(&id, entity)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *serial_repo) AddEntities(entities []model.Serial) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	stmt, err := tx.Prepare(fmt.Sprintf(
		"INSERT INTO %s (user_id,title,genre,year,seasons) VALUES ($1, $2, $3, $4, $5)", TABLE_NAME,
	))
	if err != nil {
		return err
	}
	defer stmt.Close()
	for _, serial := range entities {
		_, execErr := stmt.Exec(serial.UserID, serial.Title, serial.Genre, serial.Year, serial.Seasons)
		if execErr != nil {
			return execErr
		}
	}
	return tx.Commit()
}

func (r *serial_repo) ListEntities(limit, offset uint64) ([]model.Serial, error) {
	serials := make([]model.Serial, 0, limit)
	err := r.db.Select(&serials, fmt.Sprintf("SELECT id, user_id,title,genre,year,seasons FROM %s ORDER BY id ASC LIMIT $1 OFFSET $2", TABLE_NAME), limit, offset)
	if err != nil {
		return nil, err
	}

	return serials, nil
}

func (r *serial_repo) GetEntity(entityID int64) (*model.Serial, error) {
	serial := model.Serial{}
	err := r.db.Get(&serial, fmt.Sprintf("SELECT id, user_id, title, genre, year, seasons FROM %s WHERE id=$1", TABLE_NAME), entityID)
	if err != nil {
		return nil, getError(err, "get", entityID)
	}

	return &serial, nil
}

func (r *serial_repo) RemoveEntity(entityID int64) error {
	res, err := r.db.Exec(fmt.Sprintf(`DELETE FROM %s WHERE id=$1`, TABLE_NAME), entityID)

	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return &NotFound{
			Operation: "remove",
			Table:     TABLE_NAME,
			Id:        entityID,
		}
	}

	return nil
}

func (r *serial_repo) UpdateEntity(entity model.Serial) error {
	res, err := r.db.NamedExec(`UPDATE serial SET user_id=:user_id, title=:title, genre=:genre, year =:year, seasons=:seasons WHERE id=:id`,
		map[string]interface{}{
			"id":      entity.ID,
			"user_id": entity.UserID,
			"title":   entity.Title,
			"genre":   entity.Genre,
			"year":    entity.Year,
			"seasons": entity.Seasons,
		})

	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return &NotFound{
			Operation: "update",
			Table:     TABLE_NAME,
			Id:        entity.ID,
		}
	}

	return nil
}

func getError(err error, operation string, id int64) error {
	if errors.Is(err, sql.ErrNoRows) {
		return &NotFound{
			Operation: operation,
			Table:     TABLE_NAME,
			Id:        id,
		}
	}

	return err
}
