package ctx

import "strconv"

var nextList = map[int]chan *CTX{}

// 获取本房间全体的下一句话
func (ctx *CTX) GetNextAllMess() (chan *CTX, func()) {
	next := make(chan *CTX, 1)
	id := ctx.Being.VillaID + ctx.Being.RoomID
	nextList[id] = next
	return next, func() {
		close(next)
		delete(nextList, id)
	}
}

// 获取本房间该用户的下一句话
func (ctx *CTX) GetNextUserMess() (chan *CTX, func()) {
	next := make(chan *CTX, 1)
	id := ctx.Being.VillaID + ctx.Being.RoomID + ctx.IntUserID()
	nextList[id] = next
	return next, func() {
		close(next)
		delete(nextList, id)
	}
}

func (ctx *CTX) SendNext() {
	//先匹配个人
	if c, ok := nextList[ctx.Being.VillaID+ctx.Being.RoomID+ctx.IntUserID()]; ok {
		c <- ctx
	}
	if c, ok := nextList[ctx.Being.VillaID+ctx.Being.RoomID]; ok {
		c <- ctx
	}
}

func (ctx *CTX) IntUserID() int { x, _ := strconv.Atoi(ctx.Being.User.ID); return x }
