package web

import (
	"embed"
	"net/http"

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
