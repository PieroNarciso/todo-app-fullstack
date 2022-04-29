package configs

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
    var mongoUri string = os.Getenv("MONGO_URI")
    client, err := mongo.NewClient(options.Client().ApplyURI(mongoUri))
    if err != nil {
        log.Fatal(err)
    }
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }
    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    return client
}

var DB = ConnectDB()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
    return client.Database("test").Collection(collectionName)
}
