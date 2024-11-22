package sse

import (
	"context"

	s "github.com/r3labs/sse/v2"
)

type sseKey struct{}

type Event = s.Event

const SUBMIT_STREAM string = "submit"

func NewSseServer() *s.Server {
	server := s.New()
	server.AutoReplay = false
	_ = server.CreateStream(SUBMIT_STREAM)

	return server
}

func Close(server *s.Server) {
	server.Close()
}

func WithContext(ctx context.Context, server *s.Server) context.Context {
	return context.WithValue(ctx, sseKey{}, server)
}

func Ctx(ctx context.Context) *s.Server {
	server, ok := ctx.Value(sseKey{}).(*s.Server)
	if !ok {
		panic("no sse server context")
	}
	return server
}
