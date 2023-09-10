package test

import (
	"log"

	c "github.com/lianhong2758/RosmBot/ctx"
)

func init() {
	en := c.Register("recall", &c.PluginData{
		Name: "撤回消息",
		Help: "- {回复消息}/撤回",
	})
	en.AddWord("/撤回").Handle(func(ctx *c.CTX) {
		if ctx.Mess.Quote.QuotedMessageSendTime != 0 && ctx.Mess.Quote.OriginalMessageID != "" {
			if err := ctx.Recall(ctx.Mess.Quote.OriginalMessageID, ctx.Mess.Quote.QuotedMessageSendTime, int64(ctx.Being.RoomID)); err != nil {
				log.Println("[recall-err]", err)
			} else {
				log.Println("[recall] 撤回成功,ID: ", ctx.Mess.Quote.OriginalMessageID)
			}
		}
	})
}
