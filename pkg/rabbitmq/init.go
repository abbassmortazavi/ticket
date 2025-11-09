package rabbitmq

import (
	"log"
	"sync"
	_ "sync"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	once    sync.Once
	initErr error
)

func Init() error {
	once.Do(func() {
		conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
		if err != nil {
			initErr = err
			return
		}
		rabbitmqConn = conn
		log.Println("Connected to RabbitMQ")
	})
	return initErr
}
func IsConnected() bool {
	return rabbitmqConn != nil && !rabbitmqConn.IsClosed()
}
