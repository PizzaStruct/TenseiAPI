package mongodb

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func Connect() error {
	opts := options.Client().ApplyURI(os.Getenv("MONGO"))
	conn, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return err
	}
	if err := conn.Ping(context.Background(), nil); err != nil {
		return err
	}
	db = conn.Database(os.Getenv("DB"))
	return nil
}

func GetClient() *mongo.Database {
	return db
}
