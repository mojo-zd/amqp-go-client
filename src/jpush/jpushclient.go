package jpush

import (
	ct "constants"
	"github.com/ylywyn/jpush-api-go-client"
	"log"
	ml "models"
)

func PushMessage(amqp ml.AMQPMessage, btype int) {

	var pf jpushclient.Platform
	pf.All()

	//Audience
	var ad jpushclient.Audience
	ad.SetAlias(amqp.BuzzBody.Receivers)
	// ad.All()

	//Notice
	var notice jpushclient.Notice
	extras := make(map[string]interface{})
	extras["buzzId"] = amqp.BuzzBody.BuzzId
	extras["buzzType"] = btype
	extras["operationType"] = amqp.OperationType

	notice.SetAlert("alert_test")
	notice.SetAndroidNotice(&jpushclient.AndroidNotice{Alert: "谷歌服务", Title: amqp.BuzzBody.Title, BuilderId: 1, Extras: extras})
	notice.SetIOSNotice(&jpushclient.IOSNotice{Alert: "苹果一打", Sound: "default", Badge: 4, Extras: extras})
	notice.SetWinPhoneNotice(&jpushclient.WinPhoneNotice{Alert: "WinPhoneNotice"})

	var msg jpushclient.Message
	msg.Title = amqp.BuzzBody.Title
	msg.Content = "添加operation内容"

	payload := jpushclient.NewPushPayLoad()
	payload.SetPlatform(&pf)
	payload.SetAudience(&ad)
	payload.SetMessage(&msg)
	payload.SetNotice(&notice)

	bytes, _ := payload.ToBytes()
	log.Printf("%s\r\n", string(bytes))

	//push
	c := jpushclient.NewPushClient(ct.Secret, ct.AppKey)
	str, err := c.Send(bytes)
	if err != nil {
		log.Printf("err:%s", err.Error())
	} else {
		log.Printf("ok:%s", str)
	}
}
