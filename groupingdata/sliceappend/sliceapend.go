package main

import "fmt"

func main() {
	x := []int{5, 8, 15, 45, 52}
	fmt.Println(x)
	x = append(x, 44, 55, 66, 77, 85)
	fmt.Println(x)
	y := []int{2, 66, 32, 52, 485, 854686}
	x = append(x, y...)
	fmt.Println(x)
}
