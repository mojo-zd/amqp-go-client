package main

import (
	"github.com/astaxie/beego/utils"
	ml "loyocloud-infrastructure/tmodels"
	bizType "loyocloud-infrastructure/tmodels/enum/biztype"
	"loyocloud-notify-client/amqp"
	"loyocloud-notify-client/buzz"
)

type AmqpNotify struct {
}

func (this AmqpNotify) Analyse(amqp ml.AMQPMessage) {
	utils.Display("从消息服务器接收到的数据", amqp)
	// 如果没有接受对象则不进行消息发送
	if amqp.BuzzBody.NotifyBuzz.Receivers == nil {
		return
	}

	switch {
	case bizType.Task == amqp.BuzzType:
		buzz.TaskExcutor(amqp)
	case bizType.WorkReport == amqp.BuzzType:
		buzz.ReportExcutor(amqp)
	case bizType.Approval == amqp.BuzzType:
		buzz.WorkflowExcutor(amqp)
	default:

	}
}

func main() {
	notify := AmqpNotify{}
	amqp.ReceiveMessage(ml.NotifyQueueName, notify)
}
