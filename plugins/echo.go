package plugins

import (
	c "github.com/lianhong2758/RosmBot/ctx"
)

func init() {
	en := c.Register("echo", &c.PluginData{
		Name: "复读",
		Help: "复读...",
	})
	en.AddWord(func(ctx *c.CTX) {
		ctx.Send(c.Text(ctx.Mess.Content.Text[17:]))
	}, "复读")
}
