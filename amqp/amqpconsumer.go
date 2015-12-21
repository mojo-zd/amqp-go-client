package amqp

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	ml "loyocloud-infrastructure/tmodels"
	ct "loyocloud-notify-client/constants"

	"github.com/astaxie/beego/utils"
)

type AmqpMessageInterface interface {
	Analyse(amqp ml.AMQPMessage)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func ReceiveMessage(queueName string, amqpMessage AmqpMessageInterface) {
	conn, err := amqp.Dial(ct.AMQPUrl)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName,    // name
		ct.Durable,   // durable
		ct.Deleted,   // delete when usused
		ct.Exclusive, // exclusive
		ct.NoWait,    // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var amqp ml.AMQPMessage
			if err := json.Unmarshal(d.Body, &amqp); err != nil {
				utils.Display("从消息服务器获取到得数据", amqp)
			}
			amqpMessage.Analyse(amqp)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
