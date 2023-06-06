# RosmBot(迷迭香Bot)
RosmBot(迷迭香Bot)是连接米游社官方接口的Bot,由golang编写
# 使用方法

1直接运行

	配置config
	运行run.bat即可

2在gin框架中合并代码

    "github.com/lianhong2758/RosmBot/ctx"
	_ "github.com/lianhong2758/RosmBot/plugins"
	"github.com/lianhong2758/RosmBot/zero"
    // 初始化
    var config mysCFG

    func init() {
	f, err := os.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(f, &config)
	if err != nil {
		panic(err)
	}
	zero.BotToken = &config.BotToken
	if config.BotToken.BotID == "" || config.BotToken.BotSecret == "" {
		log.Fatalln("[init]未设置bot信息")
	    }
    }

    type mysCFG struct {
	    BotToken  zero.Token `json:"token"`
	    EventPath string     `json:"eventpath"`
    }
    func main(){
     ...
        r.POST(config.EventPath, ctx.MessReceive)
		r.GET("/file/*path", zero.GETImage)
     ...
    }
之后把config.json指定目录即可