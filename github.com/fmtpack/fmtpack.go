package main

import "fmt"

func main() {
	y := 101
	fmt.Printf("%T\n", y)
	fmt.Print(y)
	fmt.Printf("%b\t%T\t\n", y, y)
}
