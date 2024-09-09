package api

import (
	"errors"
	"net/http"

	"github.com/is1ab/Arvosana/service/db"
	"github.com/is1ab/Arvosana/service/logger"
	"github.com/is1ab/Arvosana/types"
	"github.com/labstack/echo/v4"
)

func RegisterStudent(e *echo.Group) {
	e.GET("/student", func(c echo.Context) error {
		ctx := c.Request().Context()
		l := logger.Ctx(ctx)
		q := db.Ctx(ctx)

		students, err := q.GetAllStudents(ctx)
		if err != nil {
			l.Errorln(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return c.JSON(http.StatusOK, students)
	})

	e.GET("/student/:semester", func(c echo.Context) error {
		ctx := c.Request().Context()
		l := logger.Ctx(ctx)
		q := db.Ctx(ctx)

		data := c.Param("semester")
		sem, err := types.ParseSemester(data)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		students, err := q.GetStudentsBySemester(ctx, sem)
		if err != nil {
			l.Errorln(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return c.JSON(http.StatusOK, students)
	})

	type PostStudentRequest struct {
		StudentId string `json:"student_id"`
		Semester  string `json:"semester"`
	}

	e.POST("/student", func(c echo.Context) error {
		ctx := c.Request().Context()
		l := logger.Ctx(ctx)
		q := db.Ctx(ctx)

		var data PostStudentRequest
		err := c.Bind(&data)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, errors.Unwrap(err))
		}

		sem, err := types.ParseSemester(data.Semester)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		err = q.AddStudent(ctx, db.AddStudentParams{
			StudentID: data.StudentId,
			Semester:  sem,
		})
		if err != nil {
			l.Errorln(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return c.NoContent(http.StatusCreated)
	})
}