package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

func startServer(ctx context.Context, server *http.Server, name string) error {

	go func() {
		log.Printf("%s started on %s", name, server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("%s failed: %v", name, err)
			cancelContext(ctx)
		}
	}()

	<-ctx.Done()

	log.Printf("Shutting down %s...", name)
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return server.Shutdown(shutdownCtx)
}

func cancelContext(ctx context.Context) {
	if cancel, ok := ctx.Value("cancel").(context.CancelFunc); ok {
		cancel()
	}
}
