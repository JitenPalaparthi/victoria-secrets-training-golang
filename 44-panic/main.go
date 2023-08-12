package main

import "fmt"

func main() {

	fn1 := "Victoria"
	ln1 := "Secrets & Co"

	str1 := GetFullName(&fn1, &ln1)

	fmt.Println("Full Name:", *str1)

	fn2, ln2 := new(string), new(string)

	*fn2, *ln2 = "Victoria", "Secrets & Co"
	str2 := GetFullName(fn2, ln2)

	fmt.Println("Full Name:", *str2)

	str3 := GetFullName(fn2, nil)
	fmt.Println("Full Name:", *str3)

}

func GetFullName(firstName, lastName *string) *string {

	if firstName == nil || *firstName == "" {
		panic("firstName cannot be nil")
	}

	if lastName == nil || *lastName == "" {
		panic("lastName cannot be nil")
	}

	str := *firstName + *lastName
	return &str
}

// When to raise panic:
// Without a successful execution, if your application cannot proceed further then you can consider
// raising a panic
// Ex: No database connection, Not able to write to a file, Not able to open a network connection
