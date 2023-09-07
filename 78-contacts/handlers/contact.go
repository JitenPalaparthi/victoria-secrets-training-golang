package handlers

import (
	"contacts/interfaces"
	"contacts/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ContactHandler struct {
	IContact interfaces.IContact
}

func (ch *ContactHandler) AddContact(c *gin.Context) {
	contact := new(models.Contact)
	err := c.Bind(contact)
	//err := json.NewDecoder(c.Request.Body).Decode(contact)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
	}
	//contactDb := database.ContactDB{DB: db}
	contact.Status = "active"
	contact.LastModified = time.Now().Unix()
	contact, err = ch.IContact.AddContact(contact)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
	}
	c.JSON(201, contact)
	//c.XML(200, contact)
	//c.String(200, "%v,%v,%v,%v,%v,%v", contact.Id, contact.Name, contact.Email, contact.Mobile, contact.Status, contact.LastModified)
	return
	//json.NewDecoder(c.Request.Body).Decode(contact)
}
