package contoller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/message"
	"test/models"
)

func WxApi(c *gin.Context) {

	cfg := models.GetCfg()

	appId := cfg.Section("").Key("wx.app.id").String()
	appSecret := cfg.Section("").Key("wx.app.secret").String()
	token := cfg.Section("").Key("wx.token").String()
	aesKey := cfg.Section("").Key("wx.encoding.aes.key").String()


	//配置微信参数
	config := &wechat.Config{
		AppID:          appId,
		AppSecret:      appSecret,
		Token:          token,
		EncodingAESKey: aesKey,
	}
	wc := wechat.NewWechat(config)

	// 传入request和responseWriter
	server := wc.GetServer(c.Request, c.Writer)
	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {

		//回复消息：演示回复用户发送的消息
		text := message.NewText(msg.Content)
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	})

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}
	//发送回复的消息
	_ = server.Send()

}
