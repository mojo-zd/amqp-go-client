package buzz

import (
	ct "constants"
	"jpush"
	ml "models"
)

func ReportExcutor(amqp ml.AMQPMessage) {
	// JPush发送通知到客户端
	jpush.PushMessage(amqp, ct.ReportType)
}
