package main

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// connect to RabbitMQ server
	conn, err := amqp.Dial("amqp://quest:quest@localhost:5672/")
	failOnError(err, "failed to connect to RabbitMQ")

	defer conn.Close()

	// Connect to channel
	ch, err := conn.Channel()

	// Declare queue to send information to queue
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := "Hello World!"

	err = ch.PublishWithContext(
		ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
	// Note: Decalring a queue is idempotent - it will only be created if it doesn't exists already.

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
	}
}
