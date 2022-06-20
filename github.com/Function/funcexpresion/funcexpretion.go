package main

import (
	"fmt"
)

func main() {
	f := func() {
		fmt.Println("anonymus function")
	}
	f()
	g := func(x int) {
		fmt.Println("Anonymus function with argument", x)
	}
	g(52)
}
