package main

import (
	"fmt"
	"reflect"
)

func main() {

	//var iface1 interface{} // Object obj

	var iface1 any // type inference gives nil

	fmt.Println("Value of", iface1, "Type Of", reflect.TypeOf(iface1))

	iface1 = 1000 // can directly assign a value
	fmt.Println("Value of", iface1, "Type Of", reflect.TypeOf(iface1))

	var num1 int = iface1.(int) // int(iface1) not okay becase iface1 is interface{} or any type
	// type assertion    any_var.(T) // T means type that empty interface has
	fmt.Println(num1)

	var float1 float32 = 10.1231

	iface1 = float1 // or assign a variable

	float2 := iface1.(float32)

	fmt.Println(float2)

	fmt.Println("Value of", iface1, "Type Of", reflect.TypeOf(iface1))

	var ok1 bool = true

	iface1 = ok1

	ok2 := iface1.(bool)
	fmt.Println(ok2)
	fmt.Println("Value of", iface1, "Type Of", reflect.TypeOf(iface1))

	str1 := "Victoria Secrets & Co"

	iface1 = str1
	str2 := iface1.(string)
	fmt.Println(str2)
	fmt.Println("Value of", iface1, "Type Of", reflect.TypeOf(iface1))
	//var ok3 bool = iface1.(bool) // this does not work becase it has string value, not bool
	//fmt.Println(ok3)
}
