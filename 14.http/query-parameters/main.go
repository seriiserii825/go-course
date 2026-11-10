package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fooParam := r.URL.Query().Get("foo")
	booParam := r.URL.Query().Get("boo")

	fmt.Println(fooParam, "fooParam")
	fmt.Println(booParam, "booParam")
}

func main() {
	http.HandleFunc("/", handler)

	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
