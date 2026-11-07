package main

import (
	"fmt"
	"sync"
	"time"
)

func postman(wg *sync.WaitGroup, text string) {
	defer wg.Done()
	for i := 1; i <= 3; i++ {
		fmt.Printf("Postman, delivered %s, %d time\n", text, i)
		time.Sleep(750 * time.Millisecond)
	}
}

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go postman(wg, "New York")

	wg.Add(1)
	go postman(wg, "Moscow")

	wg.Add(1)
	go postman(wg, "Canada")

	wg.Add(1)
	go postman(wg, "Tokio")

	wg.Add(1)
	go postman(wg, "London")

	wg.Wait()
}
