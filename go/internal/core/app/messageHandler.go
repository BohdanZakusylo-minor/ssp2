package app

import (
	"encoding/json"
	"fmt"
	"go-service/internal/core/domain"
	"go-service/internal/core/services"
	"net/http"
)

func HandleMessages(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)

		return
	}

	var msg domain.Message
	err := json.NewDecoder(r.Body).Decode(&msg)

	if err != nil {
		http.Error(w, "JSON is invalid", http.StatusBadRequest)

		return
	}

	fmt.Printf("Message %s: %s\n", msg.Author, msg.Body)

	services.ConnectToRabbit()

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{"status": "received"})
}
