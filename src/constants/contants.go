package constants

const (
	BuzzTask     = "task"
	BuzzReport   = "report"
	BuzzWorkflow = "workflow"

	Create     = "create"
	Finished   = "finished"
	Reviewed   = "reviewed"
	Update     = "update"
	Expire     = "expire"
	WillExpire = "willexpire"
	Redo       = "redo"
	Repeat     = "repeat"
	Done       = "done"
	Accept     = "accept"
	Reject     = "reject"
)

const (
	_ = iota // 忽略掉0 从1开始
	TaskType
	ReportType
	WorkflowType
)

const (
	AMQPUrl   = "amqp://guest:guest@localhost:5672/"
	QueueName = "mojo-work"
	Durable   = true
	Deleted   = false
	Exclusive = false
	NoWait    = false
)

// JPush
const (
	AppKey = "ee69a29f3ef8366586cabacb" // ios ee69a29f3ef8366586cabacb android 023bdeb58e5fefe8a2aeeaed
	Secret = "1a272189e01d6f5b42113bd2" //1a272189e01d6f5b42113bd2  android 516895c76543c79eac50603b
)

const (
	ReportSubmit = " 提交了 "
	ReportJob    = "工作"
	ReportReview = " 点评了你"
	ReportOfJob  = "工作的"

	TaskAssign     = " 分配了新的任务 "
	TaskSubAssign  = " 分配了子任务 "
	TaskFinished   = " 完成了任务 "
	TaskReview     = " 点评了任务 "
	TaskUpdate     = " 修改了任务 "
	TaskPrefix     = "任务 "
	TaskLeftTime   = " 还剩"
	TaskOver       = "到期，"
	TaskWillExpire = " 请抓紧时间哦。"
	TaskExpire     = " 已经超时"
	TaskRedo       = " 被打回重做"
	TaskRepeat     = "重复任务："
	TaskDone       = " 记得完成哦!"

	WFSubmit = " 向你提交了审批申请 "
	WFNeed   = "该审批需在"
	WFDo     = "内办理"
	WFYour   = "你的审批申请 "
	WFAccept = " 已办结"
	WFDone   = " 已通过"
	WFFailed = " 未通过"
)

const (
	YMDFormat = "2006-01-02" // 时间格式 '年-月-日'
	YMFormat  = "2006-01"    // 时间格式 '年-月'
)
