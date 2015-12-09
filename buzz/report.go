package buzz

import (
	ml "loyocloud-infrastructure/tmodels"
	util "loyocloud-infrastructure/utils"
	"loyocloud-notify-client/jpush"
)

func ReportExcutor(amqp ml.AMQPMessage) {
	// JPush发送通知到客户端
	jpush.PushMessage(amqp, ml.ReportType, reportTitleConcat(&amqp))
}

func reportTitleConcat(amqp *ml.AMQPMessage) string {
	title := ""
	switch {
	case amqp.OperationType == ml.Create:
		title = amqp.BuzzBody.NotifyBuzz.SenderName + ml.ReportSubmit + util.TimestampToDate(ml.YMDFormat, amqp.BuzzBody.NotifyBuzz.CreateAt) + ml.ReportJob + amqp.BuzzBody.NotifyBuzz.ReportType

	case amqp.OperationType == ml.Reviewed:
		title = amqp.BuzzBody.NotifyBuzz.SenderName + ml.ReportReview + util.TimestampToDate(ml.YMFormat, amqp.BuzzBody.NotifyBuzz.CreateAt) + ml.ReportOfJob + amqp.BuzzBody.NotifyBuzz.ReportType
	}

	return title
}
