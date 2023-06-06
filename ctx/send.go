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
	"github.com/wdvxdr1123/ZeroBot/utils/helper"
)

type H = map[string]any

const (
	sendMessage = "https://bbs-api.miyoushe.com/vila/api/bot/platform/sendMessage"
)

// 主动消息
func (ctx *CTX) Send(m MessageSegment) {
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
	req.Header.Add("x-rpc-bot_id", zero.MYSconfig.BotToken.BotID)
	req.Header.Add("x-rpc-bot_secret", zero.MYSconfig.BotToken.BotSecret)
	req.Header.Add("x-rpc-bot_villa_id", strconv.Itoa(ctx.Being.VillaID))
	req.Header.Add("Content-Type", "application/json")
}

// 消息解析
func Text(text ...any) MessageSegment {
	return MessageSegment{
		Type: "MHY:Text",
		Data: func() string {
			data, _ := json.Marshal(H{"content": Content{Text: fmt.Sprint(text...)}})
			return helper.BytesToString(data)
		}(),
	}
}

// url为图片链接,必须直链,w,h为宽高
func ImageWithText(url string, w, h, size int, text ...any) MessageSegment {
	return MessageSegment{
		Type: "MHY:Text",
		Data: func() string {
			images := ImageStr{
				URL: url,
			}
			if w != 0 {
				images.Size.Width = w
			}
			if h != 0 {
				images.Size.Height = h
			}
			if size != 0 {
				images.FileSize = size
			}
			data, _ := json.Marshal(H{"content": Content{
				Text:   fmt.Sprint(text...),
				Images: []ImageStr{images},
			}})
			return helper.BytesToString(data)
		}(),
	}
}

// url为图片链接,必须直链,w,h为宽高size大小,不需要项填0
func Image(url string, w, h, size int) MessageSegment {
	return MessageSegment{
		Type: "MHY:Text",
		Data: func() string {
			content := Content{
				ImageStr: ImageStr{URL: url},
			}
			if w != 0 {
				content.Size.Width = w
			}
			if h != 0 {
				content.Size.Height = h
			}
			if size != 0 {
				content.FileSize = size
			}
			data, _ := json.Marshal(H{"content": content})
			return helper.BytesToString(data)
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
			return helper.BytesToString(data)
		}(),
	}
}

type MessageSegment struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

// 消息模板
type Content struct {
	//图片
	ImageStr
	//文本
	Text     string `json:"text,omitempty"`
	Entities []struct {
		Entity struct {
			Type string `json:"type,omitempty"`
		} `json:"entity,omitempty"`
		Length int `json:"length,omitempty"`
		Offset int `json:"offset,omitempty"`
	} `json:"entities,omitempty"`
	Images []ImageStr `json:"images,omitempty"`
}
type ImageStr struct {
	URL      string `json:"url,omitempty"`
	FileSize int    `json:"file_size,omitempty"`
	Size     struct {
		Height int `json:"height,omitempty"`
		Width  int `json:"width,omitempty"`
	} `json:"size,omitempty"`
}
