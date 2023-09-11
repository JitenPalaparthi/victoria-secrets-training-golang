package handlers

import (
	"contacts/interfaces"
	"contacts/models"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

type ContactHandler struct {
	IContact interfaces.IContact
	Broker   string
	Topic    string
	ChMsg    chan []byte
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
		return
	}
	c.JSON(201, contact)
	//c.XML(200, contact)
	//c.String(200, "%v,%v,%v,%v,%v,%v", contact.Id, contact.Name, contact.Email, contact.Mobile, contact.Status, contact.LastModified)
	// todo write code to write a produce here

	// The below code is synchronous communication
	// if below service is not available then there is no way that an email can be sent again
	// if you want to scaleup of sending emails. have to scale up this service as well
	// there are many problems with this approach

	/*data, err := json.Marshal(contact)
	err = SendMail("http://localhost:38091/email/send", string(data))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"emai": "sent successfully"})
	}*/

	data, _ := json.Marshal(contact)

	// err = PublishMessage(ch.Broker, ch.Topic, data)
	// if err != nil {
	// 	log.Println(err)
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
	// 	return
	// }

	ch.ChMsg <- data

	return
	//json.NewDecoder(c.Request.Body).Decode(contact)
}

func SendMail(url, data string) error {
	//url := "localhost:38091/email/send"
	method := "POST"

	payload := strings.NewReader(data)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(string(body))
	return nil
}

func PublishMessage(broker, topic string, message []byte) error {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(broker),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
	defer writer.Close()
	err := writer.WriteMessages(context.TODO(),
		kafka.Message{
			//Key:   []byte("Key-A"),
			Value: message,
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func (ch *ContactHandler) Publish() {

	for data := range ch.ChMsg {
		PublishMessage(ch.Broker, ch.Topic, data)
	}

}
