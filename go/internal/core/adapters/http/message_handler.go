package http

import (
	"encoding/json"
	"go-service/internal/core/app"
	"go-service/internal/core/domain"
	"log"
	"net/http"
)

type MessageHandler struct {
	service *app.MessageService
}

func NewMessageHandler(service *app.MessageService) *MessageHandler {
	return &MessageHandler{
		service: service,
	}
}

func (h *MessageHandler) HandleMessages(w http.ResponseWriter, r *http.Request) {
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

	log.Printf("Received message from %s: %s", msg.Author, msg.Body)

	err = h.service.ProcessMessage(&msg)
	if err != nil {
		log.Printf("Error processing message: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{"status": "message received and published"})
}

