package handlers

import (
	"contacts/interfaces"
	"contacts/models"
	"fmt"
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

func (ch *ContactHandler) GetContactById(c *gin.Context) {
	fmt.Println("----------------->>>>>")
	id, ok := c.Params.Get("id")
	if !ok {
		c.String(http.StatusBadRequest, "id not provided or invalid id")
	}

	contact, err := ch.IContact.GetContactById(id)

	if err != nil {
		c.String(400, err.Error())
		//c.Abort()
		return
	}
	if contact == nil {
		c.String(200, "no record found")
		return
	}
	c.JSON(200, contact)
	return

}

func (ch *ContactHandler) DeleteContact(c *gin.Context) {
	fmt.Println("----------------->>>>>")
	id, ok := c.Params.Get("id")
	if !ok {
		c.String(http.StatusBadRequest, "id not provided or invalid id")
	}

	result, err := ch.IContact.DeleteContact(id)

	if err != nil {
		c.String(400, err.Error())
	}
	if result > 0 {
		c.String(202, "Record successfully deleted;Number of records deleted  %v", result)
	} else {
		c.String(200, "no record found")
	}
	return

}
