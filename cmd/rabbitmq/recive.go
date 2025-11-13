package rabbitmq

import (
	"log"
	"ticket/pkg/rabbitmq"
)

func Receive() {
	// Initialize RabbitMQ first
	if err := rabbitmq.Init(); err != nil {
		log.Fatal("Failed to initialize RabbitMQ:", err)
	}

	// Now you can use other functions
	rabbitmq.ConsumeQueue("new_queue")
}
