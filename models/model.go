// package models

// import (
// 	"time"
// )

// type AMQPMessage struct {
// 	StartTime     int64        `json:"startTime"`
// 	Repeat        bool         `json:"repeat"`
// 	BuzzBody      AMQPBuzzBody `json:"buzzBody"`
// 	CronExpress   string       `json:"cronExpress"`
// 	BuzzType      string       `json:"buzzType"`
// 	OperationType string       `json:"operationType"`
// }

// // 业务数据包装
// type AMQPBuzzBody struct {
// 	Receivers  []string `json:"receivers"`
// 	Sender     string   `json:"sender"`
// 	SenderName string   `json:"senderName"`
// 	ReportType string   `json:"reportType"`
// 	SubTask    bool     `json:"subTask"`
// 	Title      string   `json:"title"`
// 	CreateAt   int64    `json:"createAt"`
// 	StartAt    int64    `json:"StartAt"`
// 	EndAt      int64    `json:"EndAt"`
// 	BuzzId     string   `json:"buzzId"`
// }

// func TimestampToDate(format string, timestamp int64) string {
// 	return time.Unix(timestamp, 0).Format(format)
// }
