package main

import (
	"github.com/eatmoreapple/openwechat"
	"io"
	"os"
)

var (
	Client = openwechat.DefaultBot(openwechat.Desktop)
	Self   *openwechat.Self
)

type Wechat struct {
	Self   *openwechat.Self
	Client *openwechat.Bot
}

func init() {

	// 注册登陆二维码回调
	Client.UUIDCallback = openwechat.PrintlnQrcodeUrl
	// 注册消息处理函数
	Client.MessageHandler = func(msg *openwechat.Message) {
		//接收私发消息
		if msg.IsSendByFriend() {
			//推送消息到gotify
		}
	}

	HotLoginFilename := "./hot_login"
	if stat, err := os.Stat(HotLoginFilename); stat == nil || os.IsNotExist(err) {
		os.Create(HotLoginFilename)
	}

	if err := Client.HotLogin(openwechat.NewFileHotReloadStorage(HotLoginFilename)); err != nil {
		if err != io.EOF {
			if err.Error() == "invalid storage" || err.Error() == "failed login check" {
				os.Remove(HotLoginFilename)
			}
			panic(err)
		}
		if err := Client.Login(); err != nil {
			panic(err)
		}
	}

	//获取当前用户信息
	var err error
	Self, err = Client.GetCurrentUser()
	if err != nil {
		panic(err)
	}

}
