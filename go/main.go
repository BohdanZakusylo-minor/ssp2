package main

import (
	"fmt"
	"go-service/internal/core/app"
	"net/http"
)

func main() {
	http.HandleFunc("/messages", app.HandleMessages)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
