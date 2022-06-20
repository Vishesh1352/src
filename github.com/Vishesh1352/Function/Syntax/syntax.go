package main

import (
	"fmt"
)

func main() {
	foo()
	bar("vid")
	x, y := exam("vid", "Jolly")
	fmt.Println(x)
	fmt.Println(y)
}

func foo() {
	fmt.Println("foo")
}
func bar(s string) {
	fmt.Println("hello from bar to ", s)
}

func exam(s string, g string) (string, bool) {
	a := fmt.Sprint(s, g, `,says "hello"`)
	b := false
	return a, b
}
