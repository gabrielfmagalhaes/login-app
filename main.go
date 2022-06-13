package main

import (
	"context"
	"database/sql"
	"log"
	"login-app/internal/api"
	"login-app/platform/mongodb/dbconfig"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func newDatabaseConnection() *sql.DB {
	db, err := sql.Open(dbconfig.PostgresDriver, dbconfig.DataSourceName)

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func main() {
	e := api.SetupRouter()

	db := newDatabaseConnection()
	defer db.Close()

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
