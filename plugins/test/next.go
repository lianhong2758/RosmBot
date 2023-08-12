package test

import (
	"time"

	c "github.com/lianhong2758/RosmBot/ctx"
)

func init() {
	en := c.Register("next", &c.PluginData{
		Name: "连续对话测试",
		Help: "- 测试next",
	})
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
}
