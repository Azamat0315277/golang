package main

import (
	"bytes"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// Setup connection
	conn, err := amqp.Dial("amqp://quest:quest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"task_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exlusive
		false,        // no-wait
		nil,          //arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set Qos")

	msgs, err := ch.Consume(
		q.Name,
		"",
		false, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   //args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			dotCount := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dotCount)
			time.Sleep(t * time.Second)
			log.Println("Done")
			d.Ack(false) // -> Note below
			// Using this code, you can ensure that even if you terminate a worker using CTRL+C while it was processing a message, nothing is lost.
			// Soon after the worker terminates, all unacknowledged messages are redelivered.

			// Acknowledgement must be sent on the same channel that received the delivery.
			// Attempts to acknowledge using a different channel will result in a channel-level protocol exception.

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
