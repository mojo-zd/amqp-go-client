package buzz

import (
	ml "loyocloud-infrastructure/tmodels"
	"loyocloud-notify-client/jpush"
)

func WorkflowExcutor(amqp ml.AMQPMessage) {
	// JPush发送通知到客户端
	jpush.PushMessage(amqp, ml.WorkflowType, workflowTitleConcat(&amqp))
}

func workflowTitleConcat(amqp *ml.AMQPMessage) string {
	title := ""
	switch {
	case amqp.OperationType == ml.Create:
		title = amqp.BuzzBody.NotifyBuzz.SenderName + ml.WFSubmit + amqp.BuzzBody.NotifyBuzz.Title
	case amqp.OperationType == ml.WillExpire:
		// 提示在时间段内办理
		//title = amqp.BuzzBody.SenderName + ml.WFSubmit + amqp.BuzzBody.Title
	case amqp.OperationType == ml.Accept:
		title = ml.WFYour + amqp.BuzzBody.NotifyBuzz.Title + ml.WFAccept
	case amqp.OperationType == ml.Done:
		title = ml.WFYour + amqp.BuzzBody.NotifyBuzz.Title + ml.WFDone
	case amqp.OperationType == ml.Reject:
		title = ml.WFYour + amqp.BuzzBody.NotifyBuzz.Title + ml.WFFailed
	}

	return title
}
