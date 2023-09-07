package main

import (
	"contacts/database"
	"contacts/handlers"
	"fmt"

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

	// gin gonic

	router := gin.Default()

	// context.Context and gin.Context is a http context(containd request and response) are different.
	router.GET("/ping", func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{"message": "pong"})
		c.String(200, "pong")
	})

	contactDb := &database.ContactDB{DB: db}
	contacthandler := handlers.ContactHandler{IContact: contactDb}

	router.POST("/contact", contacthandler.AddContact)

	router.DELETE("/contact/{id}", func(ctx *gin.Context) {
		// todo here
	})

	router.Run(":58081")

}
