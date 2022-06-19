package main

import "fmt"

func main() {
	var v [5]int
	fmt.Println(v)
	v[3] = 52
	fmt.Println(v)
	fmt.Println(len(v))
}
