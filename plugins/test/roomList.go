package test

import (
	c "github.com/lianhong2758/RosmBot/ctx"
	"encoding/json"
	"strings"
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
		var u roomstr
		err = json.Unmarshal(result, &u)
		if err != nil {
			ctx.Send(c.Text("ERROR: ", err))
		}
		var msg strings.Builder
		for _, v := range u.Data.List {
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
	}, "房间列表")
}

type roomstr struct {
	Retcode int    `json:"retcode"`
	Message string `json:"message"`
	Data    struct {
		List []struct {
			GroupID   string `json:"group_id"`
			GroupName string `json:"group_name"`
			RoomList  []struct {
				RoomID   string `json:"room_id"`
				RoomName string `json:"room_name"`
				RoomType string `json:"room_type"`
				GroupID  string `json:"group_id"`
			} `json:"room_list"`
		} `json:"list"`
	} `json:"data"`
}
