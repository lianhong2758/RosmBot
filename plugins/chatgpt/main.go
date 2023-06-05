package chatgpt

import (
	"strings"
	"time"

	"github.com/FloatTech/ttl"
	c "github.com/lianhong2758/RosmBot/ctx"
)

type sessionKey struct {
	group int
	user  string
}

var (
	apiKey = ""
	cache  = ttl.NewCache[sessionKey, []chatMessage](time.Minute * 15)
)

func init() {
	en := c.Register("chatgpt", &c.PluginData{
		Name:       "chatgpt",
		Help:       "-@bot //|chatgpt [对话内容]\n",
		DataFolder: "chatgpt",
	})
	en.AddRex(func(ctx *c.CTX) {
		var messages []chatMessage
		args := ctx.Being.Rex[1]
		key := sessionKey{
			group: ctx.Being.VillaID,
			user:  ctx.Mess.User.ID,
		}
		if args == "reset" || args == "重置记忆" {
			cache.Delete(key)
			ctx.Send(c.Text("已清除上下文！"))
			return
		}
		messages = cache.Get(key)
		messages = append(messages, chatMessage{
			Role:    "user",
			Content: args,
		})
		resp, err := completions(messages, apiKey)
		if err != nil {
			ctx.Send(c.Text("请求ChatGPT失败: ", err))
			return
		}
		reply := resp.Choices[0].Message
		reply.Content = strings.TrimSpace(reply.Content)
		messages = append(messages, reply)
		cache.Set(key, messages)
		ctx.Send(c.Text(reply.Content, "\n本次消耗token: ", resp.Usage.PromptTokens, "+", resp.Usage.CompletionTokens, "=", resp.Usage.TotalTokens))
	}, `^(?:chatgpt|//)([\s\S]*)$`)
}
