package main

import (
	"bytes"
	"github.com/streadway/amqp"
	"log"
	"time"
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

	// 流量控制
	if err = channel.Qos(1, 0, false); err != nil {
		log.Fatal(err)
	}

	msgs, err := channel.Consume("queue_demo", "",
		false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
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
