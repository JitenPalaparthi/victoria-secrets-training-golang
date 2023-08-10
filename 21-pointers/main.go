package main

import (
	"errors"
	"fmt"
)

func main() {

	var num1 int = 100
	var ptr1 *int = &num1  // This holds the address of a variable
	var ptr2 **int = &ptr1 // This holds the address of another pointer

	fmt.Println("Value of num1", num1, "address of num1", &num1)
	fmt.Println("Value of num1 using ptr1", *ptr1, "address of num1 using ptr1", ptr1)
	fmt.Println("Value of num1 using ptr2", **ptr2, "address of num1 using ptr2", *ptr2)

	var ptr3 *int
	if ptr3 == nil {
		fmt.Println("ptr3 is a nil pointer")
	} else {
		fmt.Println("Value that  ptr3 is pointed to", *ptr3) // nil pointer
	}

	Incr(&num1) // valid
	fmt.Println(num1)

	Incr(ptr1) // valid
	fmt.Println(num1)

	Incr(*ptr2) // valid i.e it is a pointer to the pointer
	fmt.Println(num1)

	Incr(ptr3) // nil pointer
	fmt.Println(num1)

	ptr4 := new(int) // new allocates memonry of a given type and returns a pointer
	fmt.Println(*ptr4)
	*ptr4 = 201

	fmt.Println(*ptr4)
	ptr5 := new(any)

	fmt.Println(ptr5)
	fmt.Println(*ptr5)
	*ptr5 = 100
	fmt.Println(*ptr5)
	*ptr5 = "Victoria Secrets"
	fmt.Println(*ptr5)
}

func Incr(num *int) error {
	if num == nil {
		return errors.New("nil pointer")
	}
	*num++
	return nil
}
