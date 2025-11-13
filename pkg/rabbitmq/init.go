package rabbitmq

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
)

var (
	once    sync.Once
	initErr error
	queues  map[string]amqp.Queue
)

func Init() error {
	once.Do(func() {
		// Initialize queues map FIRST
		queues = make(map[string]amqp.Queue)
		port := viper.GetString("RABBITMQ_PORT")
		user := viper.GetString("RABBITMQ_USER")
		pass := viper.GetString("RABBITMQ_PASS")
		host := viper.GetString("RABBITMQ_HOST")
		conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", user, pass, host, port))
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
func DeclareQueue(queueName string) amqp.Queue {
	ch := GetRabbitmqChannel()
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
	ch := GetRabbitmqChannel()
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
	ch := GetRabbitmqChannel()

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
