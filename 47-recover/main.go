package main

import "fmt"

func main() {

	fn1 := "Victoria"

	str1 := GetFullName(&fn1, nil)
	//if str1 != nil {
	func() {
		defer Recoverme()
		fmt.Println(*str1)
	}()
	//}

	ln1 := "Secrets & Co"

	str2 := GetFullName(&fn1, &ln1)

	fmt.Println(*str2)
}

func GetFullName(firstName, lastName *string) *string {
	//defer fmt.Println("Is there any panic?")
	defer Recoverme()
	if firstName == nil || *firstName == "" {
		panic("firstName cannot be nil")
	}

	if lastName == nil || *lastName == "" {
		panic("lastName cannot be nil")
	}

	str := *firstName + *lastName
	return &str
}

func Recoverme() {
	if pan := recover(); pan != nil {
		fmt.Println(pan, "I will recover the whole stack except the paniced function")
	}
}
