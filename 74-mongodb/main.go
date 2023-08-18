// RrAGlf7TWLCc1BoS
// victoriasecrets
//mongodb+srv://victoriasecrets:RrAGlf7TWLCc1BoS@demo-cluster.iejivhk.mongodb.net/?retryWrites=true&w=majority"

package main

import (
	"context"
	"demo/database"
	"fmt"
)

var (
	dsn = "mongodb+srv://victoriasecrets:RrAGlf7TWLCc1BoS@demo-cluster.iejivhk.mongodb.net/?retryWrites=true&w=majority"
)

func main() {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	client, err := database.GetConnection(dsn)
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

}
