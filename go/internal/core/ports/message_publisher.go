package ports

import "go-service/internal/core/domain"

type MessagePublisher interface {
	Publish(message *domain.Message) error
}
