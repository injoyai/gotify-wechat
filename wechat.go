package main

import (
	"github.com/eatmoreapple/openwechat"
	"github.com/gotify/plugin-api"
	"io"
	"os"
)

func (this *Plugin) initWechat() {
	this.Client = openwechat.DefaultBot(openwechat.Desktop)
	// 注册登陆二维码回调
	this.Client.UUIDCallback = openwechat.PrintlnQrcodeUrl
	// 注册消息处理函数
	this.Client.MessageHandler = this.OnMessage

	HotLoginFilename := "./config/hot_login"
	if stat, err := os.Stat(HotLoginFilename); stat == nil || os.IsNotExist(err) {
		os.Create(HotLoginFilename)
	}

	if err := this.Client.HotLogin(openwechat.NewFileHotReloadStorage(HotLoginFilename)); err != nil {
		if err != io.EOF {
			if err.Error() == "invalid storage" || err.Error() == "failed login check" {
				os.Remove(HotLoginFilename)
			}
			panic(err)
		}
		if err := this.Client.Login(); err != nil {
			panic(err)
		}
	}

	//获取当前用户信息
	var err error
	this.Self, err = this.Client.GetCurrentUser()
	if err != nil {
		panic(err)
	}

}

func (this *Plugin) OnMessage(msg *openwechat.Message) {
	//接收私发消息
	if msg.IsSendByFriend() {
		//推送消息到gotify
		if this.handler != nil {
			this.handler.SendMessage(plugin.Message{
				Message:  msg.Content,
				Title:    "来自微信好友:" + msg.FromUserName,
				Priority: 0,
				Extras: map[string]interface{}{
					"微信好友": msg.FromUserName,
				},
			})
		}
	}
}
