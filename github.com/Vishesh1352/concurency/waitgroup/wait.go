package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("goroutine\t", runtime.NumGoroutine())
	go foo()
	wg.Add(1)
	bar()
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("goroutine\t", runtime.NumGoroutine())
	wg.Wait()

}
func foo() {
	fmt.Println("foo")
	wg.Done()
}

func bar() {
	fmt.Println("bar")
}
