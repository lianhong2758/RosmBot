package ctx

import (
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
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
		ctx := &CTX{
			Being: &being{
				VillaID: info.Event.Robot.VillaID,
				RoomID:  23653, //欢迎大厅
			},
			Event: &info.Event.ExtendData.EventData.SendMessage,
			Bot:   &info.Event.Robot.Template,
		}
		ctx.Send(Text(info.Event.ExtendData.EventData.JoinVilla.JoinUserNickname, "欢迎光临雪儿的小屋~"))
	case 2:
		u := new(user)
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
			},
			Event: &info.Event.ExtendData.EventData.SendMessage,
			Bot:   &info.Event.Robot.Template,
		}
		log.Println(plugins)
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
