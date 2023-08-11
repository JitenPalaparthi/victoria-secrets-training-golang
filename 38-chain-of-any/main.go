package main

import "fmt"

func main() {

	var icalc ICalc

	var a1 Any

	var num1 float32 = 10.56
	icalc = &a1
	r1 := icalc.Add(10).Add(10).Add(10.24).Add(float32(10.24)).Add(num1).Get()
	fmt.Println(r1)
	fmt.Printf("%.2f\n", r1)

}

type ICalc interface {
	Add(num any) ICalc
	Get() any
}

type Any float64

func (a Any) Add(num any) ICalc {
	switch num.(type) {
	case int:
		a = a + Any(num.(int))
	case float32:
		a = a + Any(num.(float32))
	case float64:
		a = a + Any(num.(float64))
	}
	return &a
}

func (a *Any) Get() any {
	return *a
}
