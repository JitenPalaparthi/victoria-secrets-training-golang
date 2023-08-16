package handlers

import (
	"demo/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"syscall"
	"time"
)

type EmployeeHandler struct{}

func (eh *EmployeeHandler) AddEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		e := new(models.Employee)

		err := json.NewDecoder(r.Body).Decode(e)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		if err := e.Validate(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		e.Status = "active"
		e.LastModified = time.Now().UTC().Unix() //

		// if bytes, err := io.ReadAll(r.Body); err != nil {
		// 	w.WriteHeader(http.StatusBadRequest)
		// 	w.Write([]byte(err.Error()))
		// 	return
		// }

		if bytes, err := e.ToBytes(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		} else {
			f, err := os.OpenFile("employee.txt", syscall.O_RDWR, 0644)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}
			_, err = f.Write(bytes)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			} else {
				w.Write([]byte("Data successfully stored in the file"))
				w.WriteHeader(201)
				//w.WriteHeader(http.StatusCreated)
				return
			}
		}
	}
}

func (eh *EmployeeHandler) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values.Get("id")
	if id == "" {
		w.Write([]byte("invalid id"))
		w.WriteHeader(400)
		return
	}
	//fmt.Println(id[0])
	fmt.Fprintln(w, "id query param is-> ", id)

	// Delete that row from the file
	// if no row with that id then it should say "details not found"

}
