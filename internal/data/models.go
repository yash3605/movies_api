package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("Record Not Found")
	ErrEditConflict = errors.New("edit conflict")
)

type Models struct {
	Movies MovieModel
}

func NewModels(db *sql.DB) Models{
	return Models{
		Movies: MovieModel{DB: db},
	}
}
