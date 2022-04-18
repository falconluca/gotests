package rabbitmq_test

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"testing"
)

// https://pkg.go.dev/github.com/streadway/amqp
// https://github.com/rabbitmq/rabbitmq-tutorials/tree/master/go
type AMQPClient struct {
	connection *amqp.Connection
	channel    *amqp.Channel
}

func NewAMQPConnection() *AMQPClient {
	conn, err := amqp.Dial("amqp://root:root@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}

	// channel 不是 goroutine 安全的
	channel, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	return &AMQPClient{
		connection: conn,
		channel:    channel,
	}
}

func (c AMQPClient) Release() error {
	var err error
	if c.connection != nil {
		err = c.connection.Close()
	}
	if c.channel != nil {
		err = c.channel.Close()
	}
	return err
}

func TestExchangeDeclarePassive(t *testing.T) {
	cli := NewAMQPConnection()
	defer cli.Release()

	// 检测交换机是否存在。若不存在，则抛出异常。
	if err := cli.channel.ExchangeDeclarePassive("demo:3:exchange:not-fount", "direct",
		true, false, false, false, nil); err != nil {
		t.Fatal("exchange declare passive failed:", err)
		// Exception (404) Reason: "NOT_FOUND - no exchange 'demo:3:exchange' in vhost '/'"
	}
}

func TestExchangeDeclare(t *testing.T) {
	cli := NewAMQPConnection()
	defer cli.Release()

	exchangeName := "demo:3:exchange"
	if err := cli.channel.ExchangeDeclare(exchangeName,
		"direct",
		// 是否将交换机存盘（重启不会丢失数据）
		true,
		// 当交换机没有绑定任何队列或交换机时，删除本交换机
		false,
		// 是否开启内置交换机。客户端无法将消息直接发到内置交换机，必须通过交换机发送。
		false,
		false,
		// TODO 参见 4.1.3
		nil); err != nil {
		t.Fatal(err)
	}
}

// 删除交换机
func TestDeleteChannel(t *testing.T) {
	cli := NewAMQPConnection()
	defer cli.Release()

	if err := cli.channel.ExchangeDelete("demo:3:exchange",
		// ifUnused:true 仅删除没有绑定队列或交换机的
		true, false); err != nil {
		// Exception (406) Reason: "PRECONDITION_FAILED - exchange 'demo:3:exchange' in vhost '/' in use"
		t.Fatal(err)
	}
}

func TestDeclareQueue(t *testing.T) {
	cli := NewAMQPConnection()
	defer cli.Release()

	_, err := cli.channel.QueueDeclare("my:worker:queue",
		// 是否将队列存盘（重启不会丢失数据）
		true,
		// 1)至少有一个消费者连接，2)并且当与这个队列连接的所有消费者都断开时，删除队列
		false,
		// 是否为排他队列（仅对声明它的连接可见，连接断开时自动删除队列）TODO
		false,
		false,
		nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteQueue(t *testing.T) {
	cli := NewAMQPConnection()
	defer cli.Release()

	purgedMessages, err := cli.channel.QueueDelete("my:worker:queue", true, true, false)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("purged Messages: %d\n", purgedMessages)
}

func TestQueueBind(t *testing.T) {
	cli := NewAMQPConnection()
	defer cli.Release()

	if err := cli.channel.QueueBind("sasd", "dor", "asdasd", false, nil); err != nil {
		log.Fatal(err)
	}
}

func TestQueueUnbind(t *testing.T) {
	cli := NewAMQPConnection()
	defer cli.Release()

	if err := cli.channel.QueueUnbind("CrazyQueue2", "crazyDog", "", nil); err != nil {
		log.Fatal(err)
	}
}
