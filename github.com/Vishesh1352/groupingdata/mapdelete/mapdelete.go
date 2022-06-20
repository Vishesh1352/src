package main

import "fmt"

func main() {
	m := map[string]int{
		"james":  32,
		"mepote": 23,
	}
	fmt.Println(m)
	delete(m, "james")
	fmt.Println(m)
}
