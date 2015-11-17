package models

type AMQPMessage struct {
	StartTime     int64        `json:"startTime"`
	Repeat        bool         `json:"repeat"`
	BuzzBody      AMQPBuzzBody `json:"buzzBody"`
	CronExpress   string       `json:"cronExpress"`
	BuzzType      string       `json:"buzzType"`
	OperationType string       `json:"operationType"`
}

// 业务数据包装
type AMQPBuzzBody struct {
	Receivers []string `json:"receivers"`
	Sender    string   `json:"sender"`
	Title     string   `json:"title"`
	BuzzId    string   `json:"buzzId"`
}
