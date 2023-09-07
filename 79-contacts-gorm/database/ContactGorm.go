package database

import (
	"contacts/models"

	"gorm.io/gorm"
)

type ContactGorm struct {
	DB *gorm.DB
}

// Do you need to create a table from code?
// or table is already created?
func (cg *ContactGorm) AddContact(contact *models.Contact) (*models.Contact, error) {
	if cg.DB == nil {
		return nil, ErrNilDB
	}
	cg.DB.AutoMigrate(models.Contact{}) // This creates a table
	tx := cg.DB.Create(contact)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return contact, nil
}

func (cg *ContactGorm) DeleteContact(id string) (int64, error) {
	if cg.DB == nil {
		return -1, ErrNilDB
	}
	tx := cg.DB.Delete(&models.Contact{}, id)
	if tx.Error != nil {
		return -1, tx.Error
	}
	return tx.RowsAffected, nil

}

func (cg *ContactGorm) GetContactById(id string) (*models.Contact, error) {
	if cg.DB == nil {
		return nil, ErrNilDB
	}
	contact := new(models.Contact)

	tx := cg.DB.First(contact, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return contact, nil
}
