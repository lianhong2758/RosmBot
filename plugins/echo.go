package plugins

import (
	c "github.com/lianhong2758/RosmBot/ctx"
)

func init() {
	en := c.Register("echo", &c.PluginData{
		Name: "复读",
		Help: "复读...",
	})
	en.AddRex(func(ctx *c.CTX) {
		ctx.Send(c.Text(ctx.Being.Rex[1]))
	}, "^复读(.*)")
}
