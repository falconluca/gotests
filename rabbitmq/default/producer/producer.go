package main

import (
	"github.com/streadway/amqp"
	"log"
)

func main() {
	conn, err := amqp.Dial("amqp://root:root@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	channel, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer channel.Close()

	//channel.ExchangeDeclare()

	queue, err := channel.QueueDeclare("DotQueue",
		true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	msg := "data...."
	err = channel.Publish("", queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent, // 设置消息为持久化
		ContentType:  "text/plain",
		Body:         []byte(msg),
	})
	if err != nil {
		log.Fatalf("failed to publish a msg: %s\n", msg)
	}
	log.Printf("X sent %s", msg)
}
