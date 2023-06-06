package test

import (
	"os"
	"strconv"

	c "github.com/lianhong2758/RosmBot/ctx"
)

func init() {
	en := c.Register("setting", &c.PluginData{
		Name:       "设置",
		Help:       "设置欢迎房间xxx",
		DataFolder: "setting",
	})
	en.AddRex(func(ctx *c.CTX) {
		key := ctx.Being.Rex[1]
		if key == "" {
			ctx.Send(c.Text("输入房间号错误"))
			return
		}
		_, err := strconv.Atoi(key)
		if err != nil {
			ctx.Send(c.Text("输入房间号错误"))
			return
		}
		f, err := os.Create(en.DataFolder + "welcome.txt")
		if err != nil {
			ctx.Send(c.Text("ERROR: ", err))
			return
		}
		defer f.Close()
		_, err = f.WriteString(key)
		if err != nil {
			ctx.Send(c.Text("ERROR: ", err))
			return
		}
		ctx.Send(c.Text("设置成功"))
	}, "^设置欢迎房间(.*)")
}
