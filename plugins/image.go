package plugins

import (
	"net/http"

	c "github.com/lianhong2758/RosmBot/ctx"
	"github.com/lianhong2758/RosmBot/web"
)

func init() {
	en := c.Register("image", &c.PluginData{
		Name: "图片",
		Help: "随机壁纸" + "兽耳" + "星空" + "白毛" + "我要涩涩" + "涩涩达咩" + "白丝" + "黑丝" + "丝袜" + "随机表情包" + "cos" + "盲盒" + "开盲盒",
	})
	en.AddWord(func(ctx *c.CTX) {
		path, err := web.RequestDataWith(&http.Client{}, "http://127.0.0.1/image/"+ctx.Mess.Content.Text[11:], http.MethodPost, "", "", nil)
		if err != nil {
			ctx.Send(c.Text("获取图片失败"))
		}
		ctx.Send(c.Image("http://8.134.179.136/file/"+string(path), 0, 0, 0))
	}, "随机壁纸", "兽耳", "星空", "白毛", "我要涩涩", "涩涩达咩", "白丝", "黑丝", "丝袜", "随机表情包", "cos", "盲盒", "开盲盒")
}
