package methods

import (
	"fmt"
	"math/rand"
)

type Crypto struct{}

func NewCrypto() Crypto {
	return Crypto{}
}

func (c Crypto) Pay(usd int) int {
	fmt.Println("Pay with crypto")
	fmt.Println("Usdt: ", usd, " usdt")

	return rand.Int()
}

func (c Crypto) Cancel(id int) {
	fmt.Println("Crypto was cancelled with id: ", id)
}
