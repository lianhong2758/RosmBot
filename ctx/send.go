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
	msgContentInfo := H{}
	for _, message := range m {
		switch message.Type {
		default:
			continue
		case "text":
			msgContent.Text += message.Data["text"].(string)
		case "link", "villa_room_link":
			t := message.Data["entities"].(Entities)
			t.Offset = len(utf16.Encode([]rune(msgContent.Text)))
			msgContent.Entities = append(msgContent.Entities, t)
			msgContent.Text += message.Data["text"].(string)
		case "at", "atbot":
			t := message.Data["entities"].(Entities)
			t.Offset = len(utf16.Encode([]rune(msgContent.Text)))
			msgContent.Entities = append(msgContent.Entities, t)
			msgContent.Text += message.Data["text"].(string)
			otherUID := []string{}
			if msgContentInfo["mentionedInfo"] != nil {
				otherUID = msgContentInfo["mentionedInfo"].(MentionedInfoStr).UserIDList
			}
			otherUID = append(otherUID, message.Data["uid"].(string))
			msgContentInfo["mentionedInfo"] = MentionedInfoStr{Type: 2, UserIDList: otherUID}
		case "atall":
			t := message.Data["entities"].(Entities)
			t.Offset = len(utf16.Encode([]rune(msgContent.Text)))
			msgContent.Entities = append(msgContent.Entities, t)
			msgContent.Text += message.Data["text"].(string)
			msgContentInfo["mentionedInfo"] = MentionedInfoStr{Type: 1}
		case "imagewithtext":
			msgContent.Text += message.Data["text"].(string)
			msgContent.Images = append(msgContent.Images, message.Data["imagestr"].(ImageStr))
		case "image":
			msgContent.ImageStr = message.Data["image"].(ImageStr)
		case "reply":
			id, time := message.Data["id"].(string), message.Data["time"].(int64)
			msgContentInfo["quote"] = H{"original_message_id": id, "original_message_send_time": time, "quoted_message_id": id, "quoted_message_send_time": time}
		case "badge":
			msgContent.Badge = message.Data["badge"].(*BadgeStr)
		}
	}
	var objectStr string
	if msgContent.Text != "" {
		objectStr = "MHY:Text"
	} else {
		objectStr = "MHY:Image"
	}
	msgContentInfo["content"] = msgContent
	contentStr, _ := json.Marshal(msgContentInfo)
	data, _ := json.Marshal(H{"room_id": ctx.Being.RoomID, "object_name": objectStr, "msg_content": helper.BytesToString(contentStr)})
	data, err := web.Web(&http.Client{}, sendMessage, http.MethodPost, ctx.makeHeard, bytes.NewReader(data))
	if err != nil {
		log.Println("[send-err]", err)
	}
	var sendState sendState
	_ = json.Unmarshal(data, &sendState)
	log.Println("[send]["+sendState.Message+"]", helper.BytesToString(contentStr))

}

// 转发帖子
func (ctx *CTX) SendPost(postid string) {
	contentStr := "{\"content\":{\"post_id\":\"" + postid + "\"}}"
	data, _ := json.Marshal(H{"room_id": ctx.Being.RoomID, "villa_id": ctx.Being.VillaID, "object_name": "MHY:Post", "msg_content": contentStr})
	data, err := web.Web(&http.Client{}, sendMessage, http.MethodPost, ctx.makeHeard, bytes.NewReader(data))
	if err != nil {
		log.Println("[send-err]", err)
	}
	var sendState sendState
	_ = json.Unmarshal(data, &sendState)
	log.Println("[send]["+sendState.Message+"]", contentStr)
}

// 改变发送的房间id
func (ctx *CTX) ChangeSendRoom(roomid int) {
	ctx.Being.RoomID = roomid
}

// 改变发送的大别野id
func (ctx *CTX) ChangeSendVilla(villaid, roomid int) {
	ctx.Being.VillaID, ctx.Being.RoomID = villaid, roomid
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
func (ctx *CTX) AT(uid string) MessageSegment {
	intuid, _ := strconv.Atoi(uid)
	user, _ := ctx.GetUserData(uint64(intuid))
	name := "@" + user.Data.Member.Basic.Nickname + " "
	return MessageSegment{
		Type: "at",
		Data: H{
			"text": name,
			"uid":  uid,
			"entities": Entities{
				Length: len(utf16.Encode([]rune(name))),
				Entity: H{"type": "mentioned_user", "user_id": uid},
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
		Type: "atall",
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
func ImageUrlWithText(url string, w, h, size int, text ...any) MessageSegment {
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
func ImageUrl(url string, w, h, size int) MessageSegment {
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

// 发送普通图片,w,h为宽高size大小,不需要项填0
func Image(img []byte, w, h, size int) MessageSegment {
	if url := web.UpImgByte(img); url != "" {
		return ImageUrl(url, w, h, size)
	}
	return Text("图片上传失败")
}

// 发送普通图片和文字,w,h为宽高size大小,不需要项填0,text必填
func ImageWithText(img []byte, w, h, size int, text ...any) MessageSegment {
	if url := web.UpImgByte(img); url != "" {
		return ImageUrlWithText(url, w, h, size, text...)
	}
	return Text("图片上传失败")
}

// 发送图片文件,w,h为宽高size大小,不需要项填0
func ImageFile(path string, w, h, size int) MessageSegment {
	if url := web.UpImgfile(path); url != "" {
		return ImageUrl(url, w, h, size)
	}
	return Text("图片上传失败")
}

// 发送图片文件和文字,w,h为宽高size大小,不需要项填0,text必填
func ImageFileWithText(path string, w, h, size int, text ...any) MessageSegment {
	if url := web.UpImgfile(path); url != "" {
		return ImageUrlWithText(url, w, h, size, text...)
	}
	return Text("图片上传失败")
}

// 蓝色跳转链接
func Link(url string, haveToken bool, text ...any) MessageSegment {
	t := fmt.Sprint(text...)
	return MessageSegment{
		Type: "link",
		Data: H{
			"text": t,
			"entities": Entities{
				Length: len(utf16.Encode([]rune(t))),
				Entity: H{"type": "link", "url": url, "requires_bot_access_token": haveToken},
			},
		},
	}
}

func ReplyOther(id string, time int64) MessageSegment {
	return MessageSegment{
		Type: "reply",
		Data: H{
			"id":   id,
			"time": time,
		},
	}
}

// 回复消息
func (ctx *CTX) Reply() MessageSegment {
	return ReplyOther(ctx.Event.MsgUID, ctx.Event.SendAt)
}

// 特殊结构
// 下标文字
func Badge(iocn, text, url string) MessageSegment {
	return MessageSegment{
		Type: "badge",
		Data: H{
			"badge": BadgeStr{
				Icon: iocn,
				Text: text,
				URL:  url,
			},
		},
	}
}
