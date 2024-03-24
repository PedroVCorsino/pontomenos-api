package rabbitMQ

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func ConectarRabbitMQ() *amqp091.Connection {
	conn, err := amqp091.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatalf("Falha ao conectar com RabbitMQ: %s", err)
	}
	return conn
}
