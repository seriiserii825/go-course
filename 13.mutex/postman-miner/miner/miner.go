package miner

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func miner(ctx context.Context, wg *sync.WaitGroup, transferPoint chan<- int, n int, power int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Miner nr:", n, "------------- Завершил работу -------------")
			return
		case <-time.After(1 * time.Second):
			time.Sleep(1 * time.Second)
			fmt.Println("Miner nr:", n, "Добыл уголь", power)
		}

		select {
		case <-ctx.Done():
			fmt.Println("Miner nr:", n, "------------- Завершил работу -------------")
			return
		case transferPoint <- power:
			fmt.Println("Miner nr:", n, "Передал уголь", power)
		}

		// select {
		// case <-ctx.Done():
		// 	fmt.Println("Miner nr:", n, "------------- Завершил работу -------------")
		// 	return
		// default:
		// 	fmt.Println("Miner nr:", n, "Начал добывать уголь", power)
		// 	time.Sleep(1 * time.Second)
		// 	fmt.Println("Miner nr:", n, "Добыл уголь", power)
		// 	transferPoint <- power
		// 	fmt.Println("Miner nr:", n, "Передал уголь", power)
		// }
	}
}

func MinerPool(ctx context.Context, minerCount int) <-chan int {
	coalTransferPoint := make(chan int)
	wg := &sync.WaitGroup{}
	for i := 1; i <= minerCount; i++ {
		wg.Add(1)
		go miner(ctx, wg, coalTransferPoint, i, i*10)
	}
	go func() {
		wg.Wait()
		close(coalTransferPoint)
	}()
	return coalTransferPoint
}
