package test

import (
	"encoding/json"

	c "github.com/lianhong2758/RosmBot/ctx"
	"github.com/lianhong2758/RosmBot/web"
)

const url = "http://8.134.179.136/vtbwife?id="

func init() { // 插件主体
	en := c.Register(&c.PluginData{
		Name: "抽vtb老婆",
		Help: "- /抽vtb(老婆)",
	})
	en.AddRex(`^/抽(vtb|VTB)(老婆)?$`).Handle(func(ctx *c.CTX) {
		body, err := web.GetData(url+ctx.Being.User.ID, "")
		if err != nil {
			ctx.Send(c.Text("ERROR: ", err))
			return
		}
		var r result
		err = json.Unmarshal(body, &r)
		if err != nil {
			ctx.Send(c.Text("ERROR: ", err))
			return
		}
		con, _ := web.URLToConfig(r.Imgurl)
		ctx.Send(ctx.AT(ctx.Being.User.ID), c.ImageUrlWithText(web.UpImgUrl(r.Imgurl), con.Width, con.Height, 0, "\n今天你的VTB老婆是: "+r.Name))
		ctx.Send(c.Text(r.Message))
	})
}

type result struct {
	Code    int    `json:"code"`
	Imgurl  string `json:"imgurl"`
	Name    string `json:"name"`
	Message string `json:"message"`
}
