package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lianhong2758/RosmBot/ctx"
	"github.com/lianhong2758/RosmBot/zero"

	"log"

	//导入插件
	_ "github.com/lianhong2758/RosmBot/plugins/chatgpt"
	_ "github.com/lianhong2758/RosmBot/plugins/test"
)

// 初始化

func main() {
	switch zero.MYSconfig.Types {
	case 0:
		gin.SetMode(gin.ReleaseMode)
		r := gin.New() //初始化
		log.Println("bot开始监听消息")
		r.POST(zero.MYSconfig.EventPath, ctx.MessReceive)
		r.GET("/file/*path", zero.GETImage)
		r.Run(zero.MYSconfig.Port)
	case 1:
	}

}
