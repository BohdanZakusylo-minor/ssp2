package app

import (
	"go-service/internal/core/domain"
	"go-service/internal/core/ports"
)

type MessageService struct {
	publisher ports.MessagePublisher
}

func NewMessageService(publisher ports.MessagePublisher) *MessageService {
	return &MessageService{
		publisher: publisher,
	}
}

func (s *MessageService) ProcessMessage(message *domain.Message) error {
	if err := message.Validate(); err != nil {
		return err
	}

	return s.publisher.Publish(message)
}

