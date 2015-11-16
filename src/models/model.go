package model

import (
	"gopkg.in/mgo.v2/bson"
)

type AMQPMessage struct {
	StartTime     int64  `json:"startTime"`
	Repeat        bool   `json:"repeat"`
	Obj           string `json:"obj"`
	CronExpress   string `json:"cronExpress"`
	BuzzType      string `json:"buzzType"`
	OperationType string `json:"operationType"`
}

//报告
type WReport struct {
	Id        bson.ObjectId `bson:"_id" json:"id,omitempty"`              //id
	Title     string        `bson:"title" json:"title"`                   //标题
	Content   string        `bson:"content" json:"content"`               //内容
	Creator   *ShortUser    `bson:"creator" json:"creator"`               //创建者
	Reviewers []*Reviewer   `bson:"reviewers" json:"reviewers"`           //点评人
	Members   []*Member     `bson:"members" json:"members,omitempty"`     //报告相关的人（抄送人、AT的人）
	Type      int           `bson:"type" json:"type"`                     //报告类型
	BeginAt   int64         `bson:"begin_at" json:"beginAt"`              //报告开始时间
	EndAt     int64         `bson:"end_at" json:"endAt"`                  //报告结束时间
	ProjectId string        `bson:"projectId" json:"projectId,omitempty"` //项目ID
	CreatedAt int64         `bson:"created_at" json:"createdAt"`          //创建时间
	UpdatedAt int64         `bson:"updated_at" json:"updatedAt"`          //更新时间
	UUID      string        `bson:"uuid" json:"attachmentUUId"`           //附件GUID
	// Attachments []*tm.Attachment `bson:"-" json:"attachments,omitempty"`       //附件集合
	CompanyId bson.ObjectId `bson:"company_id" json:"-"`        //公司ID
	IsDelayed bool          `bson:"isDelayed" json:"isDelayed"` //是否补签
	// CRMDatas   []*CRMData      `bson:"-" json:"crmDatas,omitempty"` //客户管理系统获取的数据
	// Discussion base.Discussion `bson:"-" json:"discuss"` //讨论
}

//点评者
type Reviewer struct {
	User       *ShortUser `bson:"user" json:"user,omitempty"`
	Viewed     bool       `bson:"viewed" json:"viewed"`
	ViewedAt   int64      `bson:"viewedAt" json:"viewedAt"`
	Reviewed   bool       `bson:"reviewed" json:"reviewed"`         //是否点评
	ReviewedAt int64      `bson:"reviewedAt" json:"reviewedAt"`     //点评时间
	Score      int        `bson:"score" json:"score"`               //评分
	Comment    string     `bson:"comment" json:"comment,omitempty"` //内容
	Status     int        `bson:"status" json:"status"`             //点评状态
}

//简单用户类型
type ShortUser struct {
	Id       bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	RealName string        `bson:"realname" json:"name"`
	Avatar   string        `bson:"avatar" json:"avatar,omitempty"`
	// Depts    []*ShortUserDept `bson:"depts" json:"depts,omitempty"`
}

//成员(任务参与人、报告抄送人、通知公告接收人、客户参与人)
type Member struct {
	User     *ShortUser `bson:"user" json:"user,omitempty"`
	Viewed   bool       `bson:"viewed" json:"viewed"`
	ViewedAt int64      `bson:"viewedAt" json:"viewedAt"`
}
