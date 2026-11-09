package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"
)

var money = atomic.Int64{}
var bank = atomic.Int64{}

func payHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading request body:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	httpRequestBodyStr := string(httpRequestBody)
	paymentAmount, err := strconv.Atoi(httpRequestBodyStr)
	if err != nil {
		fmt.Println("Error converting request body to integer:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if money.Load() >= int64(paymentAmount) {
		time.Sleep(3 * time.Second) // Simulate processing time
		money.Add(-int64(paymentAmount))
		fmt.Println("Payment successful. Remaining balance:", money.Load())
	} else {
		fmt.Println("Insufficient funds for payment. Current balance:", money.Load())
	}

	fmt.Println("Received payment amount:", paymentAmount)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading request body:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	httpRequestBodyStr := string(httpRequestBody)
	saveAmount, err := strconv.Atoi(httpRequestBodyStr)
	if err != nil {
		fmt.Println("Error converting request body to integer:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if money.Load() >= int64(saveAmount) {
		time.Sleep(3 * time.Second) // Simulate processing time
		money.Add(-int64(saveAmount))
		bank.Add(int64(saveAmount))
		fmt.Println("Saved amount:", saveAmount)
		fmt.Println("Remaining balance:", money.Load())
		fmt.Println("Bank balance:", bank.Load())
	} else {
		fmt.Println("Insufficient funds to save. Current balance:", money.Load())
	}

}

func main() {
	money.Add(50)
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/save", saveHandler)

	port := ":9091"
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Server error:", err.Error())
	}
}
