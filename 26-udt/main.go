package main

import "fmt"

func main() {

	var p1 Person // declaration

	p1.Id = 101
	p1.Name = "Victoria"
	p1.Email = "Victoria@VictoriaSecrets.com"
	p1.Mobile = "+911231231123"

	p2 := Person{Id: 102, Name: "VS", Email: "vs@vs.com", Mobile: "+9112312"} // shorthand

	p3 := Person{}
	p3.Id = 101
	p3.Name = "Victoria"
	p3.Email = "Victoria@VictoriaSecrets.com"
	p3.Mobile = "+911231231123"

	fmt.Println(p1, p2, p3)
}

type Person struct {
	Id     int
	Name   string
	Email  string
	Mobile string
}
