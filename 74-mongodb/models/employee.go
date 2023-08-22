package models

import (
	"encoding/json"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrInvalidMobile = errors.New("invalid mobile number")
)

type Employee struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"` //omit empty : if no value is given dont give any value..
	Name         string             `json:"name" bson:"name"`
	Email        string             `json:"email" bson:"email"`
	Mobile       string             `json:"mobile" bson:"mobile"`
	Status       string             `json:"status" bson:"status"`
	LastModified int64              `json:"lastModified" bson:"lastModified"`
}

func (e *Employee) Validate() error {
	if e.Name == "" {
		return errors.New("invalid name field")
	}
	if e.Email == "" {
		//return errors.New("invalid email field")
		return fmt.Errorf("invalid email field")
	}
	if e.Mobile == "" {
		return ErrInvalidMobile
	}
	return nil
}

func (e *Employee) ToBytes() ([]byte, error) {
	return json.Marshal(e)
}

func (e *Employee) ToString() (string, error) {
	bytes, err := json.Marshal(e)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
