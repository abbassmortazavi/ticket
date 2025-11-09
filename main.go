package main

import (
	"ticket/cmd"
	"ticket/cmd/rabbitmq"
)

func main() {

	cmd.Execute()
	rabbitmq.Send()
	rabbitmq.Receive()

}
