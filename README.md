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
		panic(err)
	}
	err = json.Unmarshal(f, &zero.MYSconfig)
	if err != nil {
		panic(err)
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