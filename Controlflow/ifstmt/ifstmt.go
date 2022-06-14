package main

import "fmt"

func main() {
	if true {
		fmt.Println("true")
	}
	if false {
		fmt.Println("false")
	}
	if !true {
		fmt.Println("!true")
	}
	if !false {
		fmt.Println("!false")
	}
	if 2 == 2 {
		fmt.Println("2==2")
	}
	if 2 != 2 {
		fmt.Println("2!=2")
	}
	if !(2 == 2) {
		fmt.Println("!(2==2)")
	}
	if !(2 != 2) {
		fmt.Println("!(2!=2)")
	}
}
