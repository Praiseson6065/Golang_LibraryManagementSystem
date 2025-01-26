package main

import (
	_ "LibManMicroServ/config"
	"context"
	"os"
	"os/signal"
	"syscall"

	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan
		cancel() // Cancel the context on receiving a signal
	}()
	defer cancel()
	var g errgroup.Group

	gin.SetMode(gin.ReleaseMode)

	g.Go(func() error {
		return startServer(ctx, AuthServer(), "AuthServer")
	})

	g.Go(func() error {
		return startServer(ctx, PaymentsServer(), "PaymentsServer")
	})
	g.Go(func() error {
		return startServer(ctx, APIServer(), "ApiServer")
	})
	g.Go(func() error {
		return startServer(ctx, AdminBooksServer(), "AdminBooksServer")
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}

}
