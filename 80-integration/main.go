package main

import (
	"contacts/database"
	"contacts/handlers"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var (
	DSN  string
	PORT string
)

func init() {

	DSN = os.Getenv("DSN")
	PORT = os.Getenv("PORT")

}

func main() {

	//dsn := `admin:admin123@tcp(127.0.0.1:33406)/contactsdb?allowNativePasswords=false&checkConnLiveness=false&maxAllowedPacket=0`

	if DSN == "" {
		log.Fatalln("no DSN(database connection) is provided")
	}

	if PORT == "" {
		log.Fatalln("PORT is not given")
	}

	mysqlgorm := new(database.MySql)
	var db *gorm.DB
	db2, err := mysqlgorm.GetConnection(DSN)
	if err != nil {
		fmt.Println(err)
		return
	}

	db = db2.(*gorm.DB)
	fmt.Println("Successfully connected")

	// gin gonic

	router := gin.Default()

	// context.Context and gin.Context is a http context(containd request and response) are different.
	router.GET("/ping", func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{"message": "pong"})
		c.String(200, "pong")
	})

	//contactDb := &database.ContactDB{DB: db}
	contactDb := &database.ContactGorm{DB: db}
	contacthandler := handlers.ContactHandler{IContact: contactDb}

	router.POST("/contact", contacthandler.AddContact)
	router.GET("/contact/:id", contacthandler.GetContactById)
	router.DELETE("/contact/:id", contacthandler.DeleteContact)
	//:id is for gin {id} is for gorilla

	router.Run(":" + PORT)

}
