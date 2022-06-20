package main

import "fmt"

func main() {
	x := []int{5, 8, 15, 45, 52}
	fmt.Println(x)
	x = append(x[:1], x[3:]...)
	fmt.Println(x)
}
