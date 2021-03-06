package services

import (
	"log"
	"microservice/utils"

	"github.com/streadway/amqp"
)

type messageService struct {
	queue   *amqp.Queue
	channel *amqp.Channel
}

// NewMessageService func
func NewMessageService() MessageServiceInterface {
	amqpURL := utils.GetEnvVariable("AMQP_URL")
	connection, connectionErr := amqp.Dial(amqpURL)
	failOnError(connectionErr, "Failed to connect to RabbitMQ")

	channel, channelErr := connection.Channel()
	failOnError(channelErr, "Failed to open a channel")

	queue, queueErr := channel.QueueDeclare(
		"client",
		false,
		false,
		false,
		false,
		nil,
	)

	failOnError(queueErr, "Failed to declare a queue")

	return &messageService{
		queue:   &queue,
		channel: channel,
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func (m messageService) Dequeue() ([]byte, error) {
	message, _, err := m.channel.Get(m.queue.Name, true)

	if err != nil {
		return nil, err
	}

	return message.Body, nil
}
