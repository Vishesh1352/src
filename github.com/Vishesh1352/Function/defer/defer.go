package main

import "fmt"

func main() {
	a := sum(10, 20)
	fmt.Println(a)
	defer foo()
	bar()
}
func foo() {

	fmt.Println("foo")
}
func bar() {
	fmt.Println("bar")
}
func sum(x int, y int) int {
	sum := x + y
	return sum
}
