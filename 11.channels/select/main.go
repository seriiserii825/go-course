package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	startIndex := 1
	intChn := make(chan int)
	strChn := make(chan string)
	go func() {
		i := 1
		for {
			intChn <- i
			i++
			time.Sleep(100 * time.Millisecond)
		}
	}()
	go func() {
		i := 1
		for {
			strChn <- "hi " + strconv.Itoa(i)
			i++
			time.Sleep(1000 * time.Millisecond)
		}
	}()
	for startIndex <= 20 {
		select {
		case number := <-intChn:
			fmt.Println("intChn: ", number)

		case str := <-strChn:
			fmt.Println("strChn: ", str)
		}
		startIndex++
	}
}
