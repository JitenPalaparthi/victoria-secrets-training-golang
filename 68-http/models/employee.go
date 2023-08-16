package models

import (
	"encoding/json"
	"errors"
)

type Employee struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Mobile       string `json:"mobile"`
	Status       string `json:"status"`
	LastModified int64  `json:"lastModified"`
}

func (e *Employee) Validate() error {

	if e.Id == "" || e.Name == "" || e.Email == "" || e.Mobile == "" {
		return errors.New("invalid input data")
	}
	return nil
}

func (e *Employee) ToBytes() ([]byte, error) {
	return json.Marshal(e)
}
