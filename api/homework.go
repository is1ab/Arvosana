package api

import (
	"errors"
	"net/http"
	"time"

	"github.com/is1ab/Arvosana/middleware"
	"github.com/is1ab/Arvosana/service/db"
	"github.com/is1ab/Arvosana/service/logger"
	"github.com/is1ab/Arvosana/types"
	"github.com/labstack/echo/v4"
)

func RegisterHomework(e *echo.Group) {
	e.GET("/homework", func(c echo.Context) error {
		ctx := c.Request().Context()
		l := logger.Ctx(ctx)
		q := db.Ctx(ctx)

		hws, err := q.GetAllHomeworks(ctx)
		if err != nil {
			l.Errorln(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return c.JSON(http.StatusOK, hws)
	})

	type GetHomeworkBySemesterRequest struct {
		Semester string `param:"semester"`
	}

	e.GET("/homework/:semester", func(c echo.Context) error {
		ctx := c.Request().Context()
		l := logger.Ctx(ctx)
		q := db.Ctx(ctx)

		var data GetHomeworkBySemesterRequest
		err := c.Bind(&data)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, errors.Unwrap(err))
		}

		sem, err := types.ParseSemester(data.Semester)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		hws, err := q.GetHomeworksFromSemester(ctx, sem)
		if err != nil {
			l.Errorln(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		if len(hws) == 0 {
			return echo.NewHTTPError(http.StatusNotFound)
		}

		return c.JSON(http.StatusOK, hws)
	})

	type PostHomeworkRequest struct {
		Name     string `json:"name"`
		Semester string `json:"semester"`
		BeginAt  int64  `json:"begin_at"`
		EndAt    int64  `json:"end_at"`
	}

	e.POST("/homework", func(c echo.Context) error {
		ctx := c.Request().Context()
		l := logger.Ctx(ctx)
		q := db.Ctx(ctx)

		var data PostHomeworkRequest
		err := c.Bind(&data)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, errors.Unwrap(err))
		}

		sem, err := types.ParseSemester(data.Semester)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		beginAt := time.Unix(data.BeginAt, 0)
		endAt := time.Unix(data.EndAt, 0)

		err = q.AddHomework(ctx, db.AddHomeworkParams{
			Name:     data.Name,
			Semester: sem,
			BeginAt:  types.NewDatetime(beginAt),
			EndAt:    types.NewDatetime(endAt),
		})
		if err != nil {
			l.Errorln(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return c.NoContent(http.StatusCreated)
	}, middleware.Protected)

	type PatchHomeworkRequest struct {
		OldSemester string `param:"old_semester"`
		OldName     string `param:"old_name"`
		NewSemester string `json:"new_semester,omitempty"`
		NewName     string `json:"new_name,omitempty"`
		NewBeginAt  int64  `json:"new_begin_at,omitempty"`
		NewEndAt    int64  `json:"new_end_at,omitempty"`
	}

	e.PATCH("/homework/:old_semester/:old_name", func(c echo.Context) error {
		ctx := c.Request().Context()
		l := logger.Ctx(ctx)
		q := db.Ctx(ctx)

		var data PatchHomeworkRequest
		err := c.Bind(&data)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, errors.Unwrap(err))
		}

		oldSem, err := types.ParseSemester(data.OldSemester)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		newSem, err := types.ParseNullSemester(data.NewSemester)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		newName, err := types.ParseNullString(data.NewName)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		newBeginAt, err := types.ParseNullDateTime(data.NewBeginAt)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		newEndAt, err := types.ParseNullDateTime(data.NewEndAt)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		err = q.UpdateHomework(ctx, db.UpdateHomeworkParams{
			NewName:     newName,
			NewSemester: newSem,
			NewBeginAt:  newBeginAt,
			NewEndAt:    newEndAt,
			OldSemester: oldSem,
			OldName:     data.OldName,
		})
		if err != nil {
			l.Errorln(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return c.NoContent(http.StatusOK)
	}, middleware.Protected)
}
