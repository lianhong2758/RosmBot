package ctx

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/lianhong2758/RosmBot/web"
)

const (
	host           = "https://bbs-api.miyoushe.com"
	getRoomListURL = "/vila/api/bot/platform/getVillaGroupRoomList"
	getUserDataURL = "/vila/api/bot/platform/getMember"
	recallURL      = "/vila/api/bot/platform/recallMessage"
	deleteUserURL  = "/vila/api/bot/platform/deleteVillaMember"
)

// 获取房间列表
func (ctx *CTX) GetRoomList() (r *RoomList, err error) {
	data, err := web.Web(&http.Client{}, host+getRoomListURL, http.MethodGet, ctx.makeHeard, nil)
	if err != nil {
		return nil, err
	}
	r = new(RoomList)
	err = json.Unmarshal(data, r)
	return
}

// 获取用户信息
func (ctx *CTX) GetUserData(uid uint64) (r *UserData, err error) {
	data, _ := json.Marshal(H{"uid": uid})
	data, err = web.Web(&http.Client{}, host+getUserDataURL, http.MethodGet, ctx.makeHeard, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	r = new(UserData)
	err = json.Unmarshal(data, r)
	return
}

// 踢人
func (ctx *CTX) DeleteUser(uid uint64) (err error) {
	data, _ := json.Marshal(H{"uid": uid})
	data, err = web.Web(&http.Client{}, host+deleteUserURL, http.MethodPost, ctx.makeHeard, bytes.NewReader(data))
	var r ApiCode
	_ = json.Unmarshal(data, &r)
	if r.Retcode != 0 {
		return errors.New(r.Message)
	}
	return
}

// 撤回消息,消息id,房间id,发送时间
func (ctx *CTX) Recall(msgid string, msgtime, roomid int64) (err error) {
	data, _ := json.Marshal(H{"msg_uid": msgid, "room_id": roomid, "msg_time": msgtime})
	data, err = web.Web(&http.Client{}, host+recallURL, http.MethodPost, ctx.makeHeard, bytes.NewReader(data))
	var r ApiCode
	_ = json.Unmarshal(data, &r)
	if r.Retcode != 0 {
		return errors.New(r.Message)
	}
	return
}

type RoomList struct {
	Retcode int    `json:"retcode"`
	Message string `json:"message"`
	Data    struct {
		List []struct {
			GroupID   string `json:"group_id"`
			GroupName string `json:"group_name"`
			RoomList  []struct {
				RoomID   string `json:"room_id"`
				RoomName string `json:"room_name"`
				RoomType string `json:"room_type"`
				GroupID  string `json:"group_id"`
			} `json:"room_list"`
		} `json:"list"`
	} `json:"data"`
}

type UserData struct {
	Retcode int    `json:"retcode"`
	Message string `json:"message"`
	Data    struct {
		Member struct {
			Basic struct {
				UID       string `json:"uid"`
				Nickname  string `json:"nickname"`
				Introduce string `json:"introduce"`
				Avatar    string `json:"avatar"`
				AvatarURL string `json:"avatar_url"`
			} `json:"basic"`
			RoleIDList []string `json:"role_id_list"`
			JoinedAt   string   `json:"joined_at"`
			RoleList   []struct {
				ID       string `json:"id"`
				Name     string `json:"name"`
				Color    string `json:"color"`
				RoleType string `json:"role_type"`
				VillaID  string `json:"villa_id"`
				WebColor string `json:"web_color"`
			} `json:"role_list"`
		} `json:"member"`
	} `json:"data"`
}
