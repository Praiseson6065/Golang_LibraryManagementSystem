package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

func startServer(ctx context.Context, server *http.Server, name string) error {

	go func() {

		startTime := time.Now().UnixMilli()
		log.Printf("[%d] %s started on %s", startTime, name, server.Addr)

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("[%d] %s failed: %v", startTime, name, err)
			cancelContext(ctx)
		}
	}()

	<-ctx.Done()

	shutdownTime := time.Now().UnixMilli()
	log.Printf("[%d] Shutting down %s...", shutdownTime, name)

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return server.Shutdown(shutdownCtx)
}

func cancelContext(ctx context.Context) {
	if cancel, ok := ctx.Value("cancel").(context.CancelFunc); ok {
		cancel()
	}
}
