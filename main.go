package main

import (
	"github.com/lianhong2758/RosmBot/ctx"
	"github.com/lianhong2758/RosmBot/zero"

	//导入插件
	_ "github.com/lianhong2758/RosmBot/plugins/chatgpt"
	_ "github.com/lianhong2758/RosmBot/plugins/myplugin"
	_ "github.com/lianhong2758/RosmBot/plugins/test"
)

// 初始化

func main() {
	switch zero.MYSconfig.Types {
	case 0:
		ctx.RunHttp()
	case 1:
		ctx.RunWS()
	}
}
