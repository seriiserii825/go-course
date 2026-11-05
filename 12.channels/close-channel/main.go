package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	transferPoint := make(chan int)

	go func() {
		iterations := 3 + rand.Intn(4)
		fmt.Println(iterations, "iterations")

		for i := 1; i <= iterations; i++ {
			transferPoint <- 10
			time.Sleep(100 * time.Millisecond)
		}
		close(transferPoint)
	}()

	coal := 0

	for v := range transferPoint {
		coal += v
		fmt.Println(coal, "coal")
	}

	fmt.Println("Total coal: ", coal)
}
