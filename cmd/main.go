package main

import (
	_ "LibManMicroServ/config"
	"LibManMicroServ/events"
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
		cancel()
	}()
	defer cancel()

	var eventBus = events.NewEventBus()
	var g errgroup.Group

	gin.SetMode(gin.ReleaseMode)

	g.Go(func() error {
		return startServer(ctx, APIServer(), "ApiServer")
	})

	g.Go(func() error {
		return startServer(ctx, AuthServer(eventBus), "AuthServer")
	})

	g.Go(func() error {
		return startServer(ctx, UserLendingServer(), "UserLendingServer")
	})
	g.Go(func() error {
		return startServer(ctx, UserReviewServer(), "UserReviewServer")
	})

	g.Go(func() error {
		return startServer(ctx, PaymentsServer(), "PaymentsServer")
	})
	g.Go(func() error {
		return startServer(ctx, AdminBooksServer(), "AdminBooksServer")
	})

	g.Go(func() error {
		return startServer(ctx, AdminLendingServer(), "AdminLendingServer")
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}

}
