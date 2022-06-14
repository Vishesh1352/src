package main

import "fmt"

func main() {
	m := map[string]int{
		"james":  32,
		"mepote": 23,
	}
	fmt.Println(m)

	fmt.Println(m["james"])

	fmt.Println(m["yash"]) //no value then retuns zero value

	v, ok := m["yash"] //coma ok ididiom
	fmt.Println(v)
	fmt.Println(ok)
	if v, ok := m["yash"]; ok {
		fmt.Println("this is the if print ", v)
	}
}
