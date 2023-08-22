package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetConnection(url string) (*mongo.Client, error) {

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	opts := options.Client().ApplyURI(url).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	// todo retry logic
	return mongo.Connect(context.TODO(), opts)

}

func Ping(client *mongo.Client, dbname string) error {
	if err := client.Database(dbname).RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		return err
	}
	return nil
}
