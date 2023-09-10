package test

import (
	c "github.com/lianhong2758/RosmBot/ctx"
	"strings"
)

func init() {
	en := c.Register("roomlist", &c.PluginData{
		Name: "房间列表",
		Help: "- /房间列表",
	})
	en.AddWord("/房间列表").Handle(func(ctx *c.CTX) {
		result, err := ctx.GetRoomList()
		if err != nil {
			ctx.Send(c.Text("获取信息失败", err))
		}
		var msg strings.Builder
		for _, v := range result.Data.List {
			if v.GroupID == "0" {
				continue
			}
			if msg.String() != "" {
				msg.WriteByte('\n')
			}
			msg.WriteString("#" + v.GroupName)
			msg.WriteString("(" + v.GroupID + "):")
			for _, vv := range v.RoomList {
				msg.WriteString("\n" + vv.RoomName)
				msg.WriteString("(" + vv.RoomID + ")")
			}
		}
		ctx.Send(c.Text(msg.String()))
	})
}
