package main

import "fmt"

type hotdog int

var a int
var b hotdog

func main() {
	b = 43
	a = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
}
