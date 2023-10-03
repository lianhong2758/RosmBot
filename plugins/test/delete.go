package test

import (
	"strconv"

	c "github.com/lianhong2758/RosmBot/ctx"
	"github.com/sirupsen/logrus"
)

func init() {
	en := c.Register(c.NewRegist("踢出别野", "- @机器人 踢出别野 @everyone", ""))
	en.AddRex(`踢出别野(.*)`).SetBlock(true).Rule(c.OnlyOverOwner).Handle(func(ctx *c.CTX) {
		list := ctx.Being.ATList
		if len(list) != 2 {
			return
		}
		x, _ := strconv.Atoi(list[1])
		logrus.Infof("[delete]别野%v 删除用户%v ", ctx.Being.VillaID, x)
		err := ctx.DeleteUser(uint64(x))
		ctx.Send(c.Text(err))
	})
}
