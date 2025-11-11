package rabbitmq

import (
	"context"
	"log"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	once            sync.Once
	initErr         error
	queues          map[string]amqp.Queue
	rabbitmqConn    *amqp.Connection
	rabbitmqChannel *amqp.Channel
)

func Init() error {
	once.Do(func() {
		// Initialize queues map FIRST
		queues = make(map[string]amqp.Queue)

		conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
		if err != nil {
			initErr = err
			return
		}
		rabbitmqConn = conn
		log.Println("Connected to RabbitMQ")

		// channel
		channel, err := conn.Channel()
		if err != nil {
			initErr = err
			return
		}
		rabbitmqChannel = channel
		log.Println("Channel created")
	})
	return initErr
}

// Add connection getters for safety
func GetChannel() *amqp.Channel {
	if rabbitmqChannel == nil {
		log.Fatal("RabbitMQ channel not initialized. Call Init() first.")
	}
	return rabbitmqChannel
}

func GetConnection() *amqp.Connection {
	if rabbitmqConn == nil {
		log.Fatal("RabbitMQ connection not initialized. Call Init() first.")
	}
	return rabbitmqConn
}

func DeclareQueue(queueName string) amqp.Queue {
	ch := GetChannel()
	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue(%s): %s", queueName, err)
	}

	// Safe assignment to queues map
	if queues == nil {
		queues = make(map[string]amqp.Queue)
	}
	queues[queueName] = q

	return q
}

func Consume(q *amqp.Queue) {
	ch := GetChannel()
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer to the queue(%s): %s", q.Name, err)
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func Publish(q *amqp.Queue, body string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ch := GetChannel()

	err := ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		log.Fatalf("Failed to publish a message: %s", err)
	}
	log.Printf(" [x] Sent %s\n", body)
}

func SendMessageToQueue(queue string, body string) {
	q := DeclareQueue(queue)
	Publish(&q, body)
}

func ConsumeQueue(queue string) {
	q := DeclareQueue(queue)
	Consume(&q)
}
