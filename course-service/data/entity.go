package data

import "github.com/jmoiron/sqlx"

var db *sqlx.DB

func New(dbPool *sqlx.DB) *Models {
	db = dbPool

	return &Models{
		Course: &Course{},
	}
}

type Models struct {
	Course CourseInterfaces
}
