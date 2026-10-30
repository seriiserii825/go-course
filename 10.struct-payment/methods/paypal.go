package methods

import (
	"fmt"
	"math/rand"
)

type PayPal struct{}

func NewPayPal() PayPal {
	return PayPal{}
}

func (c PayPal) Pay(usd int) int {
	fmt.Println("Pay with paypal")
	fmt.Println("Usd: ", usd)

	return rand.Int()
}

func (c PayPal) Cancel(id int) {
	fmt.Println("PayPal was cancelled with id: ", id)
}
