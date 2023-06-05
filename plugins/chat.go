package plugins

import (
	c "github.com/lianhong2758/RosmBot/ctx"
	"github.com/lianhong2758/RosmBot/zero"
)

func init() {
	en := c.Register("chat", &c.PluginData{
		Name: "@回复",
		Help: "@机器人",
	})
	en.AddWord(func(ctx *c.CTX) {
		ctx.Send(c.Text(zero.BotToken.BotName, "不在呢~"))
	}, "")
}
