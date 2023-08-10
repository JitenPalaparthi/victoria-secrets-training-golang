package demo

import "fmt"

var Msg string

func init() {
	Msg = "Victoria Secrets"
	fmt.Println("init called 1st time")

}

func init() {
	fmt.Println("init called 2nd time")
}

func init() {
	fmt.Println("init called 3nd time")
}

func GetMessage() {
	fmt.Println(Msg)
}
