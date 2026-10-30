package main

import (
	"github.com/k0kubun/pp/v3"
	"github.com/seriiserii825/go-course/10.struct-payment/methods"
	"github.com/seriiserii825/go-course/10.struct-payment/payments"
)

func main() {
	method := methods.NewBank()
	paymentModule := payments.NewPaymentModule(method)
	paymentModule.Pay("Burger", 10)
	phoneId := paymentModule.Pay("Phone", 500)
	gameId := paymentModule.Pay("Game", 100)

	paymentModule.Cancel(phoneId)

	allInfo := paymentModule.AllInfo()
	pp.Println(allInfo)

	gameInfo := paymentModule.Info(gameId)
	pp.Println(gameInfo)
}
