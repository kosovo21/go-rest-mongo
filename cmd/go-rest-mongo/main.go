package main

import (
	"context"
	"time"

	"github.com/kosovo21/go-rest-mongo/internal/data/mongodb"
	"github.com/kosovo21/go-rest-mongo/internal/rest"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		DisableColors: true,
	})
	logrus.Info("server starting...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		logrus.Fatal(err)
		return
	}

	userDao := mongodb.New(mongoClient, "go-rest-mongo")

	server := rest.New(userDao)
	logrus.Fatal(server.ListenAndServe())
	logrus.Info("server started...")
}
