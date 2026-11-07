package main

import (
	"fmt"
	"sync"
)

var number int = 0

func increase(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i <= 1000; i++ {
		number++
	}
}

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(10)

	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	wg.Wait()

	fmt.Println("Number: ", number)
}
