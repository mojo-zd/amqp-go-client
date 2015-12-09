package buzz

import (
	ct "constants"
	"jpush"
	ml "models"
)

func WorkflowExcutor(amqp ml.AMQPMessage) {
	// JPush发送通知到客户端
	jpush.PushMessage(amqp, ct.WorkflowType, workflowTitleConcat(&amqp))
}

func workflowTitleConcat(amqp *ml.AMQPMessage) string {
	title := ""
	switch {
	case amqp.OperationType == ct.Create:
		title = amqp.BuzzBody.SenderName + ct.WFSubmit + amqp.BuzzBody.Title
	case amqp.OperationType == ct.WillExpire:
		// 提示在时间段内办理
		//title = amqp.BuzzBody.SenderName + ct.WFSubmit + amqp.BuzzBody.Title
	case amqp.OperationType == ct.Accept:
		title = ct.WFYour + amqp.BuzzBody.Title + ct.WFAccept
	case amqp.OperationType == ct.Done:
		title = ct.WFYour + amqp.BuzzBody.Title + ct.WFDone
	case amqp.OperationType == ct.Reject:
		title = ct.WFYour + amqp.BuzzBody.Title + ct.WFFailed
	}

	return title
}
