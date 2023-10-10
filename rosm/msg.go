package rosm

import (
	"github.com/lianhong2758/RosmBot-MUL/message"
)

// bot主接口
type Boter interface {
	Message() any
	//获取config
	Config() any
	Send(...message.MessageSegment) any
	MakeCTX(Message any) *CTX
	//运行
	Run() error

	//Bot信息查询
	Name() string
}

// 上下文结构
type CTX struct {
	Bot     Boter
	BotType string
	Being   *Being //常用消息解析,需实现
}

// 常用数据
type Being struct {
	RoomID   int64
	RoomID2  int64
	RoomName string
	User     *UserData
	ATList   []any //at列表
	Word     string
	Rex      []string
}

// 触发者信息
type UserData struct {
	Name        string
	ID          int64
	PortraitURI string //如果直接回调没有可以不写
	Def         map[string]any
}
