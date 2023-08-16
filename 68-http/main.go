package main

import (
	"demo/handlers"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
)

var (
	PORT string
)

func main() {
	PORT = os.Getenv("PORT")
	if PORT == "" {
		flag.StringVar(&PORT, "port", "9091", "--port=9091 or -port=9091")
	}

	//flag.String(PORT, "9091", "--port=9091 or -port=9091")
	flag.Parse()
	fmt.Println(PORT)

	fmt.Println("Server started and listening on port", PORT)
	//30,000

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "pong")
	})

	http.HandleFunc("/greet", Sayhi)

	http.HandleFunc("/greetby", GreetBy)

	ehandler := &handlers.EmployeeHandler{}
	http.HandleFunc("/employee/add", ehandler.AddEmployee)

	//go http.ListenAndServe(":9092", nil)

	err := http.ListenAndServe("0.0.0.0:"+PORT, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(err, "Hello World")
}

func GreetBy(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Println("Hit here?")
		g := Greet{}
		err := json.NewDecoder(r.Body).Decode(&g) // takes data from requese body
		// use json decoder and decode it to a greet struct variable
		// if it cannot decode write it as an error

		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println(err)
		} else {
			fmt.Fprintln(w, "Data has been read", g)
		}
	} else {
		w.Write([]byte("Unimpleentd method"))
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func Sayhi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Vitoria Secrets & Co")
}

type Greet struct {
	Message string `json:"message"`
	Name    string `json:"name"`
}
