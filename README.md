# RosmBot
米游社官方接口Bot,golang编写
# 使用方法

1直接运行main.go即可

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
     ...
    }
之后把config.json指定目录即可