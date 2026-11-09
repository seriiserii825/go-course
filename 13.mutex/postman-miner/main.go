package main

import (
	"context"
	"fmt"
	"time"

	"github.com/seriiserii825/go-course/13.mutex/postman-miner/miner"
	"github.com/seriiserii825/go-course/13.mutex/postman-miner/postman"
)

func main() {
	var coal int
	var mails []string

	isCoalClosed := false
	isMailClosed := false

	minerContext, minerCancel := context.WithCancel(context.Background())
	postmanContext, postmanCancel := context.WithCancel(context.Background())

	coalTransferPoint := miner.MinerPool(minerContext, 2)
	postmanTransferPoint := postman.PostmanPool(postmanContext, 2)

	go func() {
		time.Sleep(3 * time.Second)
		minerCancel()
	}()

	go func() {
		time.Sleep(5 * time.Second)
		postmanCancel()
	}()
	for !isCoalClosed || !isMailClosed {
		select {
		case c, ok := <-coalTransferPoint:
			if !ok {
				isCoalClosed = true
				continue
			}
			coal += c
		case m, ok := <-postmanTransferPoint:
			if !ok {
				isMailClosed = true
				continue
			}
			mails = append(mails, m)
		}
	}
	fmt.Println("Total coal:", coal)
	fmt.Println("Total letters: ", len(mails))
}
