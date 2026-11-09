package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/seriiserii825/go-course/13.mutex/postman-miner/miner"
	"github.com/seriiserii825/go-course/13.mutex/postman-miner/postman"
)

func main() {
	var coal atomic.Int64
	var mails []string
	mtx := sync.Mutex{}

	workersCount := 200

	minerContext, minerCancel := context.WithCancel(context.Background())
	postmanContext, postmanCancel := context.WithCancel(context.Background())

	coalTransferPoint := miner.MinerPool(minerContext, workersCount)
	postmanTransferPoint := postman.PostmanPool(postmanContext, workersCount)

	initTime := time.Now()

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("=========== Рабочий день шахтеров закончился ========")
		minerCancel()
	}()

	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("============= Рабочий день почтальонов закончился ========")
		postmanCancel()
	}()

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for c := range coalTransferPoint {
			coal.Add(int64(c))
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for m := range postmanTransferPoint {
			mtx.Lock()
			mails = append(mails, m)
			mtx.Unlock()
		}
	}()

	wg.Wait()

	fmt.Println("Total coal:", coal.Load())
	mtx.Lock()
	fmt.Println("Total letters: ", len(mails))
	mtx.Unlock()
	fmt.Println("Total time:", time.Since(initTime))
}
