package main

func main() {

	// check whether a number is divisible by 2,4 and 8
	println("fallthrough implementation")
	divisibleBy842(32) // to show it is divisible by 8,4 and 2
	divisibleBy842(28) // to show it is divisible by 4 and 2
	divisibleBy842(6)  // to show it is divisible only by 2
	divisibleBy842(5)  // to show it reaches default
	println("wrong results as false negative.Wrong implentation of fallthrough")
	falseNegativeImplementationOfFallthrough(28) // wrong results as false negative.Wrong implentation of fallthrough
	falseNegativeImplementationOfFallthrough(6)  // wrong results as false negative.Wrong implentation of fallthrough

}

func divisibleBy842(num int) {
	switch {
	case num%8 == 0:
		println(num, "is divisible by 8")
		fallthrough
	case num%4 == 0:
		println(num, "is divisible by 4")
		fallthrough
	case num%2 == 0:
		println(num, "is divisible by 2")
	default:
		println("divisible by all or none .. Just for demo purpose, as there is no number that is not divisible by none")
	}
}

func falseNegativeImplementationOfFallthrough(num int) {
	switch {
	case num%2 == 0:
		println(num, "is divisible by 2")
		fallthrough
	case num%4 == 0:
		println(num, "is divisible by 4")
		fallthrough
	case num%8 == 0:
		println(num, "is divisible by 8")
	}
}
