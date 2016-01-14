package jpush

import (
	"github.com/ylywyn/jpush-api-go-client"
	"log"
	ml "loyocloud-infrastructure/tmodels"
	// "strconv"

	// "github.com/astaxie/beego/utils"
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

	notice.SetAndroidNotice(&jpushclient.AndroidNotice{Alert: title, Title: amqp.BuzzBody.NotifyBuzz.Title, BuilderId: 0, Extras: extras})
	notice.SetIOSNotice(&jpushclient.IOSNotice{Alert: title, Sound: "default", Badge: "+1", Extras: extras})

	options := jpushclient.Option{ApnsProduction: true}
	payload := jpushclient.NewPushPayLoad()
	payload.SetPlatform(&pf)
	payload.SetAudience(&ad)

	payload.SetNotice(&notice)
	payload.SetOptions(&options)

	bytes, _ := payload.ToBytes()
	log.Printf("%s\r\n", string(bytes))

	//push
	c := jpushclient.NewPushClient(ml.Secret, ml.AppKey)
	log.Printf("Secret %s AppKey %s\r\n", ml.Secret, ml.AppKey)
	str, err := c.Send(bytes)
	if err != nil {
		log.Printf("err:%s", err.Error())
	} else {
		log.Printf("ok:%s", str)
	}
}
