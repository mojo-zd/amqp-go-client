package buzz

import (
	ct "constants"
	"jpush"
	ml "models"
)

func TaskExcutor(amqp ml.AMQPMessage) {
	// JPush发送通知到客户端
	jpush.PushMessage(amqp, ct.TaskType, taskTitleConcat(&amqp))
}

func taskTitleConcat(amqp *ml.AMQPMessage) string {
	title := ""
	switch {
	// 分配任务
	case amqp.OperationType == ct.Create:
		if amqp.BuzzBody.SubTask {
			title = amqp.BuzzBody.SenderName + ct.TaskSubAssign + amqp.BuzzBody.Title
		} else {
			title = amqp.BuzzBody.SenderName + ct.TaskAssign + amqp.BuzzBody.Title
		}
	// 完成任务
	case amqp.OperationType == ct.Finished:
		title = amqp.BuzzBody.SenderName + ct.TaskFinished + amqp.BuzzBody.Title
	// 点评任务
	case amqp.OperationType == ct.Reviewed:
		title = amqp.BuzzBody.SenderName + ct.TaskReview + amqp.BuzzBody.Title
	// 修改了任务
	case amqp.OperationType == ct.Update:
		title = amqp.BuzzBody.SenderName + ct.TaskUpdate + amqp.BuzzBody.Title
	//即将过期
	case amqp.OperationType == ct.WillExpire:
		title = ct.TaskPrefix + amqp.BuzzBody.Title + ct.TaskLeftTime + "1小时" + ct.TaskOver + ct.TaskWillExpire
	//已经过期
	case amqp.OperationType == ct.Expire:
		title = ct.TaskPrefix + amqp.BuzzBody.Title + ct.TaskExpire
	//打回重做
	case amqp.OperationType == ct.Redo:
		title = ct.TaskPrefix + amqp.BuzzBody.Title + ct.TaskRedo
	//重复任务
	case amqp.OperationType == ct.Repeat:
		title = ct.TaskRepeat + amqp.BuzzBody.Title + ct.TaskDone
	}

	return title
}
