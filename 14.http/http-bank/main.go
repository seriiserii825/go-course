package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

var money = 50
var bank = 0
var mtx = sync.Mutex{}

func payHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		msg := "Error reading request body: " + err.Error()
		fmt.Println(msg)
		_, err = w.Write([]byte(msg))
		if err != nil {
			fmt.Println("Error writing response:", err)
		}
		return
	}
	httpRequestBodyStr := string(httpRequestBody)
	paymentAmount, err := strconv.Atoi(httpRequestBodyStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte("Error converting request body to integer: " + err.Error()))
		if err != nil {
			fmt.Println("Error writing response:", err)
		}
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	mtx.Lock()
	if money >= paymentAmount {
		money -= paymentAmount
		msg := "Payment successful. Remaining balance: " + strconv.Itoa(money)
		_, err = w.Write([]byte(msg))
		if err != nil {
			fmt.Println("Error writing response:", err)
		}
	}
	mtx.Unlock()

	fmt.Println("Received payment amount:", paymentAmount)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		msg := "Error reading request body: " + err.Error()
		fmt.Println(msg)
		_, err = w.Write([]byte(msg))
		if err != nil {
			fmt.Println("Error writing response:", err)
		}
		return
	}
	httpRequestBodyStr := string(httpRequestBody)
	saveAmount, err := strconv.Atoi(httpRequestBodyStr)
	if err != nil {
		fmt.Println("Error converting request body to integer:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	mtx.Lock()
	if money >= saveAmount {
		money -= saveAmount
		bank += saveAmount
		msg := "Save successful. Remaining balance: " + strconv.Itoa(money) + ", Bank balance: " + strconv.Itoa(bank)
		_, err = w.Write([]byte(msg))
		if err != nil {
			fmt.Println("Error writing response:", err)
		}
	}
	mtx.Unlock()
}

func main() {
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/save", saveHandler)

	port := ":9091"
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Server error:", err.Error())
	}
}
