package methods

import (
	"fmt"
	"math/rand"
)

type Bank struct{}

func NewBank() Bank {
	return Bank{}
}

func (c Bank) Pay(usd int) int {
	fmt.Println("Pay with bank")
	fmt.Println("Usd: ", usd)

	return rand.Int()
}

func (c Bank) Cancel(id int) {
	fmt.Println("Bank was cancelled with id: ", id)
}
