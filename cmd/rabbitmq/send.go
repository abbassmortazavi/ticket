package rabbitmq

import (
	"log"
	"ticket/pkg/rabbitmq"
)

func Send() {
	// Initialize RabbitMQ first
	if err := rabbitmq.Init(); err != nil {
		log.Fatal("Failed to initialize RabbitMQ:", err)
	}

	log.Println("Send")
	// Now you can use other functions
	rabbitmq.SendMessageToQueue("new_queue", "jafar.com")
}
