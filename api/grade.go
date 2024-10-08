package api

import (
	"bytes"
	"encoding/csv"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/is1ab/Arvosana/middleware"
	"github.com/is1ab/Arvosana/service/db"
	"github.com/is1ab/Arvosana/service/logger"
	"github.com/is1ab/Arvosana/types"
	"github.com/labstack/echo/v4"
)

func RegisterGrade(e *echo.Group) {
	type GetGradeInfoRequest struct {
		Semester     string `param:"semester"`
		HomeworkName string `param:"homework_name"`
	}

	e.GET("/grade/:semester/:homework_name", func(c echo.Context) error {
		ctx := c.Request().Context()
		l := logger.Ctx(ctx)
		q := db.Ctx(ctx)

		exportCsv := c.QueryParam("csv") == "1"

		var data GetGradeInfoRequest
		err := c.Bind(&data)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, errors.Unwrap(err))
		}

		sem, err := types.ParseSemester(data.Semester)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		if data.HomeworkName == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "homework_name required")
		}

		info, err := q.GetGradeInfo(ctx, db.GetGradeInfoParams{
			Semester: sem,
			Name:     data.HomeworkName,
		})
		if err != nil {
			l.Errorln(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		if exportCsv {
			var buf bytes.Buffer
			writer := csv.NewWriter(&buf)

			err := writer.Write([]string{"Student ID", "Grade"})
			if err != nil {
				l.Errorln(err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}

			for _, v := range info {
				gradeStr := strconv.FormatFloat(v.Grade, 'f', 2, 64)
				err := writer.Write([]string{v.StudentID, gradeStr})
				if err != nil {
					l.Errorln(err)
					return echo.NewHTTPError(http.StatusInternalServerError)
				}
			}

			writer.Flush()

			return c.Blob(http.StatusOK, "text/csv", buf.Bytes())
		}

		return c.JSON(http.StatusOK, info)
	})

	type PostSubmitRequest struct {
		StudentId    string   `json:"student_id"`
		HomeworkName string   `json:"homework_name"`
		Semester     string   `json:"semester"`
		SubmittedAt  int64    `json:"submitted_at"`
		Grade        *float64 `json:"grade"` // need to separate actual 0 from null values
	}

	e.POST("/submit", func(c echo.Context) error {
		ctx := c.Request().Context()
		l := logger.Ctx(ctx)
		q := db.Ctx(ctx)

		var data PostSubmitRequest
		err := c.Bind(&data)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, errors.Unwrap(err))
		}

		if data.StudentId == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "student_id required")
		}

		if data.HomeworkName == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "homework_name required")
		}

		sem, err := types.ParseSemester(data.Semester)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		submittedAt := time.Unix(data.SubmittedAt, 0)

		if data.Grade == nil {
			return echo.NewHTTPError(http.StatusBadRequest, "grade required")
		}

		info, err := q.GetSubmitInfo(ctx, db.GetSubmitInfoParams{
			StudentID: data.StudentId,
			Name:      data.HomeworkName,
			Semester:  sem,
		})
		if err != nil {
			l.Errorln(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		if submittedAt.Before(info.BeginAt.Time()) {
			return echo.NewHTTPError(http.StatusForbidden, "not yet open")
		}

		if submittedAt.After(info.EndAt.Time()) {
			return echo.NewHTTPError(http.StatusForbidden, "deadline exceeded")
		}

		err = q.SubmitGrade(ctx, db.SubmitGradeParams{
			StudentID:   info.StudentID,
			HomeworkID:  info.HomeworkID,
			SubmittedAt: types.NewDatetime(submittedAt),
			Grade:       *data.Grade,
		})
		if err != nil {
			l.Errorln(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return c.NoContent(http.StatusCreated)
	}, middleware.Protected)
}
