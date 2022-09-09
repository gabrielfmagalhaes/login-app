package main

import (
	"context"
	"log"
	"login-app/internal/api/web"
	"login-app/internal/core/repository"
	"login-app/internal/core/services"
	"login-app/platform/logger"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func newDatabaseConnection(ctx context.Context) *mongo.Client {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://user:pass@localhost:27017"))

	if err != nil {
		panic(err)
	}

	return client
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	logger := logger.NewZapLogger()

	db := newDatabaseConnection(ctx)
	defer func() {
		if err := db.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	usersCollection := db.Database("user").Collection("users")
	usersRepository := repository.NewRepository(usersCollection)

	userService := services.NewService(logger, usersRepository)

	e := web.SetupRouter(userService)

	go func() {
		if err := e.Start(":8000"); err != nil && err != http.ErrServerClosed {
			log.Fatal()
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	if err := e.Shutdown(ctx); err != nil {
	}

}
