// RrAGlf7TWLCc1BoS
// victoriasecrets
//mongodb+srv://victoriasecrets:RrAGlf7TWLCc1BoS@demo-cluster.iejivhk.mongodb.net/?retryWrites=true&w=majority"

package main

import (
	"context"
	"demo/database"
	"demo/handlers"
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var (
	url = "mongodb+srv://victoriasecrets:RrAGlf7TWLCc1BoS@demo-cluster.iejivhk.mongodb.net/?retryWrites=true&w=majority"

	PORT string
)

func main() {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	client, err := database.GetConnection(url)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Send a ping to confirm a successful connection

	if err := database.Ping(client, "demo"); err != nil {
		panic(err)
	} else {
		fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	}

	flag.StringVar(&PORT, "port", "50080", "--port=50080 or -port=50080")

	flag.Parse()

	router := mux.NewRouter()

	srv := http.Server{
		Addr:         ":" + PORT,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router, // Pass our instance of gorilla/mux in.
	}

	empDB := new(database.Employee)
	empDB.Client = client
	empDB.Dbname = "employedb"
	empDB.Collection = "employees"

	eh := new(handlers.EmployeeHandler)
	eh.DB = empDB

	router.HandleFunc("/employee/add", eh.Add)
	router.HandleFunc("/employee/delete/{id}", eh.Delete)
	srv.ListenAndServe()

	//create an employee object
	// e1 := new(models.Employee)
	// e1.Email = "Jitenp@mymail.com"
	// e1.Mobile = "12312312"
	// e1.Name = "Jiten"
	// e1.Status = "active"
	// e1.LastModified = time.Now().Unix()

	// empDB := new(database.Employee)
	// empDB.Client = client
	// empDB.Dbname = "employedb"
	// empDB.Collection = "employees"

	// result, err := empDB.Insert(context.Background(), e1)
	// fmt.Println(result, err)
}

//mongodb+srv://admin:123456@demo-cluster.iejivhk.mongodb.net/?retryWrites=true&w=majority"
