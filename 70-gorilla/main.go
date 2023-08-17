package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var (
	PORT   string
	logger *log.Logger
)

func init() {
	// f, err := os.OpenFile("log.txt", syscall.O_RDWR|syscall.O_APPEND, 0644)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	logger = log.New(os.Stdout, "", log.Lshortfile|log.Ldate|log.Ltime)
	//logger = log.New(f, "", log.Lshortfile|log.Ldate|log.Ltime)
}

func main() {

	flag.StringVar(&PORT, "port", "50090", "--port=50090 or -port=50090")
	flag.Parse()

	router := mux.NewRouter()
	srv := http.Server{
		Addr:    ":" + PORT,
		Handler: router,
	}
	logger.Println("Server started  and listening on port", PORT)

	router.Use(Auth)
	router.HandleFunc("/ping", Ping)

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		// todo
	})

	router.HandleFunc("/greet", Greet("Hello Vitoria Secrets & Co"))

	router.Use(Audit)

	if err := srv.ListenAndServe(); err != nil {
		logger.Fatalln(err)
	}

}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := r.Header.Get("user")
		authCode := r.Header.Get("authCode")

		log.Println(user, authCode)
		if user == "Victoria" && authCode == "123456" {

			log.Println("auth successfully done")
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(401)
			w.Write([]byte("un authenticated user"))
			log.Println("unauthenticated user")
			return
		}
	})
}

func Ping(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(400)
		w.Write([]byte("unimplemented http method"))
		logger.Println("unimplemented http method.")
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("pong"))
	logger.Println("ping pong successfully hit.")

}

func Greet(msg string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(msg))
	}

}

func Audit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Println("successfully executed")
		next.ServeHTTP(w, r)
	})
}

// func Square(n int) int {
// 	return n * n
// }

// func add(a, b int) int {
// 	return a + b
// }
