package ctx

import (
	"net/http"

	"RosmBot/web"
)

const (
	GETRoomList = "https://bbs-api.miyoushe.com/vila/api/bot/platform/getVillaGroupRoomList"
)

func (ctx *CTX) GetRoomList() ([]byte, error) {
	return web.Web(&http.Client{}, GETRoomList, http.MethodGet, ctx.makeHeard, nil)
}
