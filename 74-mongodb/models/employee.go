package models

type Employee struct {
	ID           int `bson:"_id"`
	Name         string
	Email        string
	Mobile       string
	Status       string
	LastModified int64
}
