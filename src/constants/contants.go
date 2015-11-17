package constants

const (
	BuzzTask     = "task"
	BuzzReport   = "report"
	BuzzWorkflow = "workflow"

	CreateOperation   = "create"
	FinishedOperation = "finished"
	ReviewedOperation = "reviewed"
	AcceptOperation   = "accept"
	RejectOperation   = "reject"
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
