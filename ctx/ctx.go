package ctx

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/lianhong2758/RosmBot/zero"
	log "github.com/sirupsen/logrus"
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
	log.Infoln("[ws]等待建立ws连接")
	for {
		header := http.Header{}
		header.Add("key", zero.MYSconfig.Key)
		// 建立WebSocket连接
		conn, _, err := websocket.DefaultDialer.Dial(zero.MYSconfig.Host, header)
		if err != nil {
			log.Errorln("[ws]服务器连接失败: ", err)
			time.Sleep(time.Second * 5)
			continue
		} else {
			log.Infoln("[ws]服务器连接成功: ", zero.MYSconfig.Host)
		}
		defer conn.Close()

		for {
			_, body, err := conn.ReadMessage()
			if err != nil {
				log.Errorln("[ws]服务器连接失败: ", err)
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
	log.Infoln("[http]bot开始监听消息")
	r.POST(zero.MYSconfig.EventPath, MessReceive)
	//r.GET("/file/*path", zero.GETImage)
	r.Run(zero.MYSconfig.Port)
}

func process(body []byte) {
	info := new(infoSTR)
	err := json.Unmarshal(body, info)
	if err != nil {
		log.Errorln("[info]", err)
		return
	}
	//调用消息处理件,触发中心
	switch info.Event.Type {
	default:
		log.Infoln("[info] (接收未知事件)", info.Event.ExtendData.EventData)
		return
	case 1:
		log.Debugln("[debug] (入群事件)", info.Event.ExtendData.EventData.JoinVilla)
		log.Infof("[info] (入群事件)[%d] %s(%d)", info.Event.Robot.VillaID, info.Event.ExtendData.EventData.JoinVilla.JoinUserNickname, info.Event.ExtendData.EventData.JoinVilla.JoinUID)
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
		log.Debugln("[debug] (添加bot)", info.Event.ExtendData.EventData.CreateRobot)
		log.Infof("[info] (添加Bot事件)[%d]", info.Event.Robot.VillaID)
		ctx := &CTX{
			Being: &being{
				VillaID: info.Event.Robot.VillaID,
			},
			Event: &info.Event.ExtendData.EventData,
			Bot:   &info.Event.Robot.Template,
		}
		ctx.runFuncAll(Create)
	case 4:
		log.Debugln("[debug] (删除bot)", info.Event.ExtendData.EventData.DeleteRobot)
		log.Infof("[info] (删除Bot事件)[%d]", info.Event.Robot.VillaID)
		ctx := &CTX{
			Being: &being{
				VillaID: info.Event.Robot.VillaID,
			},
			Event: &info.Event.ExtendData.EventData,
			Bot:   &info.Event.Robot.Template,
		}
		ctx.runFuncAll(Delete)
	case 5:
		log.Debugln("[debug] (接收表态)", info.Event.ExtendData.EventData.AddQuickEmoticon)
		log.Infof("[info] (表态事件)[%d] %d:%s", info.Event.Robot.VillaID, info.Event.ExtendData.EventData.AddQuickEmoticon.UID, info.Event.ExtendData.EventData.AddQuickEmoticon.Emoticon)
		ctx := &CTX{
			Being: &being{
				VillaID: info.Event.Robot.VillaID,
				User:    &user{ID: strconv.Itoa(info.Event.ExtendData.EventData.AddQuickEmoticon.UID)},
				RoomID:  info.Event.ExtendData.EventData.AddQuickEmoticon.RoomID,
			},
			Event: &info.Event.ExtendData.EventData,
			Bot:   &info.Event.Robot.Template,
		}
		ctx.emoticonNext()
		ctx.runFuncAll(Quick)
	//case 6:
	//回调审核
	case 2:
		log.Debugln("[debug] (接收消息)", info.Event.ExtendData.EventData.SendMessage.Content)
		u := new(mess)
		err = json.Unmarshal([]byte(info.Event.ExtendData.EventData.SendMessage.Content), u)
		if err != nil {
			log.Errorln("[info]", err)
			return
		}
		log.Infof("[info] (接收消息)[%d] %s:%s", info.Event.Robot.VillaID, u.User.Name, u.Content.Text)
		ctx := &CTX{
			Mess: u,
			Being: &being{
				VillaID: info.Event.Robot.VillaID,
				RoomID:  info.Event.ExtendData.EventData.SendMessage.RoomID,
				User:    &u.User,
				ATList:  u.MentionedInfo.UserIDList,
			},
			Event: &info.Event.ExtendData.EventData,
			Bot:   &info.Event.Robot.Template,
		}
		//消息处理(切割加去除尾部空格)
		word := strings.TrimSpace(ctx.Mess.Content.Text[len(ctx.Bot.Name)+2:])
		ctx.Being.Word = word
		//ctx next
		if ctx.sendNext() {
			return
		}
		//关键词触发
		if m, ok := caseAllWord[word]; ok {
			if m.RulePass(ctx) {
				m.Handler(ctx)
				log.Debugf("调用插件: %s - 匹配关键词: %s", m.PluginNode.Name, word)
			}
			return
		}
		//正则匹配
		for regex, m := range caseRegexp {
			if match := regex.FindStringSubmatch(word); len(match) > 0 {
				if m.RulePass(ctx) {
					ctx.Being.Rex = match
					m.Handler(ctx)
					log.Debugf("调用插件: %s - 匹配关键词: %s", m.PluginNode.Name, word)
				}
				if m.Block {
					return
				}
			}
		}
	}
}

// 遍历内部的函数,并执行
func (ctx *CTX) runFuncAll(types string) {
	for _, m := range caseOther[types] {
		if m.RulePass(ctx) {
			m.Handler(ctx)
			log.Debugf("调用插件: %s - 类型: %s", m.PluginNode.Name, types)
		}
		if m.Block {
			return
		}
	}
}
