package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Movies MovieModel
	Users  UserModel
}

// takes pointer to database object and returns new Models
func NewModels(db *sql.DB) Models {
	return Models{
		Movies: MovieModel{
			DB: db,
		},
		Users: UserModel{DB: db},
	}
}
