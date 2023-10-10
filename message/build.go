package message

import (
	"fmt"
	"github.com/lianhong2758/RosmBot-MUL/tool/web"
)

type H = map[string]any

type Message []MessageSegment

type MessageSegment struct {
	Type string `json:"type"`
	Data H      `json:"data"`
}

// 消息解析
func Text(text ...any) MessageSegment {
	return MessageSegment{
		Type: "text",
		Data: H{"text": fmt.Sprint(text...)},
	}
}

// at用户
func AT(uid, name string) MessageSegment {
	return MessageSegment{
		Type: "at",
		Data: H{
			"text": name,
			"uid":  uid,
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
			"uid":  botid,
		},
	}
}

// at all
func ATAll() MessageSegment {
	return MessageSegment{
		Type: "atall",
	}
}

func ImageUrlWithText(url string, text ...any) MessageSegment {
	return MessageSegment{
		Type: "image",
		Data: H{
			"url": url,
		},
	}
}

// url为图片链接,必须直链,w,h为宽高size大小,不需要项填0
func ImageUrl(url string) MessageSegment {
	return MessageSegment{
		Type: "image",
		Data: H{
			"url": url,
		},
	}
}

// 发送普通图片
func Image(img []byte) MessageSegment {
	if url := web.UpImgByte(img); url != "" {
		return ImageUrl(url)
	}
	return Text("图片上传失败")
}

// 发送普通图片和文字,text必填
func ImageWithText(img []byte, text ...any) MessageSegment {
	if url := web.UpImgByte(img); url != "" {
		return ImageUrlWithText(url, text...)
	}
	return Text("图片上传失败")
}

// 发送图片文件
func ImageFile(path string) MessageSegment {
	if url := web.UpImgfile(path); url != "" {
		return ImageUrl(url)
	}
	return Text("图片上传失败")
}

// 发送图片文件和文字,text必填
func ImageFileWithText(path string, text ...any) MessageSegment {
	if url := web.UpImgfile(path); url != "" {
		return ImageUrlWithText(url, text...)
	}
	return Text("图片上传失败")
}

// 蓝色跳转链接
func Link(url string, haveToken bool, text ...any) MessageSegment {
	return MessageSegment{
		Type: "link",
		Data: H{
			"text":  fmt.Sprint(text...),
			"url":   url,
			"token": haveToken,
		},
	}
}

func ReplyOther(id string, more ...H) MessageSegment {
	var t H
	if len(more) > 0 {
		t = more[0]
	}
	return MessageSegment{
		Type: "reply",
		Data: H{
			"id":   id,
			"more": t,
		},
	}
}
