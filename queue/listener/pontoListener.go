package listener

import (
	"log"

	"pontomenos-api/services"

	"github.com/rabbitmq/amqp091-go"
)

type PontoListener struct {
	rabbitMQChannel *amqp091.Channel
	registryService *services.RegistroPontoService
}

func NewPontoListener(channel *amqp091.Channel, registryService *services.RegistroPontoService) *PontoListener {
	return &PontoListener{
		rabbitMQChannel: channel,
		registryService: registryService,
	}
}

func (pl *PontoListener) Listen(queueName string) {
	msgs, err := pl.rabbitMQChannel.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		log.Fatalf("Erro ao registrar consumidor: %s", err)
		return
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Recebido: %s", d.Body)
			pl.registryService.ProcessarMensagem(d.Body)
		}
	}()

	log.Printf(" [*] Aguardando por mensagens. Para sair pressione CTRL+C")
	<-forever
}
