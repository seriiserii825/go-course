package main

import (
	"fmt"
	"math/rand"
	"time"
)

func mine(n int) int {
	fmt.Println("Go to mine start", n)
	time.Sleep(1 * time.Second)
	fmt.Println("Go to mine finished", n)
	return 10
}

func mineWithChannel(transferPoint chan int, n int) {
	fmt.Println("Go to mine start", n)
	time.Sleep(1 * time.Second)
	fmt.Println("Go to mine finished", n)
	transferPoint <- 10
	fmt.Println("Give coal")
}

func consecutiveMine() {
	coal := 0

	initTime := time.Now()

	coal += mine(1)
	coal += mine(2)
	coal += mine(3)
	coal += mine(4)
	coal += mine(5)
	coal += mine(6)
	coal += mine(7)
	coal += mine(8)
	coal += mine(9)
	coal += mine(10)

	fmt.Println("Got coal: ", coal)
	fmt.Println("Time was passed: ", time.Since(initTime))
}

func parallelMine() {

	transferPoint := make(chan int)

	coal := 0

	initTime := time.Now()

	go mineWithChannel(transferPoint, 1)
	go mineWithChannel(transferPoint, 2)
	go mineWithChannel(transferPoint, 3)
	go mineWithChannel(transferPoint, 4)
	go mineWithChannel(transferPoint, 5)
	go mineWithChannel(transferPoint, 6)
	go mineWithChannel(transferPoint, 7)
	go mineWithChannel(transferPoint, 8)
	go mineWithChannel(transferPoint, 9)
	go mineWithChannel(transferPoint, 10)

	coal += <-transferPoint
	coal += <-transferPoint
	coal += <-transferPoint
	coal += <-transferPoint
	coal += <-transferPoint
	coal += <-transferPoint
	coal += <-transferPoint
	coal += <-transferPoint
	coal += <-transferPoint
	coal += <-transferPoint

	fmt.Println("Got coal: ", coal)
	fmt.Println("Time was passed: ", time.Since(initTime))
}

func parallelMineWithBuffer() {
	transferPoint := make(chan int, 10)

	coal := 0

	initTime := time.Now()

	go mineWithChannel(transferPoint, 1)
	go mineWithChannel(transferPoint, 2)
	go mineWithChannel(transferPoint, 3)
	go mineWithChannel(transferPoint, 4)
	go mineWithChannel(transferPoint, 5)
	go mineWithChannel(transferPoint, 6)
	go mineWithChannel(transferPoint, 7)
	go mineWithChannel(transferPoint, 8)
	go mineWithChannel(transferPoint, 9)
	go mineWithChannel(transferPoint, 10)

	time.Sleep(400 * time.Millisecond)
	coal += <-transferPoint
	time.Sleep(400 * time.Millisecond)
	coal += <-transferPoint
	time.Sleep(400 * time.Millisecond)
	coal += <-transferPoint
	time.Sleep(400 * time.Millisecond)
	coal += <-transferPoint
	time.Sleep(400 * time.Millisecond)
	coal += <-transferPoint
	time.Sleep(400 * time.Millisecond)
	coal += <-transferPoint
	time.Sleep(400 * time.Millisecond)
	coal += <-transferPoint
	time.Sleep(400 * time.Millisecond)
	coal += <-transferPoint
	time.Sleep(400 * time.Millisecond)
	coal += <-transferPoint
	time.Sleep(400 * time.Millisecond)
	coal += <-transferPoint

	fmt.Println("Got coal: ", coal)
	fmt.Println("Time was passed: ", time.Since(initTime))
}

func parallelMineFor() {
	transferPoint := make(chan int)

	coal := 0

	initTime := time.Now()

	totalIterations := 1000

	for i := 1; i <= totalIterations; i++ {
		go mineWithChannel(transferPoint, i)
	}

	for i := 1; i <= totalIterations; i++ {
		coal += <-transferPoint
	}

	fmt.Println("Got coal: ", coal)
	fmt.Println("Time was passed: ", time.Since(initTime))
}

func randomIterations() {
	transferPoint := make(chan int)

	timeSpent := time.Now()

	go func() {
		iterations := 3 + rand.Intn(40)
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
	totalTime := time.Since(timeSpent)

	fmt.Println("Total coal: ", coal)
	fmt.Println(totalTime, "totalTime")
}

func main() {
	consecutiveMine()
	// parallelMine()
	// parallelMineWithBuffer()
	// parallelMineFor()
	// randomIterations()
}
