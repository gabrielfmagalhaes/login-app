package main

import (
	"context"
	"login-app/internal/api"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	e := api.SetupRouter()

	go func() {
		if err := e.Start(":8000"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
