package test

import (
	"encoding/json"

	c "github.com/lianhong2758/RosmBot/ctx"
	"github.com/lianhong2758/RosmBot/web"
	"github.com/wdvxdr1123/ZeroBot/utils/helper"
)

func init() {
	//插件注册
	en := c.Register(&c.PluginData{ //插件英文索引
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
	en.AddRex(`^解析([\s\S]*)$`).Handle(func(ctx *c.CTX) {
		info := new(c.H)
		err := json.Unmarshal(helper.StringToBytes(ctx.Being.Rex[1]), info)
		if err != nil {
			ctx.Send(c.Text("解析失败", err))
			return
		}
		ctx.Send(c.MYContent(info))
	})
}
