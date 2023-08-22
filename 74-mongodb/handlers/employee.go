package handlers

import (
	"context"
	"demo/database"
	"demo/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployeeHandler struct {
	DB *database.Employee
}

func (eh *EmployeeHandler) Add(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte("not implemented method"))
		return
	}

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

	e.Status = "active"
	e.LastModified = time.Now().Unix()
	result, err := eh.DB.Insert(context.TODO(), e)
	if err != nil {
		glog.Errorln(err)
		w.WriteHeader(400)
		w.Write([]byte("there seems to be some thing went wrong.Please try again or contact admin"))
		return
	}

	w.Write([]byte(result.(primitive.ObjectID).String()))
	w.WriteHeader(201)
	return

}

func (eh *EmployeeHandler) Delete(w http.ResponseWriter, r *http.Request) {

	if r.Method != "DELETE" {
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte("not Implemented method"))
		return
	}
	id := mux.Vars(r)["id"]
	if id == "" {
		w.WriteHeader(400)
		w.Write([]byte("invalid id"))
		return
	}

	result, err := eh.DB.Delete(context.TODO(), id)
	if err != nil {
		glog.Errorln(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(fmt.Sprint(result)))
	w.WriteHeader(202)
	return

}
