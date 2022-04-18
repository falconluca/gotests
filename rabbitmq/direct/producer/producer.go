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

	exchangeName := "exchange_demo"
	if err = channel.ExchangeDeclare(exchangeName, "direct",
		// 持久化
		true,
		// 非自动删除
		false,
		false, false, nil); err != nil {
		log.Fatal(err)
	}

	queue, err := channel.QueueDeclare("queue_demo",
		true,
		false,
		// 非排他
		false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	routingKey := "routingKey_demo"
	if err = channel.QueueBind(queue.Name, exchangeName, routingKey, false, nil); err != nil {
		log.Fatal(err)
	}

	msg := "Greetings!!!"
	if err = channel.Publish(exchangeName, routingKey, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent, // 设置消息为持久化
		ContentType:  "text/plain",
		Body:         []byte(msg),
	}); err != nil {
		log.Fatal(err)
	}
}
