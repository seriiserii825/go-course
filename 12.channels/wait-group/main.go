package main

import (
	"fmt"
	"time"
)

func postman(text string, done chan bool) {
	for i := 1; i <= 3; i++ {
		fmt.Printf("Postman, delivered %s, %d time\n", text, i)
		time.Sleep(750 * time.Millisecond)
	}

	done <- true
}

func main() {
	done := make(chan bool)

	go postman("New York", done)
	go postman("Moscow", done)

	<-done
	<-done
}
