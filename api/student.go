package api

import (
	"errors"
	"net/http"

	"github.com/is1ab/Arvosana/middleware"
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

	type GetStudentInfoRequest struct {
		Semester  string `param:"semester"`
		StudentId string `param:"student_id"`
	}

	e.GET("/student/:semester/:student_id", func(c echo.Context) error {
		ctx := c.Request().Context()
		l := logger.Ctx(ctx)
		q := db.Ctx(ctx)

		var data GetStudentInfoRequest
		err := c.Bind(&data)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, errors.Unwrap(err))
		}

		sem, err := types.ParseSemester(data.Semester)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		if data.StudentId == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "student_id required")
		}

		info, err := q.GetStudentInfo(ctx, db.GetStudentInfoParams{
			Semester:  sem,
			StudentID: data.StudentId,
		})
		if err != nil {
			l.Errorln(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, info)
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

		if data.StudentId == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "student_id required")
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
	}, middleware.Protected)

	type UpdateStudentRequest struct {
		OldSemester  string `param:"old_semester"`
		OldStudentId string `param:"old_student_id"`
		NewSemester  string `json:"new_semester"`
		NewStudentId string `json:"new_student_id"`
	}

	e.PUT("/student/:old_semester/:old_student_id", func(c echo.Context) error {
		ctx := c.Request().Context()
		l := logger.Ctx(ctx)
		q := db.Ctx(ctx)

		var data UpdateStudentRequest
		err := c.Bind(&data)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, errors.Unwrap(err))
		}

		old_sem, err := types.ParseSemester(data.OldSemester)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		if data.OldStudentId == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "old_student_id required")
		}

		new_sem, err := types.ParseSemester(data.NewSemester)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		if data.NewStudentId == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "new_student_id required")
		}

		err = q.UpdateStudent(ctx, db.UpdateStudentParams{
			OldSemester:  old_sem,
			OldStudendID: data.OldStudentId,
			NewSemester:  new_sem,
			NewStudentID: data.NewStudentId,
		})
		if err != nil {
			l.Errorln(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.NoContent(http.StatusCreated)
	}, middleware.Protected)

	type DeleteStudentRequest struct {
		Semester  string `param:"semester"`
		StudentId string `param:"student_id"`
	}

	e.DELETE("/student/:semester/:student_id", func(c echo.Context) error {
		ctx := c.Request().Context()
		l := logger.Ctx(ctx)
		q := db.Ctx(ctx)

		var data DeleteStudentRequest
		err := c.Bind(&data)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, errors.Unwrap(err))
		}

		sem, err := types.ParseSemester(data.Semester)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		if data.StudentId == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "student_id required")
		}

		err = q.DeleteStudent(ctx, db.DeleteStudentParams{
			Semester:  sem,
			StudentID: data.StudentId,
		})
		if err != nil {
			l.Errorln(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.NoContent(http.StatusOK)
	})
}
