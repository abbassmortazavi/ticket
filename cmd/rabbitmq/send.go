package rabbitmq

import (
	"context"
	"log"
	"ticket/pkg/rabbitmq"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Send() {
	if !rabbitmq.IsConnected() {
		log.Println("RabbitMQ not connected, initializing...")
		if err := rabbitmq.Init(); err != nil {
			log.Fatalf("Failed to initialize RabbitMQ: %v", err)
		}
	}

	conn := rabbitmq.GetRabbitmqConn()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

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

	body := "halet chetore........."
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
