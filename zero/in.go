package zero

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func init() {
	f, err := os.ReadFile("config.json")
	if err != nil {
		fmt.Println("请输入选择的连接方式:\n0:http连接\n1:ws正向连接")
		fmt.Scanln(&MYSconfig.Types)
		if MYSconfig.Types != 0 {
			MYSconfig.Host = "ws://47.120.13.24/ws/id"
			MYSconfig.Key = "123"
			MYSconfig.BotToken.BotPubKey = "-----BEGIN PUBLIC KEY----- abcabc123 -----END PUBLIC KEY----- "
		} else {
			MYSconfig.EventPath = "/"
			MYSconfig.Port = "0.0.0.0:80"
		}
		f, err := os.Create("config.json")
		if err != nil {
			log.Fatalln(err)
		}
		defer f.Close()
		configdata, _ := json.MarshalIndent(MYSconfig, "", "  ")
		_, _ = f.Write(configdata)
		log.Fatalln("创建初始化配置完成\n请填写config.json文件后重新运行本程序\n字段解释:\ntoken:机器人基本信息\nhttp连接需要填写:eventpath:回调路径,port:端口\nws协议需要填写:host:ws的连接地址,key:请求头验证秘钥")
	}
	err = json.Unmarshal(f, &MYSconfig)
	if err != nil {
		log.Fatalln(err)
	}
	if MYSconfig.BotToken.BotID == "" || MYSconfig.BotToken.BotSecret == "" {
		log.Fatalln("[init]未设置bot信息")
	}
	//备份
	MYSconfig.BotToken.BotSecretConst = MYSconfig.BotToken.BotSecret
	//修正
	var pubKeynext strings.Builder
	s := strings.Fields(MYSconfig.BotToken.BotPubKey)
	for k, v := range s {
		if k < 2 || k > len(s)-4 {
			pubKeynext.WriteString(v)
			pubKeynext.WriteString(" ")
		} else {
			pubKeynext.WriteString(v)
			pubKeynext.WriteString("\n")
		}
	}
	MYSconfig.BotToken.BotPubKey = strings.TrimSpace(pubKeynext.String()) + "\n"
	//加密验证
	MYSconfig.BotToken.BotSecret = Sha256HMac(MYSconfig.BotToken.BotPubKey, MYSconfig.BotToken.BotSecret)
}

// HMAC/SHA256加密
func Sha256HMac(pubKey string, botSecret string) string {
	h := hmac.New(sha256.New, []byte(pubKey))
	raw := []byte(botSecret)
	h.Write(raw)
	return hex.EncodeToString(h.Sum(nil))
}
