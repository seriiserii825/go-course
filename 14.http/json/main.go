package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/k0kubun/pp/v3"
)

type Payment struct {
	Description string `json:"description"`
	USD         int    `json:"usd"`
	FullName    string `json:"fullName"`
	Address     string `json:"address"`
}

var mtx = sync.Mutex{}
var money = 1000
var paymentHistory = make([]Payment, 0)

func payHandler(w http.ResponseWriter, r *http.Request) {
	var payment Payment

	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		fmt.Println("Error decoding payment:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	mtx.Lock()
	if money-payment.USD >= 0 {
		money -= payment.USD
		paymentHistory = append(paymentHistory, payment)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Payment successful"))
		pp.Println("Payment successful:", payment)
		pp.Println("Remaining money:", money)
	} else {
		w.WriteHeader(http.StatusPaymentRequired)
		w.Write([]byte("Insufficient funds"))
		fmt.Printf("Payment failed: %+v, Remaining money: %d\n", payment, money)
	}
	mtx.Unlock()
}

func main() {
	http.HandleFunc("/pay", payHandler)

	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
