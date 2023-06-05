package main

import (
	"github.com/gin-gonic/gin"

	"encoding/json"
	"log"
	"os"

	"github.com/lianhong2758/RosmBot/ctx"
	_ "github.com/lianhong2758/RosmBot/plugins"
	"github.com/lianhong2758/RosmBot/zero"
)

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

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New() //初始化
	log.Println("bot开始监听消息")
	r.POST(config.EventPath, ctx.MessReceive)
	r.Run("0.0.0.0:80")
}

type mysCFG struct {
	BotToken  zero.Token `json:"token"`
	EventPath string     `json:"eventpath"`
}
