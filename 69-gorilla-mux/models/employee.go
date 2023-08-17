package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

var (
	ErrInvalidMobile = errors.New("invalid mobile filed")
)

type Employee struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Mobile       string `json:"mobile"`
	Status       string `json:"status"`
	LastModified int64  `json:"lastModified"`
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
