package consumer

import (
	"email-todo-service/internal/dto"
	"email-todo-service/internal/service"

	"github.com/rabbitmq/amqp091-go"

	"encoding/json"
	"log"
)

type EmailConsumer struct {
	emailService *service.EmailService
}

func NewEmailConsumer(es *service.EmailService) *EmailConsumer {
	return &EmailConsumer{emailService: es}
}

func (c *EmailConsumer) Start(ch *amqp091.Channel) error {
	err := ch.ExchangeDeclare(
		"email.notifications",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	q, err := ch.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	err = ch.QueueBind(q.Name, "", "email.notifications", false, nil)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {
			var event dto.UserRegisteredEvent
			if err := json.Unmarshal(msg.Body, &event); err != nil {
				msg.Nack(false, true)
				log.Printf("Failed to unmarshal event: %v", err)
				continue
			}

			log.Printf("Received registration event for %s (%s)", event.FullName, event.Email)

			c.emailService.SendWelcomeEmail(event.Email, event.FullName)
			msg.Ack(false)
		}
	}()

	log.Println("Email consumer started, listening to email.notifications")
	return nil
}
