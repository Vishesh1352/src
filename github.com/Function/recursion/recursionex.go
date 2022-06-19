package main

import "fmt"

func main() {
	n := factorial(4)
	fmt.Println(n)
	lf := loopFact(4)
	fmt.Println(lf)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
func loopFact(n int) int {
	tatal := 1
	for ; n > 0; n-- {
		tatal *= n
	}
	return tatal
}
