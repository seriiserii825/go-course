package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// writeJSON сериализует v и пишет его с указанным статусом.
func writeJSON(w http.ResponseWriter, status int, v any) {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write response:", err)
	}
}

// writeError пишет ErrorDto с указанным статусом.
func writeError(w http.ResponseWriter, status int, err error) {
	http.Error(w, NewErrorDto(err.Error()).ToString(), status)
}
