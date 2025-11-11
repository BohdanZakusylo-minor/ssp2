package main

import (
	httpadapter "go-service/internal/core/adapters/http"
	"go-service/internal/core/adapters/rabbitmq"
	"go-service/internal/core/app"
	"log"
	"net/http"
)

func main() {
	publisher, err := rabbitmq.NewRabbitMQService("amqp://user:password@rabbit-mq:5672/", "southpark_messages")

	if err != nil {
		log.Fatal(err)
	}
	defer publisher.(*rabbitmq.RabbitMQService).Close()

	messageService := app.NewMessageService(publisher)

	handler := httpadapter.NewMessageHandler(messageService)
	http.HandleFunc("/messages", handler.HandleMessages)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
