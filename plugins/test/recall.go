package test

import (
	c "github.com/lianhong2758/RosmBot/ctx"
	log "github.com/sirupsen/logrus"
)

func init() {
	en := c.Register(&c.PluginData{
		Name: "撤回消息",
		Help: "- {回复消息}/撤回",
	})
	en.AddWord("/撤回").Rule(c.OnlyReply, c.OnlyOverOwner).Handle(func(ctx *c.CTX) {
		if err := ctx.Recall(ctx.Mess.Quote.OriginalMessageID, ctx.Mess.Quote.QuotedMessageSendTime, int64(ctx.Being.RoomID)); err != nil {
			log.Errorln("[recall]", err)
		} else {
			log.Infoln("[recall] 撤回成功,ID: ", ctx.Mess.Quote.OriginalMessageID)
		}
	})
}
