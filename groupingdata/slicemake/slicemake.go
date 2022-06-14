//make is in build function

package main

import "fmt"

func main() {
	x := make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x = append(x, 10, 20)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x = append(x, 20, 30)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
