# RosmBot(迷迭香Bot)
RosmBot(迷迭香Bot)是连接米游社官方接口的Bot,由golang编写
# 使用方法

1直接运行

	配置config
	运行run.bat即可

2在gin框架中合并代码

  	"github.com/lianhong2758/RosmBot/ctx"
	"github.com/lianhong2758/RosmBot/zero"

	//导入插件
	_ "github.com/lianhong2758/RosmBot/plugins/chatgpt"
	_ "github.com/lianhong2758/RosmBot/plugins/test"
	// 初始化
	func init() {
	f, err := os.ReadFile("config.json")
	if err != nil {
		f, err := os.Create("config.json")
		if err != nil {
			log.Fatalln(err)
		}
		defer f.Close()
		configdata, _ := json.MarshalIndent(zero.MYSconfig, "", "  ")
		_, _ = f.Write(configdata)
		log.Fatalln("创建初始化配置完成\n请填写config.json文件后重新运行本程序\n字段解释:\ntoken:机器人基本信息,eventpath:回调路径,port:端口\nhost:你的服务器地址/外部访问地址,如果不是80端口,需要加上端口号,结尾不需要加\"/\"")
	}
	err = json.Unmarshal(f, &zero.MYSconfig)
	if err != nil {
		log.Fatalln(err)
	}
	if zero.MYSconfig.BotToken.BotID == "" || zero.MYSconfig.BotToken.BotSecret == "" {
		log.Fatalln("[init]未设置bot信息")
		}
	}
    func main(){
     ...
        r.POST(config.EventPath, ctx.MessReceive)
		r.GET("/file/*path", zero.GETImage)
     ...
    }
之后把config.json指定目录即可