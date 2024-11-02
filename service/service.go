package service

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/is1ab/Arvosana/api"
	"github.com/is1ab/Arvosana/service/db"
	"github.com/is1ab/Arvosana/service/logger"
	"github.com/is1ab/Arvosana/web"
	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type Service struct {
	router *echo.Echo
	logger *zap.SugaredLogger
	ctx    context.Context
}

func NewService() (*Service, error) {
	e := echo.New()
	l := logger.NewLogger()
	q, err := db.NewQueries()
	if err != nil {
		return nil, fmt.Errorf("error connecting to DB")
	}

	ctx := context.Background()
	ctx = logger.WithContext(ctx, l)
	ctx = db.WithContext(ctx, q)

	e.Use(mw.CORS())

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	})

	e.Use(mw.RequestLoggerWithConfig(mw.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogMethod: true,
		LogValuesFunc: func(c echo.Context, v mw.RequestLoggerValues) error {
			l := logger.Ctx(c.Request().Context())
			l.Info(
				"request",
				zap.Int("status", v.Status),
				zap.String("method", v.Method),
				zap.String("uri", v.URI),
			)

			return nil
		},
	}))

	webGroup := e.Group("")
	web.RegisterFrontend(webGroup)

	apiGroup := e.Group("/api")
	api.RegisterHomework(apiGroup)
	api.RegisterStudent(apiGroup)
	api.RegisterGrade(apiGroup)

	return &Service{
		router: e,
		logger: l,
		ctx:    ctx,
	}, nil
}

func (s *Service) Start() error {
	err := s.router.Start(":8080")
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("server crashed: %w", err)
	}

	return nil
}

func (s *Service) Shutdown() error {
	timeout := 10 * time.Second

	var err error
	ctx, cancel := context.WithTimeout(s.ctx, timeout)
	defer cancel()

	err2 := s.router.Shutdown(ctx)
	err = errors.Join(err, fmt.Errorf("error shutting down: %w", err2))

	err3 := logger.Close(s.logger)
	err = errors.Join(err, fmt.Errorf("error closing logger: %w", err3))

	return err
}
