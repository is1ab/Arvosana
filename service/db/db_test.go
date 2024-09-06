package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

func TestAddHomework(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		t.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://../../db/migrations", "sqlite3", driver)
	if err != nil {
		t.Fatal(err)
	}

	err = m.Up()
	if err != nil {
		t.Fatal(err)
	}

	conn := New(db)

	ctx := context.Background()
	err = conn.AddHomework(ctx, AddHomeworkParams{
		Name:     "HW0",
		Deadline: time.Now().Add(3 * 24 * time.Hour),
	})
	if err != nil {
		t.Fatal(err)
	}

	hws, err := conn.GetAllHomeworks(ctx)
	if err != nil {
		t.Fatal(err)
	}

	for _, hw := range hws {
		fmt.Printf("%+v\n", hw)
	}
}
