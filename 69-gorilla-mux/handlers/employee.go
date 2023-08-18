package handlers

import (
	"demo/models"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"syscall"
	"time"

	"github.com/golang/glog"
)

type EmployeeHandler struct {
	MU *sync.Mutex
}

func (eh *EmployeeHandler) Add(chanId chan int) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// fetch a value from employees/employe-id
		// increment that value
		// fetch details from request body
		// assign the value to employee object
		// creat a file with the id and write employee data into that file
		// rewrite the emoployee-id file with a new number
		//f, err := os.Open("employees/employee-id")
		if r.Method != "POST" {
			w.WriteHeader(http.StatusNotImplemented)
			w.Write([]byte("unimplemented method"))
			return
		}

		// now

		e := new(models.Employee)
		err := json.NewDecoder(r.Body).Decode(e)
		if err != nil {
			glog.Errorln(err)
			w.WriteHeader(400)
			w.Write([]byte("there seems to be some thing went wrong.Please try again or contact admin"))
			return
		}
		err = e.Validate()
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}

		eh.MU.Lock()
		bytes, err := os.ReadFile("employees/employee-id")

		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}

		str := string(bytes)
		_id, err := strconv.Atoi(str)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}

		e.Id = _id + 1
		e.Status = "active"
		e.LastModified = time.Now().Unix()

		f, err := os.OpenFile("employees/"+strconv.Itoa(e.Id), syscall.O_RDWR|syscall.O_CREAT, 0777)
		if err != nil {
			log.Println(err)
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}
		buf, err := e.ToBytes()
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}

		_, err = f.Write(buf)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}

		// f, err = os.OpenFile("employees/employee-id", syscall.O_RDWR, 0644)
		// if err != nil {
		// 	w.WriteHeader(400)
		// 	w.Write([]byte(err.Error()))
		// 	return
		// }

		// _, err = f.Write([]byte(strconv.Itoa(e.Id)))
		// if err != nil {
		// 	w.WriteHeader(400)
		// 	w.Write([]byte(err.Error()))
		// 	return
		// }

		chanId <- e.Id
		w.WriteHeader(201)
		w.Write(buf)
		eh.MU.Unlock()

	}

}
