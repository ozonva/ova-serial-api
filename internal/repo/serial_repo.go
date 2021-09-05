package repo

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"ova-serial-api/internal/model"
)

type serial_repo struct {
	db sqlx.DB
}

func NewSerialRepo(db *sqlx.DB) Repo {
	return &serial_repo{db: *db}
}

func (r *serial_repo) AddEntity(entity model.Serial) (int64, error) {
	var id int64
	query, err := r.db.PrepareNamed(`INSERT INTO serial (user_id,title,genre,year,seasons) VALUES (:user_id,:title,:genre,:year,:seasons) RETURNING id`)

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
	return nil
}

func (r *serial_repo) ListEntities(limit, offset uint64) ([]model.Serial, error) {
	serials := make([]model.Serial, 0, limit)
	err := r.db.Select(&serials, "SELECT id, user_id,title,genre,year,seasons FROM serial ORDER BY id ASC LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return nil, err
	}

	return serials, nil
}

func (r *serial_repo) GetEntity(entityID int64) (*model.Serial, error) {
	serial := model.Serial{}
	err := r.db.Get(&serial, "SELECT id, user_id, title, genre, year, seasons FROM serial WHERE id=$1", entityID)
	if err != nil {
		return nil, getError(err)
	}

	return &serial, nil
}

func (r *serial_repo) RemoveEntity(entityID int64) error {
	res, err := r.db.Exec(`DELETE FROM serial WHERE id=$1`, entityID)

	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return &NotFound{}
	}

	return nil
}

func (r *serial_repo) UpdateEntity(entityID int64, entity model.Serial) error {
	res, err := r.db.NamedExec(`UPDATE serial SET user_id=:user_id, title=:title, genre=:genre, year =:year, seasons=:seasons, where id=:id`,
		map[string]interface{}{
			"id":      entityID,
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
		return &NotFound{}
	}

	return nil
}

func getError(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return &NotFound{}
	}

	return err
}
