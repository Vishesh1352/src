package main

import "fmt"

func main() {

	p1 := struct {
		first string
		last  string
		age   int
	}{first: "james",
		last: "hold",
		age:  55,
	}

	fmt.Println(p1)
}
