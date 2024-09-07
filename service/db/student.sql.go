// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: student.sql

package db

import (
	"context"

	"github.com/is1ab/Arvosana/types"
)

const addStudent = `-- name: AddStudent :exec
INSERT INTO student (id, semester)
VALUES (?, ?)
`

type AddStudentParams struct {
	ID       int64          `json:"id"`
	Semester types.Semester `json:"semester"`
}

func (q *Queries) AddStudent(ctx context.Context, arg AddStudentParams) error {
	_, err := q.db.ExecContext(ctx, addStudent, arg.ID, arg.Semester)
	return err
}
