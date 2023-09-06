package database

import (
	"contacts/models"
	"database/sql"
	"errors"
)

type ContactDB struct {
	//DB *sql.DB
	*sql.DB // promoted field
}

func (c *ContactDB) AddContact(contact *models.Contact) (int64, error) {
	//result, err := c.DB.Exec("INSERT INTO contacts(name,email,mobile,status,lastModified) VALUES(?,?,?,?,?)", contact.Name, contact.Email, contact.Mobile, contact.Status, contact.LastModified)
	result, err := c.Exec("INSERT INTO contacts(name,email,mobile,status,lastModified) VALUES(?,?,?,?,?)", contact.Name, contact.Email, contact.Mobile, contact.Status, contact.LastModified)

	if err != nil {
		return -1, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (c *ContactDB) DeleteContact(id int) (int64, error) {
	//result, err := c.DB.Exec("INSERT INTO contacts(name,email,mobile,status,lastModified) VALUES(?,?,?,?,?)", contact.Name, contact.Email, contact.Mobile, contact.Status, contact.LastModified)
	result, err := c.Exec("DELETE from contacts where id=?", id)

	if err != nil {
		return -1, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return -1, err
	}
	return affected, nil
}

func (c *ContactDB) GetContactById(id int) (*models.Contact, error) {

	row := c.QueryRow("SELECT id,name,email,mobile,status,lastModified FROM contacts WHERE id = ?", id)

	contact := models.Contact{}

	if err := row.Scan(&contact.Id, &contact.Name, &contact.Email, &contact.Mobile, &contact.Status, &contact.LastModified); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("no contact found")
		}
		return &contact, err
	}

	return &contact, nil
}
