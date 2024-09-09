package middleware

import (
	"github.com/is1ab/Arvosana/env"
	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"
)

var Protected echo.MiddlewareFunc = mw.KeyAuth(func(auth string, c echo.Context) (bool, error) {
	return auth == env.SECRET_KEY, nil
})
