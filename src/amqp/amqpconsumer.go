package amqp

import (
	"buzz"
	ct "constants"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	ml "models"

	"github.com/astaxie/beego/utils"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func ReceiveMessage() {
	conn, err := amqp.Dial(ct.AMQPUrl)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		ct.QueueName, // name
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
			analyze(amqp)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func analyze(amqp ml.AMQPMessage) {

	switch {
	case ct.BuzzTask == amqp.BuzzType:
		buzz.TaskExcutor(amqp)
	case ct.BuzzReport == amqp.BuzzType:
		buzz.ReportExcutor(amqp)
	case ct.BuzzWorkflow == amqp.BuzzType:
		//buzz.WorkflowExcutor(buzzBody, amqp.OperationType, amqp.BuzzType)
	default:

	}
}
