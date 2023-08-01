package test

import (
	"encoding/json"
	"os"
	"strconv"

	c "github.com/lianhong2758/RosmBot/ctx"
)

var list = map[int]int{}

func init() {
	en := c.Register("setting", &c.PluginData{
		Name:       "入群欢迎",
		Help:       "-设置欢迎房间xxx",
		DataFolder: "join",
	})
	//读取
	jd, err := os.ReadFile("data/init/F.txt")
	if err == nil {
		// 将 JSON 数据转换为 map
		err = json.Unmarshal(jd, &list)
		if err != nil {
			panic(err)
		}
	}
	en.AddRex(func(ctx *c.CTX) {
		key := ctx.Being.Rex[1]
		if key == "" {
			ctx.Send(c.Text("输入房间号错误"))
			return
		}
		intkey, err := strconv.Atoi(key)
		if err != nil {
			ctx.Send(c.Text("输入房间号错误"))
			return
		}
		//添加一个欢迎别野
		list[ctx.Being.VillaID] = intkey
		jsonData, _ := json.Marshal(list)

		err = os.WriteFile(en.DataFolder+"join.txt", jsonData, 0644)
		if err != nil {
			panic(err)
		}
		ctx.Send(c.Text("设置成功"))
	}, "^设置欢迎房间(.*)")
	en.AddOther(func(ctx *c.CTX) {
		ctx.ChangeSendRoom(list[ctx.Being.VillaID])
		if ctx.Being.RoomID != 0 {
			r, _ := ctx.GetVillaData()
			villaName := r.Data.Villa.Name
			ctx.Send(ctx.AT(ctx.Being.User.ID), c.Text("欢迎光临", villaName, "~"))
		}
	}, c.Join)
}
