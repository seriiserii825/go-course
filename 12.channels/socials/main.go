package main

import (
	"fmt"
	"time"
)

type Message struct {
	Author string
	Text   string
}

func main() {
	messageChan1 := make(chan Message)
	messageChan2 := make(chan Message)

	go func() {
		for {
			messageChan1 <- Message{
				Author: "Friend 1",
				Text:   "Hello",
			}
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			messageChan2 <- Message{
				Author: "Friend 2",
				Text:   "How are you?",
			}
			time.Sleep(3 * time.Second)
		}
	}()

	for {
		select {
		case msg1 := <-messageChan1:
			fmt.Println("Message from: ", msg1.Author, "message: ", msg1.Text)

		case msg2 := <-messageChan2:
			fmt.Println("Message from: ", msg2.Author, "message: ", msg2.Text)
		}
	}

}
