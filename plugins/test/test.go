package test

import (
	"time"

	c "github.com/lianhong2758/RosmBot/ctx"
)

func init() {
	en := c.Register(&c.PluginData{
		Name: "测试",
		Help: "- 测试\n" +
			"- 测试下标\n" +
			"- 测试预览\n" +
			"- 测试全体next\n" +
			"- 测试个人next\n" +
			"- 测试表情",
	})
	en.AddWord("测试").Handle(func(ctx *c.CTX) {
		ctx.Send(c.Text("你好"), ctx.AT("76257069"), c.Link("www.baidu.com", false, "百度一下"), ctx.RoomLink("23648"), c.Text("[爱心]"))
	})
	en.AddWord("测试下标跳转房间").Handle(func(ctx *c.CTX) {
		s := c.BadgeStr{
			Icon: "http://8.134.179.136/favicon.ico",
			Text: "10248",
			URL:  "https://dby.miyoushe.com/chat/463/10248",
		}
		ctx.Send(c.Text("大别野房间~"), c.Badge(s))
	})
	en.AddWord("测试下标").Handle(func(ctx *c.CTX) {
		s := c.BadgeStr{
			Icon: "http://8.134.179.136/favicon.ico",
			Text: "清雪官方",
			URL:  "http://8.134.179.136",
		}
		ctx.Send(c.Text("清雪官网~"), c.Badge(s))
	})
	en.AddWord("测试预览").Handle(func(ctx *c.CTX) {
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
	})
	en.AddWord("测试全体next").Handle(func(ctx *c.CTX) {
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
	})
	en.AddWord("测试个人next").Handle(func(ctx *c.CTX) {
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
	})
	en.AddWord("测试表情").Handle(func(ctx *c.CTX) {
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
	})
	en.AddWord("测试视频").Handle(func(ctx *c.CTX) {
		ctx.Send(c.Text("测试开始"))
		s := c.PreviewStr{
			Icon:       "http://8.134.179.136/favicon.ico",
			URL:        "http://8.134.179.136/file?path=CSGO/1.mp4",
			IsIntLink:  true,
			SourceName: "清雪API",
			Title:      "测试视频",
			Content:    "CSGO精彩击杀,完美竞技平台",
		}
		ctx.Send(c.Text("视频测试"), c.Preview(s))
	})
	en.AddWord("测试组合").Handle(func(ctx *c.CTX) {
		ctx.Send(c.Text("测试开始"))
		s := c.PreviewStr{
			Icon:       "http://8.134.179.136/favicon.ico",
			URL:        "http://8.134.179.136/file?path=CSGO/1.mp4",
			IsIntLink:  true,
			SourceName: "清雪API",
			Title:      "测试视频",
			Content:    "CSGO精彩击杀,完美竞技平台",
		}
		ss := c.BadgeStr{
			Icon: "http://8.134.179.136/favicon.ico",
			Text: "清雪官方",
			URL:  "http://8.134.179.136",
		}
		ctx.Send(c.Text("视频测试"), c.Preview(s), c.Badge(ss))
	})
}

//ctx有消息的全部信息,ctx.Being有简单的消息信息获取
//ctx.ChangeRoome可以改变发送的消息房间等
//ctx.Send(...)是发送消息的基本格式
