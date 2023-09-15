package ctx

import "github.com/lianhong2758/RosmBot/zero"
import log "github.com/sirupsen/logrus"

func (m *Matcher) RulePass(ctx *CTX) bool {
	for _, v := range m.Rules {
		if !v(ctx) {
			return false
		}
	}
	return true
}

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
	return IsMaster(ctx.Being.User.ID)
}

func OnlyMaster(ctx *CTX) bool {
	return ctx.IsMaster()
}

// 别野房东权限以上
func OnlyOverOwner(ctx *CTX) bool {
	data, err := ctx.GetVillaData()
	if err != nil {
		log.Errorln("[ctx](", ctx.Being.VillaID, ")获取别野信息失败:", err)
	}
	return ctx.Being.User.ID == data.Data.Villa.OwnerUID || ctx.IsMaster()
}

// 触发消息是否是回复消息
func OnlyReply(ctx *CTX) bool {
	return ctx.Mess.Quote.QuotedMessageSendTime != 0 && ctx.Mess.Quote.OriginalMessageID != ""
}
