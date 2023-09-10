package test

import (
	c "github.com/lianhong2758/RosmBot/ctx"
	"github.com/lianhong2758/RosmBot/web"
)

func init() {
	//插件注册
	en := c.Register("echo", &c.PluginData{ //插件英文索引
		Name: "复读",      //中文插件名
		Help: "- 复读...", //插件帮助
	})
	en.AddRex("^复读(.*)").SetBlock(true).Rule(func(ctx *c.CTX) bool { return true }, c.OnlyMaster).Handle(func(ctx *c.CTX) { //正则的触发方式
		ctx.Send(c.Text(ctx.Being.Rex[1])) //发送文字信息
	})
	en.AddRex("^复图(.*)文字(.*)").Handle(func(ctx *c.CTX) {
		con, _ := web.URLToConfig(ctx.Being.Rex[1])
		ctx.Send(c.ImageUrlWithText(web.UpImgUrl(ctx.Being.Rex[1]), con.Width, con.Height, 0, ctx.Being.Rex[2]))
	})
	en.AddRex("^复纯图(.*)").Handle(func(ctx *c.CTX) {
		con, _ := web.URLToConfig(ctx.Being.Rex[1])
		ctx.Send(c.ImageUrl(web.UpImgUrl(ctx.Being.Rex[1]), con.Width, con.Height, 0))
	})
}
