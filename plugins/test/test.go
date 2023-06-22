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
		ctx.Send(c.Text("你好"), ctx.AT("76257069"), c.Link("www.baidu.com", "百度一下"), c.ATAll(), ctx.RoomLink("23648"), c.Text("[爱心]"))
	}, "测试")
}

//ctx有消息的全部信息,ctx.Being有简单的消息信息获取
//ctx.ChangeRoome可以改变发送的消息房间等
//ctx.Send(...)是发送消息的基本格式
