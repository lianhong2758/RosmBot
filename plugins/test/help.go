package test

import (
	"strings"

	c "github.com/lianhong2758/RosmBot/ctx"
)

func init() {
	en := c.Register("help", &c.PluginData{
		Name: "帮助菜单",
		Help: "help",
	})
	en.AddWord(func(ctx *c.CTX) {
		var msg strings.Builder
		msg.WriteString("*****菜单********")
		for _, v := range c.GetPlugins() {
			msg.WriteString("\n")
			msg.WriteString("#")
			msg.WriteString(v.Name)
			msg.WriteString("\n")
			msg.WriteString(v.Help)
			msg.WriteString("\n")
		}
		msg.WriteString("*****************")
		ctx.Send(c.Text(msg.String()))
	}, "help", "帮助")
}
