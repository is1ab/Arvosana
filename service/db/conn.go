package db

import (
	"database/sql"

	"github.com/is1ab/Arvosana/env"

	_ "github.com/mattn/go-sqlite3"
)

func NewDb() (*sql.DB, error) {
	return sql.Open("sqlite3", env.DB)
}
