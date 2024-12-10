package database

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

// ket noi database

func ConnectDatabase(nameCollection string) *mongo.Collection {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil
	}
	database := os.Getenv("URL")
	namedatabse := os.Getenv("NAMEDATABASE")
	clientOptions := options.Client().ApplyURI(database)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil

	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	println("Connected to MongoDB!")

	return client.Database(namedatabse).Collection(nameCollection)

}
