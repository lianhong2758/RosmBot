package ctx

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/lianhong2758/RosmBot/zero"
	"github.com/wdvxdr1123/ZeroBot/utils/helper"
)

func MessReceive(c *gin.Context) {
	body, _ := c.GetRawData()
	c.JSON(200, map[string]any{"message": "", "retcode": 0}) //确认接收
	sign := c.GetHeader("x-rpc-bot_sign")
	if verify(sign, helper.BytesToString(body), zero.MYSconfig.BotToken.BotSecretConst, zero.MYSconfig.BotToken.BotPubKey) {
		process(body)
	}
}

func RunWS() {
	log.Println("[ws]等待建立ws连接")
	for {
		header := http.Header{}
		header.Add("key", zero.MYSconfig.Key)
		// 建立WebSocket连接
		conn, _, err := websocket.DefaultDialer.Dial(zero.MYSconfig.Host, header)
		if err != nil {
			log.Println("[ws]服务器连接失败: ", err)
			time.Sleep(time.Second * 5)
			continue
		} else {
			log.Println("[ws]服务器连接成功: ", zero.MYSconfig.Host)
		}
		defer conn.Close()

		for {
			_, body, err := conn.ReadMessage()
			if err != nil {
				log.Println("[ws]服务器连接失败: ", err)
				break
			}
			process(body)
			//延时
			time.Sleep(time.Millisecond * 100)
		}
	}
}
func RunHttp() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New() //初始化
	log.Println("[http]bot开始监听消息")
	r.POST(zero.MYSconfig.EventPath, MessReceive)
	//r.GET("/file/*path", zero.GETImage)
	r.Run(zero.MYSconfig.Port)
}

func process(body []byte) {
	info := new(infoSTR)
	err := json.Unmarshal(body, info)
	if err != nil {
		log.Println("[info-err]", err)
		return
	}
	//调用消息处理件,触发中心
	switch info.Event.Type {
	default:
		log.Println("[info] (接收未知事件)", info.Event.ExtendData.EventData)
		return
	case 1:
		log.Printf("[info] (入群事件)[%d] %s(%d)\n", info.Event.Robot.VillaID, info.Event.ExtendData.EventData.JoinVilla.JoinUserNickname, info.Event.ExtendData.EventData.JoinVilla.JoinUID)
		ctx := &CTX{
			Being: &being{
				VillaID: info.Event.Robot.VillaID,
				User:    &user{ID: strconv.Itoa(info.Event.ExtendData.EventData.JoinVilla.JoinUID), Name: info.Event.ExtendData.EventData.JoinVilla.JoinUserNickname},
			},
			Event: &info.Event.ExtendData.EventData,
			Bot:   &info.Event.Robot.Template,
		}
		ctx.runFuncAll(Join)
	case 3:
		log.Printf("[info] (添加Bot事件)[%d]\n", info.Event.Robot.VillaID)
		ctx := &CTX{
			Being: &being{
				VillaID: info.Event.Robot.VillaID,
			},
			Event: &info.Event.ExtendData.EventData,
			Bot:   &info.Event.Robot.Template,
		}
		ctx.runFuncAll(Create)
	case 4:
		log.Printf("[info] (删除Bot事件)[%d]\n", info.Event.Robot.VillaID)
		ctx := &CTX{
			Being: &being{
				VillaID: info.Event.Robot.VillaID,
			},
			Event: &info.Event.ExtendData.EventData,
			Bot:   &info.Event.Robot.Template,
		}
		ctx.runFuncAll(Delete)
	case 5:
		log.Printf("[info] (表态事件)[%d] %d:%s\n", info.Event.Robot.VillaID, info.Event.ExtendData.EventData.AddQuickEmoticon.UID, info.Event.ExtendData.EventData.AddQuickEmoticon.Emoticon)
		ctx := &CTX{
			Being: &being{
				VillaID: info.Event.Robot.VillaID,
				User:    &user{ID: strconv.Itoa(info.Event.ExtendData.EventData.AddQuickEmoticon.UID)},
				RoomID:  info.Event.ExtendData.EventData.AddQuickEmoticon.RoomID,
			},
			Event: &info.Event.ExtendData.EventData,
			Bot:   &info.Event.Robot.Template,
		}
		ctx.runFuncAll(Quick)
	//case 6:
	case 2:
		//log.Println(info.Event.ExtendData.EventData.SendMessage.Content)
		u := new(mess)
		err = json.Unmarshal([]byte(info.Event.ExtendData.EventData.SendMessage.Content), u)
		if err != nil {
			log.Println("[info-err]", err)
			return
		}
		log.Printf("[info] (接收消息)[%d] %s:%s\n", info.Event.Robot.VillaID, u.User.Name, u.Content.Text)
		ctx := &CTX{
			Mess: u,
			Being: &being{
				VillaID: info.Event.Robot.VillaID,
				RoomID:  info.Event.ExtendData.EventData.SendMessage.RoomID,
				User:    &u.User,
			},
			Event: &info.Event.ExtendData.EventData,
			Bot:   &info.Event.Robot.Template,
		}
		//消息处理(切割加去除尾部空格)
		word := strings.TrimSpace(ctx.Mess.Content.Text[len(ctx.Bot.Name)+2:])
		ctx.Being.Word = word
		//ctx next
		ctx.SendNext()
		//关键词触发
		if f, ok := caseAllWord[word]; ok {
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

// 遍历内部的函数,并执行
func (ctx *CTX) runFuncAll(types string) {
	for _, f := range caseOther[types] {
		f(ctx)
	}
}
