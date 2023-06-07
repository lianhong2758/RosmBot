package test

import (
	c "github.com/lianhong2758/RosmBot/ctx"
)

func init() {
	en := c.Register("test", &c.PluginData{
		Name: "测试",
		Help: "测试",
	})
	en.AddWord(func(ctx *c.CTX) {
		ctx.Send(c.Text("你好"), ctx.AT(76257069), c.Link("www.baidu.com", "百度一下"), c.ATAll(), ctx.RoomLink("23648"), c.Text("[爱心]"))
	}, "测试")
}
