package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// Setting up connection
	conn, err := amqp.Dial("amqp:://quest:quest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Declare queue
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

	failOnError(err, "Failed to declare a queue")
	msgs, err := ch.Consume(
		q.Name,
		"",
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	failOnError(err, "Failed to register a consumer")

	// Telling the server to deliver us the messages from the queue.
	// Since it will push us messages asynchronously, reading the messages from a channel will be (returned by amqp::Consume) in a goroutine.
	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever // blocking

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
