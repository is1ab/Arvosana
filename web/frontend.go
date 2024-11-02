package web

import (
	"embed"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"
)

//go:embed build/*
var output embed.FS

func RegisterFrontend(e *echo.Group) {
	e.Use(mw.StaticWithConfig(mw.StaticConfig{
		Skipper: func(c echo.Context) bool {
			skipList := []string{"/api"}

			for _, v := range skipList {
				if strings.HasPrefix(c.Request().URL.Path, v) {
					return true
				}
			}

			return false
		},
		Root:       "build",
		Filesystem: http.FS(output),
		HTML5:      true,
	}))
}
