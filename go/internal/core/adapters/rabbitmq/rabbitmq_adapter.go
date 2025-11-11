package rabbitmq

import (
	"encoding/json"
	"go-service/internal/core/domain"
	"go-service/internal/core/ports"
	"github.com/streadway/amqp"
)

type RabbitMQService struct {
	channel   *amqp.Channel
	queueName string
}

func NewRabbitMQService(connectionURL, queueName string) (ports.MessagePublisher, error) {
	conn, err := amqp.Dial(connectionURL)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	_, err = channel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		channel.Close()
		conn.Close()
		return nil, err
	}

	return &RabbitMQService{channel: channel, queueName: queueName}, nil
}

func (r *RabbitMQService) Publish(message *domain.Message) error {
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return r.channel.Publish("", r.queueName, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	})
}

func (r *RabbitMQService) Close() error {
	return r.channel.Close()
}

