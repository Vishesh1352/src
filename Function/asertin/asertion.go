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
	fmt.Println("first :", s.first, "last:", s.last, "-the secretAgenet")
}
func (p person) speak() {
	fmt.Println("first :", p.first, "last:", p.last, "-the person speak()")
}

type human interface {
	speak()
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("pass in to barr", h.(person).first)
	case seceretAgenet:
		fmt.Println("pass in to barr", h.(seceretAgenet).first)
	}
	fmt.Println("i ws passed into bar", h)
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
	bar(sa1)
	bar(p1)
	bar(p2)
}
