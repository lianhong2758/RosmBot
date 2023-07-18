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
	en.AddRex(func(ctx *c.CTX) { //正则的触发方式
		ctx.Send(c.Text(ctx.Being.Rex[1])) //发送文字信息
	}, "^复读(.*)") //正则
	en.AddRex(func(ctx *c.CTX) {
		ctx.Send(c.ImageUrlWithText(web.UpImgUrl(ctx.Being.Rex[1]), 0, 0, 0, ctx.Being.Rex[2]))
	}, "^复图(.*)文字(.*)")
	en.AddRex(func(ctx *c.CTX) {
		ctx.Send(c.ImageUrl(web.UpImgUrl(ctx.Being.Rex[1]), 0, 0, 0))
	}, "^复纯图(.*)")
}
