package services

import (
	"backend/utils"
	"log"

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

// Enqueue func
func (m messageService) Enqueue(message []byte) error {
	err := m.channel.Publish(
		"",
		m.queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         message,
			DeliveryMode: amqp.Persistent,
		})

	return err
}
