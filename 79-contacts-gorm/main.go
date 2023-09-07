package main

import (
	"contacts/database"
	"contacts/handlers"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := `admin:admin123@tcp(127.0.0.1:33406)/contactsdb?allowNativePasswords=false&checkConnLiveness=false&maxAllowedPacket=0`
	//db, err := database.GetConnection(dsn)
	/*
		mysqldb := new(database.MySql)
		var db *sql.DB
		db1, err := mysqldb.GetConnection(dsn)

		if err != nil {
			fmt.Println(err)
			return
		}
		db = db1.(*sql.DB)
		fmt.Println("Successfully connected", db.Stats())
	*/

	mysqlgorm := new(database.MySqlGorm)
	var db *gorm.DB
	db2, err := mysqlgorm.GetConnection(dsn)
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

	router.Run(":58081")

}
