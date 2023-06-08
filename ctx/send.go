package ctx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"unicode/utf16"

	"github.com/lianhong2758/RosmBot/web"
	"github.com/lianhong2758/RosmBot/zero"
	"github.com/wdvxdr1123/ZeroBot/utils/helper"
)

type H = map[string]any

const (
	sendMessage = "https://bbs-api.miyoushe.com/vila/api/bot/platform/sendMessage"
)

// 主动消息
func (ctx *CTX) Send(m ...MessageSegment) {
	msgContent := new(Content)
	for _, message := range m {
		switch message.Type {
		case "text":
			msgContent.Text += message.Data["text"].(string)
		case "at", "atbot", "link", "villa_room_link":
			t := message.Data["entities"].(Entities)
			t.Offset = len(utf16.Encode([]rune(msgContent.Text)))
			msgContent.Entities = append(msgContent.Entities, t)
			msgContent.Text += message.Data["text"].(string)
		case "imagewithtext":
			msgContent.Text += message.Data["text"].(string)
			msgContent.Images = append(msgContent.Images, message.Data["imagestr"].(ImageStr))
		case "image":
			msgContent.ImageStr = message.Data["image"].(ImageStr)
		}
	}
	var objectStr string
	if msgContent.Text != "" {
		objectStr = "MHY:Text"
	} else {
		objectStr = "MHY:Image"
	}
	contentStr, _ := json.Marshal(H{"content": msgContent})
	data, _ := json.Marshal(H{"room_id": ctx.Being.RoomID, "object_name": objectStr, "msg_content": helper.BytesToString(contentStr)})
	data, err := web.Web(&http.Client{}, sendMessage, http.MethodPost, ctx.makeHeard, bytes.NewReader(data))
	if err != nil {
		log.Println("[send-err]", err)
	}
	var sendState sendState
	_ = json.Unmarshal(data, &sendState)
	log.Println("[send]["+sendState.Message+"]", helper.BytesToString(contentStr))

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
		Type: "text",
		Data: H{"text": fmt.Sprint(text...)},
	}
}

// at用户
func (ctx *CTX) AT(uid uint64) MessageSegment {
	user, _ := ctx.GetUserData(uid)
	name := "@" + user.Data.Member.Basic.Nickname + " "
	return MessageSegment{
		Type: "at",
		Data: H{
			"text": name,
			"entities": Entities{
				Length: len(utf16.Encode([]rune(name))),
				Entity: H{"type": "mentioned_user", "user_id": strconv.Itoa(int(uid))},
			},
		},
	}
}

// atbot
func ATBot(botid, botname string) MessageSegment {
	name := "@" + botname + " "
	return MessageSegment{
		Type: "atbot",
		Data: H{
			"text": name,
			"entities": Entities{
				Length: len(utf16.Encode([]rune(name))),
				Entity: H{"type": "mentioned_robot", "bot_id": botid},
			},
		},
	}
}

// at all
func ATAll() MessageSegment {
	name := "@全体成员 "
	return MessageSegment{
		Type: "atbot",
		Data: H{
			"text": name,
			"entities": Entities{
				Length: len(utf16.Encode([]rune(name))),
				Entity: H{"type": "mention_all"},
			},
		},
	}
}

// goto the room
func (ctx *CTX) RoomLink(roomid string) MessageSegment {
	r, _ := ctx.GetRoomList()
	RoomName := roomid
GroupFor:
	for _, v := range r.Data.List {
		for _, vv := range v.RoomList {
			if vv.RoomID == roomid {
				RoomName = vv.RoomName
				break GroupFor
			}
		}
	}
	name := "#" + RoomName + " "
	return MessageSegment{
		Type: "villa_room_link",
		Data: H{
			"text": name,
			"entities": Entities{
				Length: len(utf16.Encode([]rune(name))),
				Entity: H{"type": "villa_room_link", "villa_id": strconv.Itoa(ctx.Being.VillaID), "room_id": roomid},
			},
		},
	}
}

// url为图片链接,必须直链,w,h为宽高
func ImageWithText(url string, w, h, size int, text ...any) MessageSegment {
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
	return MessageSegment{
		Type: "imagewithtext",
		Data: H{
			"text":     fmt.Sprint(text...),
			"imagestr": images,
		},
	}
}

// url为图片链接,必须直链,w,h为宽高size大小,不需要项填0
func Image(url string, w, h, size int) MessageSegment {
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
	return MessageSegment{
		Type: "image",
		Data: H{
			"image": images,
		},
	}
}

// 蓝色跳转链接
func Link(url string, text ...any) MessageSegment {
	t := fmt.Sprint(text...)
	return MessageSegment{
		Type: "link",
		Data: H{
			"text": t,
			"entities": Entities{
				Length: len(utf16.Encode([]rune(t))),
				Entity: H{"type": "link", "url": url},
			},
		},
	}
}

type Message []MessageSegment
type MessageSegment struct {
	Type string `json:"type"`
	Data H      `json:"data"`
}

// 消息模板
type Content struct {
	//图片
	ImageStr
	//文本
	Text     string     `json:"text,omitempty"`
	Entities []Entities `json:"entities,omitempty"`
	Images   []ImageStr `json:"images,omitempty"`
}
type ImageStr struct {
	URL      string `json:"url,omitempty"`
	FileSize int    `json:"file_size,omitempty"`
	Size     struct {
		Height int `json:"height,omitempty"`
		Width  int `json:"width,omitempty"`
	} `json:"size,omitempty"`
}
type Entities struct {
	Entity H   `json:"entity,omitempty"`
	Length int `json:"length,omitempty"`
	Offset int `json:"offset,omitempty"`
}
