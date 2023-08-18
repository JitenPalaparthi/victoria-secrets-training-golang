package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Employee struct {
	Client     *mongo.Client
	Dbname     string
	Collection string
}

func (e *Employee) Insert() error {
	_, err := e.Client.Database(e.Dbname).Collection(e.Collection).InsertOne(context.TODO(), e)
	return err
}
