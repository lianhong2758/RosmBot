package test

import (
	"time"

	c "github.com/lianhong2758/RosmBot/ctx"
)

func init() {
	en := c.Register("test", &c.PluginData{
		Name: "测试",
		Help: "- 测试\n" +
			"- 测试下标\n" +
			"- 测试预览\n" +
			"- 测试全体next\n" +
			"- 测试个人next\n" +
			"- 测试表情",
	})
	en.AddWord(func(ctx *c.CTX) {
		ctx.Send(c.Text("你好"), ctx.AT("76257069"), c.Link("www.baidu.com", false, "百度一下"), c.ATAll(), ctx.RoomLink("23648"), c.Text("[爱心]"))
	}, "测试")
	en.AddWord(func(ctx *c.CTX) {
		s := c.BadgeStr{
			Icon: "http://8.134.179.136/favicon.ico",
			Text: "清雪官方",
			URL:  "http://8.134.179.136",
		}
		ctx.Send(c.Text("清雪官网~"), c.Badge(s))
	}, "测试下标")
	en.AddWord(func(ctx *c.CTX) {
		s := c.PreviewStr{
			Icon:       "http://8.134.179.136/favicon.ico",
			URL:        "http://8.134.179.136",
			ImageURL:   "http://8.134.179.136/ippic",
			IsIntLink:  true,
			SourceName: "我是喵喵喵~",
			Title:      "这是一个标题测试",
			Content:    "我是具体内容",
		}
		ctx.Send(c.Text("测试"), c.Preview(s))
	}, "测试预览")
	en.AddWord(func(ctx *c.CTX) {
		next, stop := ctx.GetNextAllMess()
		defer stop()
		ctx.Send(c.Text("测试开始"))
		for i := 0; i < 3; i++ {
			select {
			case <-time.After(time.Second * 60):
				ctx.Send(c.Text("时间太久了"))
				return
			case ctx2 := <-next:
				ctx.Send(c.Text("这是全体下一句话:", ctx2.Being.Word))
			}
		}
	}, "测试全体next")
	en.AddWord(func(ctx *c.CTX) {
		next, stop := ctx.GetNextUserMess()
		defer stop()
		ctx.Send(c.Text("测试开始"))
		for i := 0; i < 3; i++ {
			select {
			case <-time.After(time.Second * 60):
				ctx.Send(c.Text("时间太久了"))
				return
			case ctx2 := <-next:
				ctx.Send(c.Text("这是个人下一句话:", ctx2.Being.Word))
			}
		}
	}, "测试个人next")
	en.AddWord(func(ctx *c.CTX) {
		result := ctx.Send(c.Text("测试开始,表态此条消息"))
		next, stop := ctx.GetNextAllEmoticon(result.Data.BotMsgID)
		defer stop()
		for i := 0; i < 3; i++ {
			select {
			case <-time.After(time.Second * 60):
				ctx.Send(c.Text("时间太久了"))
				return
			case ctx2 := <-next:
				ctx.Send(c.Text("这是表态结果:\n", ctx2.Event.AddQuickEmoticon))
			}
		}
	}, "测试表情")
}

//ctx有消息的全部信息,ctx.Being有简单的消息信息获取
//ctx.ChangeRoome可以改变发送的消息房间等
//ctx.Send(...)是发送消息的基本格式
