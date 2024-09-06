package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/is1ab/Arvosana/service"
)

func main() {
	service, err := service.NewService()
	if err != nil {
		log.Fatalln(err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		err := service.Start()
		if err != nil {
			log.Fatalln(err)
		}
	}()

	<-ctx.Done()
	err = service.Shutdown()
	if err != nil {
		log.Fatalln(err)
	}
}
