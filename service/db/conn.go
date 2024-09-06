package db

import (
	"context"
	"database/sql"

	"github.com/is1ab/Arvosana/env"

	_ "github.com/mattn/go-sqlite3"
)

type dbKey struct{}

func NewDb() (*sql.DB, error) {
	return sql.Open("sqlite3", env.DB)
}

func WithContext(ctx context.Context, db *sql.DB) context.Context {
	return context.WithValue(ctx, dbKey{}, db)
}

func Ctx(ctx context.Context) *sql.DB {
	db, ok := ctx.Value(dbKey{}).(*sql.DB)
	if !ok {
		panic("no db context")
	}
	return db
}
