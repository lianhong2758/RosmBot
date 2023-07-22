## RosmBot(迷迭香Bot)
RosmBot(迷迭香Bot)是大别野(Villa)相关Bot-SDK,由golang编写
## 使用方法

1直接运行
```
	运行run.bat即可
```
2在gin框架中合并代码
```
  	"github.com/lianhong2758/RosmBot/ctx"
	"github.com/lianhong2758/RosmBot/zero"

	//导入插件
	_ "github.com/lianhong2758/RosmBot/plugins/chatgpt"
	_ "github.com/lianhong2758/RosmBot/plugins/test"
	
    func main(){
     ...
        r.POST(config.EventPath, ctx.MessReceive)
     ...
    }
```
之后运行即可

## 插件编写教程

1注册插件
```
func init() {
	en := c.Register("chat", &c.PluginData{//第一个参数是插件名,用于区分插件
		Name: "@回复",        			   //插件名,用于help
		Help: "- @机器人",				   //帮助信息,用于help
		DataFolder: "chat",				   //可选,创建插件的数据文件夹,不需要数据存储则不需要填写
	})
	en.AddWord(func(ctx *c.CTX) {			//创建一个完全匹配指令,之后的匹配后执行的函数
		ctx.Send(c.Text(zero.MYSconfig.BotToken.BotName, "不在呢~"))
	}, "")									//这里是匹配词
}
```
2获取触发时传送的数据
```
ctx.Being里有所有需要的数据,结构如下
type Being struct {
	RoomID  int		//房间号
	VillaID int		//大别野号
	User    *user	//触发者的信息
	Word    string	//如果是word触发(完全匹配触发),则这里是触发词
	Rex     []string//如果是rex触发(正则匹配触发),则这里是正则全匹配的数组
}
```
3发送消息
```
1)文本或者图片消息
ctx.Send(xxx)
xxx有很多,可以无限续接,逗号分开
其中文本消息用c.Text(any)
byte图片用c.Image(img []byte, w, h, size int)
url图片用c.ImageUrl(url string, w, h, size int)
at用ctx.AT(id)
reply用ctx.reply()
其余看源码学习...
2)帖子消息
 ctx.SendPost(postid string)
```
4更改发送房间
```
ctx.ChangeSendRoom(roomid int)//更改发送房间
ctx.ChangeSendVilla(villaid, roomid int)//更改发送别野
```
5部分接口(可能存在没有及时更新,导致调用出错的情况,如有请反馈)
```
ctx.GetRoomList()//获取房间列表
ctx.GetUserData(uid uint64)//获取某人信息
ctx.DeleteUser(uid uint64)//踢人
ctx.Recall(msgid, string, roomid uint64, msgtime int64)//撤回消息
```
6启用插件
```
如果编写的插件没有在plugins/test里面,请手动在main.go里面进行导入注册
```
## 特别鸣谢
[ZeroBot](https://github.com/wdvxdr1123/ZeroBot)提供部分代码借鉴
## 相关地址

- [大别野Bot开放平台](https://open.miyoushe.com/#/login)

- [官方API文档](https://webstatic.mihoyo.com/vila/bot/doc/)

- [SDK交流大别野](https://dby.miyoushe.com/chat/1722/23652)