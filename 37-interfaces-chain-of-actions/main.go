package main

import "fmt"

type ICalculator interface {
	Add(int) ICalculator
	Sub(int) ICalculator
	Mul(int) ICalculator
	Get() int
}

type IntType struct {
	Result int
}

func (i *IntType) Add(num int) ICalculator {
	i.Result = i.Result + num
	return i
}
func (i *IntType) Sub(num int) ICalculator {
	i.Result = i.Result - num
	return i
}
func (i *IntType) Mul(num int) ICalculator {
	if i.Result == 0 {
		i.Result = 1
	}
	i.Result = i.Result * num
	return i
}

func (i *IntType) Get() int {
	return i.Result
}

func main() {
	var icalc ICalculator

	icalc = new(IntType)

	r1 := icalc.Add(10).Add(20).Add(20).Add(30).Add(30).Get()
	fmt.Println("Result r1:", r1)

	icalc = new(IntType)
	r2 := icalc.Add(10).Sub(20).Add(20).Mul(2).Get()
	fmt.Println("Result r2:", r2)

}
