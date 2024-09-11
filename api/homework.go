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
}
