package main

import (
	"fmt"
)

func main() {
	a := sum(10, 20, 30, 40, 50)
	fmt.Println(a)
}

func sum(x ...int) int {
	sum := 0
	for i, v := range x {
		sum = sum + v
		fmt.Println("index", i, "sum is", sum)
	}
	fmt.Println("total sum", sum)
	return sum

}
