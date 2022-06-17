package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}
type ByName []Person

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

func main() {
	people := []Person{
		{"bob", 31},
		{"yash", 24},
		{"vid", 22},
		{"bharat", 20},
	}
	fmt.Println(people)
	sort.Sort(ByName(people))
	fmt.Println(people)
}
