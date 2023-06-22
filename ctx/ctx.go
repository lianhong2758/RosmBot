package ctx

import (
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lianhong2758/RosmBot/zero"
	"github.com/wdvxdr1123/ZeroBot/utils/helper"
)

func MessReceive(c *gin.Context) {
	body, _ := c.GetRawData()
	info := new(infoSTR)
	err := json.Unmarshal(body, info)
	if err != nil {
		log.Println("[info-err]", err)
		c.JSON(200, gin.H{"message": "", "retcode": 0})
		return
	}
	c.JSON(200, map[string]any{"message": "", "retcode": 0}) //确认接收
	//调用消息处理件,触发中心
	switch info.Event.Type {
	default:
		log.Println(info.Event.ExtendData.EventData)
		return
	case 1:
		log.Printf("[info] (入群事件) %s(%d)", info.Event.ExtendData.EventData.JoinVilla.JoinUserNickname, info.Event.ExtendData.EventData.JoinVilla.JoinUID)
		//设置欢迎房间
		wel, _ := os.ReadFile("data/setting/welcome.txt")
		welcomeRoom, _ := strconv.Atoi(helper.BytesToString(wel))
		ctx := &CTX{
			Being: &being{
				VillaID: info.Event.Robot.VillaID,
				RoomID:  welcomeRoom, //欢迎大厅
				User:    &user{ID: strconv.Itoa(info.Event.ExtendData.EventData.JoinVilla.JoinUID), Name: info.Event.ExtendData.EventData.JoinVilla.JoinUserNickname},
			},
			Event: &info.Event.ExtendData.EventData.SendMessage,
			Bot:   &info.Event.Robot.Template,
		}
		if ctx.Being.RoomID != 0 {
			ctx.Send(Text(info.Event.ExtendData.EventData.JoinVilla.JoinUserNickname, "欢迎光临", zero.MYSconfig.BotToken.BotName, "的小屋~"))
		}
	case 2:
		u := new(mess)
		err = json.Unmarshal([]byte(info.Event.ExtendData.EventData.SendMessage.Content), u)
		if err != nil {
			log.Println("[info-err]", err)
			c.JSON(200, gin.H{"message": "", "retcode": 0})
			return
		}
		log.Println("[info] (接收消息)", u.User.Name, ":", u.Content.Text)
		ctx := &CTX{
			Mess: u,
			Being: &being{
				VillaID: info.Event.Robot.VillaID,
				RoomID:  info.Event.ExtendData.EventData.SendMessage.RoomID,
				User:    &u.User,
			},
			Event: &info.Event.ExtendData.EventData.SendMessage,
			Bot:   &info.Event.Robot.Template,
		}
		//消息处理
		word := ctx.Mess.Content.Text[11:]
		//关键词触发
		if f, ok := caseAllWord[word]; ok {
			ctx.Being.Word = word
			f(ctx)
			return
		}
		//正则匹配
		for regex, f := range caseRegexp {
			if match := regex.FindStringSubmatch(word); len(match) > 0 {
				ctx.Being.Rex = append(ctx.Being.Rex, match...)
				f(ctx)
				return
			}
		}
	}
}
