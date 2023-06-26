package ctx

import (
	"github.com/lianhong2758/RosmBot/zero"
)

// 是否是主人权限
func IsMaster(userID string) bool {
	for _, v := range zero.MYSconfig.BotToken.Master {
		if v == userID {
			return true
		}
	}
	return false
}

func (ctx *CTX) IsMaster() bool {
	for _, v := range zero.MYSconfig.BotToken.Master {
		if v == ctx.Being.User.ID {
			return true
		}
	}
	return false
}
