package main

import (
	"github.com/gin-gonic/gin"

	"encoding/json"
	"log"
	"os"

	"github.com/lianhong2758/RosmBot/ctx"
	"github.com/lianhong2758/RosmBot/zero"

	//导入插件
	_ "github.com/lianhong2758/RosmBot/plugins/chatgpt"
	_ "github.com/lianhong2758/RosmBot/plugins/test"
)

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

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New() //初始化
	log.Println("bot开始监听消息")
	r.POST(zero.MYSconfig.EventPath, ctx.MessReceive)
	r.GET("/file/*path", zero.GETImage)
	r.Run("0.0.0.0:" + zero.MYSconfig.Port)
}
