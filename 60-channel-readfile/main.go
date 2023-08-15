package main

import (
	"os"
	"syscall"
)

func main() {
	os.OpenFile("data.ext", syscall.O_RDONLY, 0644)
	// To explain it later
}
