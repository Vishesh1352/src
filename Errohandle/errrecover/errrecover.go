package main

import (
	"fmt"

)

func main() {
	fo()
	fmt.Print("Hello, world!")
}
func fo(){
	defer func ()  {
		if r := recover(); r!=nil{
			fmt.Println("Recover in f",r)
		}
	}()
	fmt.Println("calling g")
	g(0)
}

func g (i int){
	if i>3{
		fmt.Println("PANIKING!")
		panic(fmt.Sprintf("%T\n",i))
	}
	defer fmt.Println("defer in go", i)

	fmt.Println("Printing in g",i)
	g(i+1)
}
