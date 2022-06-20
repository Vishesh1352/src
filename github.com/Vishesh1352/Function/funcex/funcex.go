package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}
type seceretAgenet struct {
	person
	first string
	ltk   bool //ltk stnds for liansence to kill
}

func (s seceretAgenet) speak() {
	fmt.Println("first :", s.first, "last:", s.last)
}

func main() {
	sa1 :=
		seceretAgenet{
			person: person{
				first: "james",
				last:  "bald",
				age:   98,
			},
			first: "adding to test collesion",
			ltk:   true,
		}
	p1 :=
		person{
			first: "james",
			last:  "hold",
			age:   55,
		}
	p2 :=
		person{
			first: "carb",
			last:  "kelly",
		}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(sa1)
	fmt.Println(sa1.first, sa1.person.first)
	sa1.speak()
}
