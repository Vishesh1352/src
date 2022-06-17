package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	people := []Person{
		{"bob", 31},
		{"yash", 24},
		{"vid", 22},
		{"bharat", 20},
	}
	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Sprintln(people)
}
