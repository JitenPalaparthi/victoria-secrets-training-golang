package main

import (
	"contacts/database"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	dsn := `admin:admin123@tcp(127.0.0.1:33306)/contactsdb?allowNativePasswords=false&checkConnLiveness=false&maxAllowedPacket=0`
	db, err := database.GetConnection(dsn)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully connected")
	// contact := models.Contact{Name: "Jiten", Email: "JitenP@Outlook.Com", Mobile: "9618558500", Status: "active", LastModified: time.Now().Unix()}

	// contactDb := database.ContactDB{DB: db}

	// id, err := contactDb.AddContact(contact)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Contact successfully inserted with the id", id)
	// }

	contactDb := database.ContactDB{DB: db}
	contact, err := contactDb.GetContactById(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(contact)
	}

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

}
