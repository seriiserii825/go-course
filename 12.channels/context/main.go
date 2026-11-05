package main

import (
	"context"
	"fmt"
	"time"
)

func foo(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Foo end")
			return
		default:
			fmt.Println("Foo")
		}
		time.Sleep(100 * time.Millisecond)
	}
}
func boo(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Boo end")
			return
		default:
			fmt.Println("Boo")
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	parentContext, parentCancel := context.WithCancel(context.Background())
	childContext, childCancel := context.WithCancel(parentContext)

	go foo(parentContext)
	go boo(childContext)

	time.Sleep(1 * time.Second)
	childCancel()

	time.Sleep(3 * time.Second)
	parentCancel()
}
