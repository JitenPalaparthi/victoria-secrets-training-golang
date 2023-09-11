package database

import (
	"database/sql"
	"errors"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	Max_No_Of_Retries  = 5
	Max_Time_For_Retry = 5 // 10 sec
)

var (
	ErrNilDB = errors.New("nil db object")
)

func GetConnection(dsn string) (*sql.DB, error) {
	counter := 1
loop:
	db, err := sql.Open("mysql", dsn)
	if err == nil {
		err = db.Ping()
	}
	if err != nil {
		log.Println("trying to connect to database-->", counter, "times")
		//return nil, err
		if counter >= Max_No_Of_Retries {
			return nil, err
		}
		counter++
		time.Sleep(time.Second * time.Duration(Max_Time_For_Retry))
		goto loop
	}

	return db, nil
}
