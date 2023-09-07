package interfaces

import "contacts/models"

type IContact interface {
	AddContact(contact models.Contact) (int, error)
	GetContactById(id int) (models.Contact, error)
	DeleteContact(id int) (int64, error)
}
