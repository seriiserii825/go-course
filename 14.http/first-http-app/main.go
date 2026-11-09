package main

import (
	"fmt"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {
	str := "Hello, World!"
	b := []byte(str)
	_, err := w.Write(b)
	if err != nil {
		fmt.Println("have error:", err.Error())
		return
	} else {
		fmt.Println("Response sent successfully")
	}
}
func payHandler(w http.ResponseWriter, r *http.Request) {
	str := "Payment successful!"
	b := []byte(str)
	_, err := w.Write(b)
	if err != nil {
		fmt.Println("have error:", err.Error())
		return
	} else {
		fmt.Println("Payment response sent successfully")
	}
}
func cancelHandler(w http.ResponseWriter, r *http.Request) {
	str := "Payment cancelled!"
	b := []byte(str)
	_, err := w.Write(b)
	if err != nil {
		fmt.Println("have error:", err.Error())
		return
	} else {
		fmt.Println("Payment cancellation response sent successfully")
	}
}

func main() {
	http.HandleFunc("/", handle)
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/cancel", cancelHandler)
	port := ":9091"

	fmt.Println("Server is running on port ", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Server error:", err.Error())
	}
}
