# RosmBot(迷迭香Bot)
RosmBot(迷迭香Bot)是连接米游社官方接口的Bot,由golang编写
# 使用方法

1直接运行

	运行run.bat即可

2在gin框架中合并代码

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

之后运行即可

# 特别鸣谢
ZeroBot提供部分代码借鉴