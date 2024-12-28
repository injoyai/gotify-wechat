package main

import (
	"github.com/eatmoreapple/openwechat"
	"github.com/gotify/plugin-api"
)

var _ plugin.Messenger = (*Plugin)(nil)

// GetGotifyPluginInfo returns gotify plugin info
func GetGotifyPluginInfo() plugin.Info {
	return plugin.Info{
		Version: "v1.0",
		Author:  "injoy",
		Name:    "Gotify微信插件",
	}
}

// NewGotifyPluginInstance creates a plugin instance for a user context.
func NewGotifyPluginInstance(ctx plugin.UserContext) plugin.Plugin {
	p := &Plugin{
		ctx: ctx,
	}
	p.initWechat()
	return p
}

// Plugin is plugin instance
type Plugin struct {
	ctx     plugin.UserContext
	enabled bool
	handler plugin.MessageHandler

	Client *openwechat.Bot
	Self   *openwechat.Self
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

func main() {
	panic("this should be built as go plugin")
}
