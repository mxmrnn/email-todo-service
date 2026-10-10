package main

import (
	"email-todo-service/internal/config"
	"log"

	"email-todo-service/internal/consumer"
	"email-todo-service/internal/service"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	conn, err := amqp091.Dial(cfg.RabbitMQ.URL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	emailService := service.NewEmailService()
	consumer := consumer.NewEmailConsumer(emailService)

	if err := consumer.Start(ch); err != nil {
		log.Fatalf("Failed to start consumer: %v", err)
	}

	log.Println("Email Service started. Waiting for events...")

	// Бесконечный цикл
	select {}
}
