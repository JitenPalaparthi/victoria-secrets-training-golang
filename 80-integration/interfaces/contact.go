package interfaces

import "contacts/models"

type IContact interface {
	AddContact(contact *models.Contact) (*models.Contact, error)
	GetContactById(id string) (*models.Contact, error)
	DeleteContact(id string) (int64, error)
}
