package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/email/send", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "POST" {
			data := make(map[string]any)

			err := json.NewDecoder(r.Body).Decode(&data)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusBadRequest)
			}

			// assume that you write email logic over here
			fmt.Fprintln(w, data)

		} else {

			w.Write([]byte("unimplemented method"))
			w.WriteHeader(http.StatusNotImplemented)
			return
		}

	})

	fmt.Println("Server started up and running on port 38091")

	http.ListenAndServe(":38091", nil)

}
