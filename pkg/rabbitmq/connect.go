package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

func GetRabbitmqConn() *amqp.Connection {
	return rabbitmqConn
}
