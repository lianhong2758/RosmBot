package zero

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func init() {
	f, err := os.ReadFile("config.json")
	if err != nil {
		fmt.Println("请输入选择的连接方式:\n0:http连接\n1:ws正向连接")
		fmt.Scanln(&MYSconfig.Types)
		if MYSconfig.Types != 0 {
			MYSconfig.Host = "ws://127.0.0.1/ws/123"
			MYSconfig.Key = "123"
		}
		f, err := os.Create("config.json")
		if err != nil {
			log.Fatalln(err)
		}
		defer f.Close()
		configdata, _ := json.MarshalIndent(MYSconfig, "", "  ")
		_, _ = f.Write(configdata)
		log.Fatalln("创建初始化配置完成\n请填写config.json文件后重新运行本程序\n字段解释:\ntoken:机器人基本信息\nhttp连接需要填写:eventpath:回调路径,port:端口\nhost:你的服务器地址/外部访问地址,如果不是80端口,需要加上端口号,结尾不需要加\"/\"\nws协议需要填写:host:ws的连接地址,key:请求头验证秘钥")
	}
	err = json.Unmarshal(f, &MYSconfig)
	if err != nil {
		log.Fatalln(err)
	}
	if MYSconfig.BotToken.BotID == "" || MYSconfig.BotToken.BotSecret == "" {
		log.Fatalln("[init]未设置bot信息")
	}
}
