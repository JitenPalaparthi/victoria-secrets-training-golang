package database

import (
	"context"
	"demo/models"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrNil = errors.New("nil connection")
)

type Employee struct {
	Client     *mongo.Client
	Dbname     string
	Collection string
}

func (e *Employee) Insert(ctx context.Context, employee *models.Employee) (any, error) {
	if e.Client == nil {
		return nil, ErrNil
	}
	result, err := e.Client.Database(e.Dbname).Collection(e.Collection).InsertOne(ctx, employee)

	return result.InsertedID, err
}

func (e *Employee) Delete(ctx context.Context, id string) (int64, error) {
	if e.Client == nil {
		return 0, ErrNil
	}
	objid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}
	result, err := e.Client.Database(e.Dbname).Collection(e.Collection).DeleteOne(ctx, bson.D{{"_id", objid}})

	return result.DeletedCount, err
}

func (e *Employee) GetByID(ctx context.Context, id string) (*models.Employee, error) {
	if e.Client == nil {
		return nil, ErrNil
	}
	objid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	result := e.Client.Database(e.Dbname).Collection(e.Collection).FindOne(ctx, bson.D{{"_id", objid}})
	emp := new(models.Employee)
	err = result.Decode(emp)
	if err != nil {
		return nil, err
	}
	return emp, nil
}

func (e *Employee) GetByName(ctx context.Context, name string) (*models.Employee, error) {
	if e.Client == nil {
		return nil, ErrNil
	}
	//objid, err := primitive.ObjectIDFromHex(id)
	// if err != nil {
	// 	return nil, err
	// }

	result := e.Client.Database(e.Dbname).Collection(e.Collection).FindOne(ctx, bson.D{{"name", name}})
	emp := new(models.Employee)
	err := result.Decode(emp)
	if err != nil {
		return nil, err
	}
	return emp, nil
}
