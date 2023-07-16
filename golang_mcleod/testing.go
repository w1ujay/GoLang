package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Begin CPU", runtime.NumCPU())
	fmt.Println("Begin GS", runtime.NumGoroutine())
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		fmt.Println("hello From 1st")
		wg.Done()
	}()
	go func() {
		fmt.Println("hello From 2nd")
		wg.Done()
	}()

	fmt.Println("mid CPU", runtime.NumCPU())
	fmt.Println("mid gs", runtime.NumCPU())

	wg.Wait()

	fmt.Println("About to Exit")
	fmt.Println("end CPU", runtime.NumCPU())
	fmt.Println("end gs", runtime.NumGoroutine())
}
