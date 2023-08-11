package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	c := new(Console)
	fw := new(FW)
	fw.FileName = "log.txt"

	sms := new(SMS)
	sms.FROM = "9618558500"
	sms.TO = "979273192"
	fmt.Fprintln(c, "Hello Victoria Secrets")
	fmt.Fprintln(fw, "Hello Victoria Secrets")
	fmt.Fprintln(sms, "Hello Victoria Secrets")
}

//Write(p []byte) (n int, err error)

type Console struct{}

func (c *Console) Write(bytes []byte) (n int, err error) {
	return fmt.Println(string(bytes))
	//return 0, nil
}

type FW struct {
	FileName string
}

func (f *FW) Write(bytes []byte) (n int, err error) {
	if f == nil {
		return 0, errors.New("nil file object")
	}

	err = os.WriteFile(f.FileName, bytes, 0644)
	if err != nil {
		return 0, err
	}
	return len(bytes), nil

}

type SMS struct {
	FROM, TO string
}

func (s *SMS) Write(bytes []byte) (n int, err error) {
	if s == nil {
		return 0, errors.New("nil file object")
	}

	fmt.Println("Message has been sent:", string(bytes))

	return len(bytes), nil

}
