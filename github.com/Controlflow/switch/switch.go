package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("this will not print")
	case (2 == 3):
		fmt.Println("this should not print")
	case (3 == 3):
		fmt.Println("prints")
		fallthrough
	case (4 == 4):
		fmt.Println("true")
		fallthrough
	default:
		fmt.Println("just default")
	}
}
