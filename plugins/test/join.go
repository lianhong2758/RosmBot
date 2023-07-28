package test

import (
	"os"
	"strconv"

	c "github.com/lianhong2758/RosmBot/ctx"
	"github.com/lianhong2758/RosmBot/zero"
	"github.com/wdvxdr1123/ZeroBot/utils/helper"
)

func init() {
	en := c.Register("setting", &c.PluginData{
		Name:       "入群欢迎",
		Help:       "-设置欢迎房间xxx",
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
	en.AddOther(func(ctx *c.CTX) {
		wel, _ := os.ReadFile("data/setting/welcome.txt")
		welcomeRoom, _ := strconv.Atoi(helper.BytesToString(wel))
		ctx.ChangeSendRoom(welcomeRoom)
		if ctx.Being.RoomID != 0 {
			ctx.Send(c.Text(ctx.Being.User.Name, "欢迎光临", zero.MYSconfig.BotToken.BotName, "的小屋~"))
		}
	}, c.Join)
}
