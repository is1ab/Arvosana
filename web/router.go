package web

import (
	"embed"
	"net/http"

	"github.com/is1ab/Arvosana/service/db"
	"github.com/is1ab/Arvosana/service/logger"
	"github.com/is1ab/Arvosana/web/template/pages"
	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"
)

//go:embed static/*
var static embed.FS

func RegisterRouter(e *echo.Group) {
	e.Group("/static", mw.StaticWithConfig(mw.StaticConfig{
		Browse:     true,
		Root:       "static",
		Filesystem: http.FS(static),
	}))

	e.GET("/homework/latest", func(c echo.Context) error {
		ctx := c.Request().Context()
		l := logger.Ctx(ctx)
		q := db.Ctx(ctx)

		hws, err := q.GetAllHomeworks(ctx)
		if err != nil {
			l.Errorln(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		_hws := make([]pages.Homework, len(hws))
		for i := range _hws {
			_hws[i].Name = hws[i].Name
			_hws[i].BeginAt = hws[i].BeginAt
			_hws[i].EndAt = hws[i].EndAt
		}

		html, err := Render(ctx, pages.Homeworks(_hws))
		if err != nil {
			l.Errorln(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return c.HTML(http.StatusOK, html)
	})

	e.GET("/homework/:semester", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "")
	})

	e.GET("/test", func(c echo.Context) error {
		ctx := c.Request().Context()
		l := logger.Ctx(ctx)

		html, err := Render(ctx, pages.Test())
		if err != nil {
			l.Errorln(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return c.HTML(http.StatusOK, html)
	})
}
