package main

import (
	"fmt"
	"time"
)

func postman(text string) {
	for i := 1; i <= 3; i++ {
		fmt.Printf("Postman, delivered %s, %d time\n", text, i)
		time.Sleep(750 * time.Millisecond)
	}
}

func main() {
	postman("New York")
	postman("Moscow")
}
