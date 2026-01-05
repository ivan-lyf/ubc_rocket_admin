package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("UBC Rocket Admin API is running on :8080")
	
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "ok", "service": "rocket-admin-api"}`))
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}