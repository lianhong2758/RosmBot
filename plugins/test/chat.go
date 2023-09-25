package test

import (
	c "github.com/lianhong2758/RosmBot/ctx"
	"github.com/lianhong2758/RosmBot/zero"
)

func init() {
	en := c.Register(c.NewRegist("@回复", "- @机器人", ""))
	en.AddWord("").SetBlock(true).Handle(func(ctx *c.CTX) {
		ctx.Send(c.Text(zero.MYSconfig.BotToken.BotName, "不在呢~"))
	})
}
