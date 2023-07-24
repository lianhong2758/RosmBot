package ctx

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
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
		//消息处理(切割加去除尾部空格)
		word := strings.TrimSpace(ctx.Mess.Content.Text[len(ctx.Bot.Name)+2:])
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
