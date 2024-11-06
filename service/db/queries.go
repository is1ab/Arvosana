package db

import (
	"context"
	"database/sql"

	"github.com/is1ab/Arvosana/env"

	_ "github.com/mattn/go-sqlite3"
)

type queriesKey struct{}

func NewQueries() (*Queries, error) {
	conn, err := sql.Open("sqlite3", env.DB)
	if err != nil {
		return nil, err
	}

	_, err = conn.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		return nil, err
	}

	return New(conn), nil
}

func WithContext(ctx context.Context, q *Queries) context.Context {
	return context.WithValue(ctx, queriesKey{}, q)
}

func Ctx(ctx context.Context) *Queries {
	db, ok := ctx.Value(queriesKey{}).(*Queries)
	if !ok {
		panic("no db context")
	}
	return db
}
