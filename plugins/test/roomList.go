package plugins

import (
	c "github.com/lianhong2758/RosmBot/ctx"
)

func init() {
	en := c.Register("roomlist", &c.PluginData{
		Name: "房间列表",
		Help: "房间列表",
	})
	en.AddWord(func(ctx *c.CTX) {
		result, err := ctx.GetRoomList()
		if err != nil {
			ctx.Send(c.Text("获取信息失败", err))
		}
		ctx.Send(c.Text(string(result)))
	}, "房间列表")
}
