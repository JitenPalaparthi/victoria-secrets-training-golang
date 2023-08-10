package main

import "fmt"

func main() {
	p1 := Person{Id: 101, Name: "Jitenrere", Email: "Jiten3123@outlook.com", Mobile: "96185534234", Addr: Address{Country: "India", PinCode: "3243242", Line1: "Door No 43-43", Street: "Maha Temple Road", City: "Guntur", State: "AP"}}
	p1.Addr.PinCode = "54445454"
	fmt.Println(p1)
	e1 := Employee{Id: 101, Name: "Jitenrere", Email: "Jiten3123@outlook.com", Mobile: "96185534234", Address: Address{Country: "India", PinCode: "3243242", Line1: "Door No 43-43", Street: "Maha Temple Road", City: "Guntur", State: "AP"}}
	e1.City = "Hyderabad"
	e1.Id = 23123
	e1.Address.Id = 12312

	cd1 := CountryCode{}
	cd1.int = 92
	cd1.string = "Pakisthan"
	cd2 := CountryCode{int: 91, string: "India"}
	fmt.Println(cd1, cd2)

	as1 := struct {
		Name, Email string
	}{
		Name:  "Jiten",
		Email: "Jiten1231@123.com",
	}

	fmt.Println(as1)
}

type Person struct {
	Id     int
	Name   string
	Email  string
	Mobile string
	Addr   Address // composition
	//ShippingAddrs []Address
}

type Address struct {
	Id                                           int
	Line1, City, Street, State, Country, PinCode string
}

type Employee struct {
	Id      int
	Name    string
	Email   string
	Mobile  string
	Address // promoted fileds
}

type int1 = int

type CountryCode struct { // structs only with types
	int
	int1
	string
}

type Student struct {
	Id          int
	Name        string
	Email       string
	SocialMedia struct { // Inline embedded struct
		Linkedin, FaceBook, Twitter, Instagram string
	}
}
