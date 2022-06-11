package main

import "fmt"

func main() {
	fmt.Println("hello go.")
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Print(i)
		}
	}
	foo()
}
func foo() {
	fmt.Print("im in foo.")
}
