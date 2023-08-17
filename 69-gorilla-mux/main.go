package main

import (
	"demo/handlers"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

var (
	PORT   string
	chanId chan int
)

func main() {

	flag.StringVar(&PORT, "port", "50080", "--port=50080 or -port=50080")
	flag.Parse()
	chanId = make(chan int)

	router := mux.NewRouter()

	srv := http.Server{
		Addr:         ":" + PORT,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router, // Pass our instance of gorilla/mux in.
	}

	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	router.HandleFunc("/greet/{name}", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		fmt.Println(params["name"])

		fmt.Fprintln(w, params["name"])

		return
	})

	eh := new(handlers.EmployeeHandler)
	router.HandleFunc("/employee/add", eh.Add(chanId))

	router.HandleFunc("/employee/delete/{id}", func(w http.ResponseWriter, r *http.Request) {})
	router.HandleFunc("/employee/get/{id}", func(w http.ResponseWriter, r *http.Request) {})

	fmt.Println("Server started and listning on port->", PORT)
	go IncrementEmployeeId(chanId)
	if err := srv.ListenAndServe(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func IncrementEmployeeId(chanId chan int) {
	for id := range chanId {
		f, err := os.OpenFile("employees/employee-id", syscall.O_RDWR, 0644)
		if err != nil {
			log.Println(err)
		}
		_, err = f.Write([]byte(strconv.Itoa(id)))
		if err != nil {
			log.Println(err)
		}
	}
}
