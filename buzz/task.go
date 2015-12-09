package buzz

import (
	ml "loyocloud-infrastructure/tmodels"
	"loyocloud-notify-client/jpush"
)

func TaskExcutor(amqp ml.AMQPMessage) {
	// JPush发送通知到客户端
	jpush.PushMessage(amqp, ml.TaskType, taskTitleConcat(&amqp))
}

func taskTitleConcat(amqp *ml.AMQPMessage) string {
	title := ""
	switch {
	// 分配任务
	case amqp.OperationType == ml.Create:
		if amqp.BuzzBody.NotifyBuzz.SubTask {
			title = amqp.BuzzBody.NotifyBuzz.SenderName + ml.TaskSubAssign + amqp.BuzzBody.NotifyBuzz.Title
		} else {
			title = amqp.BuzzBody.NotifyBuzz.SenderName + ml.TaskAssign + amqp.BuzzBody.NotifyBuzz.Title
		}
	// 完成任务
	case amqp.OperationType == ml.Finished:
		title = amqp.BuzzBody.NotifyBuzz.SenderName + ml.TaskFinished + amqp.BuzzBody.NotifyBuzz.Title
	// 点评任务
	case amqp.OperationType == ml.Reviewed:
		title = amqp.BuzzBody.NotifyBuzz.SenderName + ml.TaskReview + amqp.BuzzBody.NotifyBuzz.Title
	// 修改了任务
	case amqp.OperationType == ml.Update:
		title = amqp.BuzzBody.NotifyBuzz.SenderName + ml.TaskUpdate + amqp.BuzzBody.NotifyBuzz.Title
	//即将过期
	case amqp.OperationType == ml.WillExpire:
		title = ml.TaskPrefix + amqp.BuzzBody.NotifyBuzz.Title + ml.TaskLeftTime + "1小时" + ml.TaskOver + ml.TaskWillExpire
	//已经过期
	case amqp.OperationType == ml.Expire:
		title = ml.TaskPrefix + amqp.BuzzBody.NotifyBuzz.Title + ml.TaskExpire
	//打回重做
	case amqp.OperationType == ml.Redo:
		title = ml.TaskPrefix + amqp.BuzzBody.NotifyBuzz.Title + ml.TaskRedo
	//重复任务
	case amqp.OperationType == ml.Repeat:
		title = ml.TaskRepeat + amqp.BuzzBody.NotifyBuzz.Title + ml.TaskDone
	}

	return title
}
