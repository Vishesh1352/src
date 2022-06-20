package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{10, 20, 50, 12, 63, 54, 2, 55, 5, 9, 0, 5, 5, 0}
	xs := []string{"a", "z", "d", "f", "r", "h"}

	sort.Ints(xi)
	fmt.Println(xi)
	sort.Strings(xs)
	fmt.Sprintln(xs)
}
