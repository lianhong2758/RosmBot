package ctx

var nextMessList = map[int]chan *CTX{}
var nextEmoticonList = map[string]chan *CTX{}

// 获取本房间全体的下一句话
func (ctx *CTX) GetNextAllMess() (chan *CTX, func()) {
	next := make(chan *CTX, 1)
	id := ctx.Being.VillaID + ctx.Being.RoomID
	nextMessList[id] = next
	return next, func() {
		close(next)
		delete(nextMessList, id)
	}
}

// 获取本房间该用户的下一句话
func (ctx *CTX) GetNextUserMess() (chan *CTX, func()) {
	next := make(chan *CTX, 1)
	id := ctx.Being.VillaID + ctx.Being.RoomID + ctx.IntUserID()
	nextMessList[id] = next
	return next, func() {
		close(next)
		delete(nextMessList, id)
	}
}

func (ctx *CTX) sendNext() (block bool) {
	if len(nextMessList) == 0 {
		return false
	}
	//先匹配个人
	if c, ok := nextMessList[ctx.Being.VillaID+ctx.Being.RoomID+ctx.IntUserID()]; ok {
		c <- ctx
		return true
	}
	if c, ok := nextMessList[ctx.Being.VillaID+ctx.Being.RoomID]; ok {
		c <- ctx
		return true
	}
	return false
}

// 获取本消息-全体的下一表态
func (ctx *CTX) GetNextAllEmoticon(botMsgID string) (chan *CTX, func()) {
	next := make(chan *CTX, 1)
	nextEmoticonList[botMsgID] = next
	return next, func() {
		close(next)
		delete(nextEmoticonList, botMsgID)
	}
}

func (ctx *CTX) emoticonNext() {
	if len(nextEmoticonList) == 0 {
		return
	}
	if c, ok := nextEmoticonList[ctx.Event.AddQuickEmoticon.BotMsgID]; ok {
		c <- ctx
	}
}
