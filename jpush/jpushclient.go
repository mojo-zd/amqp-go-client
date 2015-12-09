package jpush

import (
	"github.com/ylywyn/jpush-api-go-client"
	"log"
	ml "loyocloud-infrastructure/tmodels"
)

func PushMessage(amqp ml.AMQPMessage, btype int, title string) {

	var pf jpushclient.Platform
	pf.Add(jpushclient.ANDROID)
	pf.Add(jpushclient.IOS)

	//Audience
	var ad jpushclient.Audience
	ad.SetAlias(amqp.BuzzBody.NotifyBuzz.Receivers)
	// ad.All()

	//Notice
	var notice jpushclient.Notice
	extras := make(map[string]interface{})
	extras["buzzId"] = amqp.BuzzBody.NotifyBuzz.BuzzId
	extras["buzzType"] = btype
	extras["operationType"] = amqp.OperationType

<<<<<<< HEAD:jpush/jpushclient.go
	notice.SetAndroidNotice(&jpushclient.AndroidNotice{Alert: title, Title: amqp.BuzzBody.NotifyBuzz.Title, BuilderId: +1, Extras: extras})
	notice.SetIOSNotice(&jpushclient.IOSNotice{Alert: title, Sound: "default", Badge: +1, Extras: extras})
=======
	notice.SetAndroidNotice(&jpushclient.AndroidNotice{Alert: title, Title: amqp.BuzzBody.Title, BuilderId: +1, Extras: extras})
	notice.SetIOSNotice(&jpushclient.IOSNotice{Alert: title, Sound: "default", Badge: +1, Extras: extras})
	// notice.SetWinPhoneNotice(&jpushclient.WinPhoneNotice{Alert: "WinPhoneNotice"})
>>>>>>> d3d6b3ae4f8438c0065d46605d5ccd2de458f7e2:src/jpush/jpushclient.go

	var msg jpushclient.Message
	msg.Title = amqp.BuzzBody.NotifyBuzz.Title
	msg.Content = "添加operation内容"

	payload := jpushclient.NewPushPayLoad()
	payload.SetPlatform(&pf)
	payload.SetAudience(&ad)
	payload.SetMessage(&msg)
	payload.SetNotice(&notice)

	bytes, _ := payload.ToBytes()
	log.Printf("%s\r\n", string(bytes))

	//push
	c := jpushclient.NewPushClient(ml.Secret, ml.AppKey)
	str, err := c.Send(bytes)
	if err != nil {
		log.Printf("err:%s", err.Error())
	} else {
		log.Printf("ok:%s", str)
	}
}
