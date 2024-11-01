// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: homework.sql

package db

import (
	"context"

	"github.com/is1ab/Arvosana/types"
)

const addHomework = `-- name: AddHomework :exec
INSERT INTO homework (name, semester, begin_at, end_at)
VALUES (?, ?, ?, ?)
`

type AddHomeworkParams struct {
	Name     string         `json:"name"`
	Semester types.Semester `json:"semester"`
	BeginAt  types.Datetime `json:"begin_at"`
	EndAt    types.Datetime `json:"end_at"`
}

func (q *Queries) AddHomework(ctx context.Context, arg AddHomeworkParams) error {
	_, err := q.db.ExecContext(ctx, addHomework,
		arg.Name,
		arg.Semester,
		arg.BeginAt,
		arg.EndAt,
	)
	return err
}

const getAllHomeworks = `-- name: GetAllHomeworks :many
SELECT name, semester, begin_at, end_at FROM homework
`

type GetAllHomeworksRow struct {
	Name     string         `json:"name"`
	Semester types.Semester `json:"semester"`
	BeginAt  types.Datetime `json:"begin_at"`
	EndAt    types.Datetime `json:"end_at"`
}

func (q *Queries) GetAllHomeworks(ctx context.Context) ([]GetAllHomeworksRow, error) {
	rows, err := q.db.QueryContext(ctx, getAllHomeworks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAllHomeworksRow{}
	for rows.Next() {
		var i GetAllHomeworksRow
		if err := rows.Scan(
			&i.Name,
			&i.Semester,
			&i.BeginAt,
			&i.EndAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getHomeworksFromSemester = `-- name: GetHomeworksFromSemester :many
SELECT name, begin_at, end_at FROM homework
WHERE semester = ?
`

type GetHomeworksFromSemesterRow struct {
	Name    string         `json:"name"`
	BeginAt types.Datetime `json:"begin_at"`
	EndAt   types.Datetime `json:"end_at"`
}

func (q *Queries) GetHomeworksFromSemester(ctx context.Context, semester types.Semester) ([]GetHomeworksFromSemesterRow, error) {
	rows, err := q.db.QueryContext(ctx, getHomeworksFromSemester, semester)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetHomeworksFromSemesterRow{}
	for rows.Next() {
		var i GetHomeworksFromSemesterRow
		if err := rows.Scan(&i.Name, &i.BeginAt, &i.EndAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
