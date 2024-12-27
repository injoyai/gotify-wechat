package main

import (
	"github.com/eatmoreapple/openwechat"
	"github.com/gotify/plugin-api"
)

var _ plugin.Plugin = (*Plugin)(nil)

// GetGotifyPluginInfo returns gotify plugin info
func GetGotifyPluginInfo() plugin.Info {
	return plugin.Info{
		Name:       "Gotify微信插件",
		ModulePath: "github.com/gotify/server/example/minimal",
	}
}

// NewGotifyPluginInstance creates a plugin instance for a user context.
func NewGotifyPluginInstance(ctx plugin.UserContext) plugin.Plugin {
	return &Plugin{
		Wechat: &Wechat{},
	}
}

// Plugin is plugin instance
type Plugin struct {
	ctx     plugin.UserContext
	enabled bool
	handler plugin.MessageHandler

	*Wechat
}

// Enable implements plugin.Plugin
func (c *Plugin) Enable() error {
	c.enabled = true
	return nil
}

// Disable implements plugin.Disable
func (c *Plugin) Disable() error {
	c.enabled = false
	return nil
}

// SetMessageHandler implements plugin.Messenger
func (c *Plugin) SetMessageHandler(h plugin.MessageHandler) {
	c.handler = h
}

func (c *Plugin) SendMessage(msg *openwechat.Message) error {
	if c.handler != nil {
		return c.handler.SendMessage(plugin.Message{
			Message:  msg.Content,
			Title:    "来自微信好友:" + msg.FromUserName,
			Priority: 0,
			Extras: map[string]interface{}{
				"微信好友": msg.FromUserName,
			},
		})
	}
	return nil
}
