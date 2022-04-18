package main

import (
	"bytes"
	"github.com/streadway/amqp"
	"log"
	"time"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	channel, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer channel.Close()

	queue, err := channel.QueueDeclare("DotQueue",
		// 支持消息持久化
		true,
		// 非自动删除
		false,
		// 非排他
		false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	//channel.QueueBind()
	//channel.Get() // 拉模式

	msgs, err := channel.Consume(queue.Name, "",
		false, false, false, false, nil)
	//true, false, false, false, nil)
	if err != nil {
		log.Fatalf("failed to register a consumer: %s", err)
	}

	forever := make(chan bool)

	go func() {
		for msg := range msgs {
			log.Printf("received a msg: %s", msg.Body)
			dotCount := bytes.Count(msg.Body, []byte("."))
			time.Sleep(time.Duration(dotCount) * time.Second)
			log.Println("done")
			msg.Ack(false)
		}
	}()

	log.Println("main goroutine waiting for msg...")
	<-forever
}
