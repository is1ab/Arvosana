package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/is1ab/Arvosana/types"
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

	now := time.Now().UTC().Truncate(time.Hour)
	delta := 3 * 24 * time.Hour
	conn := New(db)

	ctx := context.Background()
	err = conn.AddHomework(ctx, AddHomeworkParams{
		Name:     "HW0",
		Semester: types.TimeToSemester(now),
		Deadline: types.NewDatetime(now.Add(delta)),
	})
	if err != nil {
		t.Fatal(err)
	}

	hws, err := conn.GetAllHomeworks(ctx)
	if err != nil {
		t.Fatal(err)
	}

	hw := hws[0]
	deadline := hw.Deadline.Time()
	if deadline.Sub(now) != delta {
		t.Fatalf("incorrect result: expect %v, got %v", delta, deadline.Sub(now))
	}
}
