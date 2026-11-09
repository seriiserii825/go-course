package postman

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Postman(ctx context.Context, wg *sync.WaitGroup, transferPoint chan<- string, n int, mail string) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Posman nr:", n, "------------- Закончил работу -------------")
			return
		default:
			fmt.Println("Mailer nr:", n, "Взял письмо")
			time.Sleep(1 * time.Second)
			fmt.Println("Mailer nr:", n, "Донес письмо до почты", mail)

			transferPoint <- mail
			fmt.Println("Mailer nr:", n, "Передал письмо", mail)
		}
	}
}

func PostmanPool(ctx context.Context, postmanCount int) <-chan string {
	mailTransferPoint := make(chan string)
	wg := &sync.WaitGroup{}

	for i := 1; i <= postmanCount; i++ {
		wg.Add(1)
		go Postman(ctx, wg, mailTransferPoint, i, postmanToMail(i))
	}
	go func() {
		wg.Wait()
		close(mailTransferPoint)
	}()
	return mailTransferPoint
}

func postmanToMail(postmanNumber int) string {
	postmanMap := map[int]string{
		1: "Family hello",
		2: "Inviter from friend",
		3: "Info from autoservice",
	}

	mail, ok := postmanMap[postmanNumber]
	if !ok {
		return "Lottery"
	}
	return mail
}
