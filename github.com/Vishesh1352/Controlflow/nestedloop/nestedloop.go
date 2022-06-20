package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println("oter loop :", i)
		for j := 0; j < 3; j++ {
			fmt.Println("inner loop :", j)
		}
	}
}
