package main

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("anonymus function")
	}()

	func(x int) {
		fmt.Println("Anonymus function with argument", x)
	}(52)
}
