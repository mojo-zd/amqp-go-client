package buzz

import (
	ct "constants"
	"jpush"
	ml "models"
)

func ReportExcutor(amqp ml.AMQPMessage) {
	// JPush发送通知到客户端
	jpush.PushMessage(amqp, ct.ReportType, reportTitleConcat(&amqp))
}

func reportTitleConcat(amqp *ml.AMQPMessage) string {
	title := ""
	switch {
	case amqp.OperationType == ct.Create:
		title = amqp.BuzzBody.SenderName + ct.ReportSubmit + ml.TimestampToDate(ct.YMDFormat, amqp.BuzzBody.CreateAt) + ct.ReportJob + amqp.BuzzBody.ReportType

	case amqp.OperationType == ct.Reviewed:
		title = amqp.BuzzBody.SenderName + ct.ReportReview + ml.TimestampToDate(ct.YMFormat, amqp.BuzzBody.CreateAt) + ct.ReportOfJob + amqp.BuzzBody.ReportType
	}

	return title
}
