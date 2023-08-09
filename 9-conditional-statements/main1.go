package main

func main() {
	a(4, 5)
}

func a(m, n int) int {
	if m == 0 {
		return n + 1
	} else if m > 0 && n == 0 {
		return a(m-1, 1)
	} else {
		return a(m-1, a(m, n-1))
	}
}
