package database

import (
	"errors"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	Max_No_Of_Retries  = 5
	Max_Time_For_Retry = 5 // 10 sec
)

var (
	ErrNilDB = errors.New("nil db object")
)

type IConnection interface {
	GetConnection(dsn string) (any, error)
}

type MySql struct {
}

func (m1 *MySql) GetConnection(dsn string) (any, error) {
	counter := 1
loop:
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
