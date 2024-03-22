package sender

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"pontomenos-api/models"

	"github.com/rabbitmq/amqp091-go"
)

type PontoSender struct {
    rabbitMQConnection *amqp091.Connection
}

func NewPontoSender(rabbitMQConnection *amqp091.Connection) *PontoSender {
    return &PontoSender{rabbitMQConnection: rabbitMQConnection}
}

func (ps *PontoSender) EnviaRegistroParaFila(login string) error {
    ch, err := ps.rabbitMQConnection.Channel()
    if err != nil {
        return err
    }
    defer ch.Close()

    // Declara a fila como durável
    queue, err := ch.QueueDeclare(
        "registroPontoEvento", // nome da fila
        true,                  // durável
        false,                 // delete quando não utilizada
        false,                 // exclusiva
        false,                 // no-wait
        nil,                   // argumentos
    )
    if err != nil {
        return err
    }

    registro := models.RegistroPontoEvento{
        DataHora: time.Now(),
        Login:    login,
    }

    body, err := json.Marshal(registro)
    if err != nil {
        return err
    }

    // Publica a mensagem com contexto e marca como persistente
    err = ch.PublishWithContext(
        context.Background(), // contexto
        "",                  // exchange
        queue.Name,          // routing key (nome da fila)
        false,               // mandatory
        false,               // immediate
        amqp091.Publishing{
            ContentType:  "application/json",
            Body:         body,
            DeliveryMode: amqp091.Persistent, // mensagem persistente
        },
    )
    if err != nil {
        return err
    }

    log.Printf(" [x] Sent %s", body)
    return nil
}
