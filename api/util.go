package api

import (
	"github.com/is1ab/Arvosana/service/logger"
	"github.com/is1ab/Arvosana/service/sse"
	"github.com/labstack/echo/v4"
)

func RegisterUtil(e *echo.Group) {
	e.GET("/sse", func(c echo.Context) error {
		ctx := c.Request().Context()
		l := logger.Ctx(ctx)
		s := sse.Ctx(ctx)

		go func() {
			<-c.Request().Context().Done()
			l.Debugf("client disconnected: %s", c.RealIP())
		}()

		s.ServeHTTP(c.Response(), c.Request())
		return nil
	})
}
