package main

import (
	"contacts/database"
	"contacts/models"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	dsn := `admin:admin123@tcp(127.0.0.1:33306)/contactsdb?allowNativePasswords=false&checkConnLiveness=false&maxAllowedPacket=0`
	db, err := database.GetConnection(dsn)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully connected", db.Stats())
	// contact := models.Contact{Name: "Jiten", Email: "JitenP@Outlook.Com", Mobile: "9618558500", Status: "active", LastModified: time.Now().Unix()}

	// contactDb := database.ContactDB{DB: db}

	// id, err := contactDb.AddContact(contact)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Contact successfully inserted with the id", id)
	// }

	// contactDb := database.ContactDB{DB: db}
	// contact, err := contactDb.GetContactById(1)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(contact)
	// }

	// // connConfig := mysql.Config{
	// // 	User:   "admin",
	// // 	Passwd: "admin123",
	// // 	Net:    "tcp",
	// // 	Addr:   "127.0.0.1:33306",
	// // 	DBName: "contactsdb",
	// // }
	// //fmt.Println(connConfig.FormatDSN())
	// //db, err := sql.Open("mysql", connConfig.FormatDSN())
	// db, err := sql.Open("mysql", dsn)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// err = db.Ping()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// gin gonic

	router := gin.Default()

	// context.Context and gin.Context is a http context(containd request and response) are different.
	router.GET("/ping", func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{"message": "pong"})
		c.String(200, "pong")
	})

	router.POST("/contact", func(c *gin.Context) {

		contact := new(models.Contact)
		err := c.Bind(contact)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		}

		contactDb := database.ContactDB{DB: db}
		contact.Status = "active"
		contact.LastModified = time.Now().Unix()
		contact, err = contactDb.AddContact(contact)

		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		}

		c.JSON(201, contact)
		return
		//json.NewDecoder(c.Request.Body).Decode(contact)
	})

	router.DELETE("/contact/{id}", func(ctx *gin.Context) {
		// todo here
	})

	router.Run(":58081")

}
