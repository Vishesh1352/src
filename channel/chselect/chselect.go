package main

import (
	"fmt"

)

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)
	//send
	go send(eve, odd, quit)
	//receive
	receive(eve, odd, quit)
	fmt.Println("about to exit")
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("EVEN", v)

		case v := <-o:
			fmt.Println("ODD", v)

		case i, ok := <-q:
			if !ok {
				fmt.Println("quit", i, ok)
				return
			} else {
				fmt.Println("from comma ok", i)
				return
			}
		}
	}
}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(q)
}
