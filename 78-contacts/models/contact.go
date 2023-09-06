package models

type Contact struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Mobile       string `json:"mobile"`
	Status       string `json:"status"`
	LastModified int64  `json:"lastModified"`
}
