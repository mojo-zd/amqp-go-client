package amqp

import (
	"encoding/json"
	"github.com/astaxie/beego/utils"
	"github.com/streadway/amqp"
	"log"
	ml "loyocloud-infrastructure/tmodels"
	util "loyocloud-infrastructure/utils"
)

type AmqpMessageInterface interface {
	Analyse(amqp ml.AMQPMessage)
}

func ReceiveMessage(queueName string, amqpMessage AmqpMessageInterface) {
	//ml.DeptAndChildDept
	conn, err := amqp.Dial(ml.AMQPUrl)
	util.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	util.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName,    // name
		ml.Durable,   // durable
		ml.Deleted,   // delete when usused
		ml.Exclusive, // exclusive
		ml.NoWait,    // no-wait
		nil,          // arguments
	)
	util.FailOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	util.FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var amqp ml.AMQPMessage
			if err := json.Unmarshal(d.Body, &amqp); err != nil {
				utils.Display("从消息服务器获取到得数据", amqp)
			}
			amqpMessage.Analyse(amqp)
			//analyze(amqp)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
