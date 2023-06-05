package ctx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/lianhong2758/RosmBot/web"
	"github.com/lianhong2758/RosmBot/zero"
)

type H = map[string]any

const (
	sendMessage = "https://bbs-api.miyoushe.com/vila/api/bot/platform/sendMessage"
)

// 主动消息
func (ctx *CTX) Send(msg any) {
	m, ok := msg.(MessageSegment)
	if !ok {
		log.Println("[send-err]", "数据格式错误")
		return
	}
	data, _ := json.Marshal(H{"room_id": ctx.Being.RoomID, "object_name": m.Type, "msg_content": m.Data})
	data, err := web.Web(&http.Client{}, sendMessage, http.MethodPost, ctx.makeHeard, bytes.NewReader(data))
	if err != nil {
		log.Println("[send-err]", err)
	}
	var sendState sendState
	_ = json.Unmarshal(data, &sendState)
	log.Println("[send]["+sendState.Message+"]", m.Data)

}

func (ctx *CTX) makeHeard(req *http.Request) {
	req.Header.Add("x-rpc-bot_id", zero.BotToken.BotID)
	req.Header.Add("x-rpc-bot_secret", zero.BotToken.BotSecret)
	req.Header.Add("x-rpc-bot_villa_id", strconv.Itoa(ctx.Being.VillaID))
	req.Header.Add("Content-Type", "application/json")
}

// 消息解析
func Text(text ...any) MessageSegment {
	return MessageSegment{
		Type: "MHY:Text",
		Data: func() string {
			data, _ := json.Marshal(H{
				"content": H{
					"text": fmt.Sprint(text...),
				},
			})
			return string(data)
		}(),
	}
}

func Image(url string, text ...any) MessageSegment {
	return MessageSegment{
		Type: "MHY:Text",
		Data: func() string {
			data, _ := json.Marshal(H{
				"content": H{
					"images": []any{H{
						"size": H{
							"width":  2800,
							"height": 3400,
						},
						"url":       url,
						"file_size": 9999,
					}},
					"text":     fmt.Sprint(text...),
					"entities": nil,
				},
			})
			return string(data)
		}(),
	}
}

func Link(url string, text ...any) MessageSegment {
	t := fmt.Sprint(text...)
	offset, lenght := 0, 0
	for i := 0; i <= len([]rune(t))-len([]rune(url)); i++ {
		if string(t[i:i+len(url)]) == url {
			offset = i
			lenght = len([]rune(url))
			break
		}
	}
	return MessageSegment{
		Type: "MHY:Text",
		Data: func() string {
			data, _ := json.Marshal(H{
				"content": H{
					"text": t,
					"entities": []any{H{
						"offset": offset,
						"length": lenght,
						"entity": H{
							"type": "link",
							"url":  url,
						}},
					},
				},
			})
			return string(data)
		}(),
	}
}

type MessageSegment struct {
	Type string `json:"type"`
	Data string `json:"data"`
}
