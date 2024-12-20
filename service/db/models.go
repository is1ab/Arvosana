// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"github.com/is1ab/Arvosana/types"
)

type Grade struct {
	ID          int64          `json:"id"`
	StudentID   int64          `json:"student_id"`
	HomeworkID  int64          `json:"homework_id"`
	SubmittedAt types.Datetime `json:"submitted_at"`
	Grade       float64        `json:"grade"`
}

type Homework struct {
	ID       int64          `json:"id"`
	Name     string         `json:"name"`
	Semester types.Semester `json:"semester"`
	BeginAt  types.Datetime `json:"begin_at"`
	EndAt    types.Datetime `json:"end_at"`
}

type Student struct {
	ID        int64          `json:"id"`
	StudentID string         `json:"student_id"`
	Semester  types.Semester `json:"semester"`
}
