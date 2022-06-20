package main

import (
	"fmt"
)

func main() {
	li := []int{10, 20, 30, 50, 40, 100}
	s := sum(li...)
	fmt.Println(s)
	e := even(sum, li...)
	fmt.Println(e)
	od := odd(sum, li...)
	fmt.Println(od)
}

func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
func even(f func(xi ...int) int, v ...int) int {
	var yi []int
	for _, v := range yi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
func odd(f func(xi ...int) int, vz ...int) int {
	var yi []int
	for _, v := range yi {
		if v%2 != 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
