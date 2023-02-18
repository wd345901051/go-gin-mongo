package models

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Mongo = InitMongo()
)

func InitMongo() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://8.130.106.57:27017"))
	if err != nil {
		log.Fatal(err)
		return nil
	}
	db := client.Database("im")
	return db
}
