package test

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
	en.AddRex(func(ctx *c.CTX) {
		ctx.Send(c.ImageWithText(ctx.Being.Rex[1], 0, 0, 0, ctx.Being.Rex[2]))
	}, "^复图(.*)文字(.*)")
	en.AddRex(func(ctx *c.CTX) {
		ctx.Send(c.Image(ctx.Being.Rex[1], 0, 0, 0))
	}, "^复纯图(.*)")
}
