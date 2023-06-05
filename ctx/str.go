package ctx

import ()

// 回调的请求结构
type infoSTR struct {
	Event struct {
		Robot struct { 
			Template tem `json:"template"`
			VillaID  int `json:"villa_id"`
		} `json:"robot"`
		Type       int `json:"type"`
		ExtendData struct {
			EventData struct {
				SendMessage      sendmessage      `json:"SendMessage"`
				JoinVilla        joinVilla        `json:"JoinVilla"`
				CreateRobot      changeRobot      `json:"CreateRobot"`
				DeleteRobot      changeRobot      `json:"DeleteRobot"`
				AddQuickEmoticon addQuickEmoticon `json:"AddQuickEmoticon"`
				AuditCallback    auditCallback    `json:"AuditCallback"`
			} `json:"EventData"`
		} `json:"extend_data"`
		CreatedAt int64  `json:"created_at"`
		ID        string `json:"id"`
		SendAt    int    `json:"send_at"`
	} `json:"event"`
}
type tem struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	Icon     string `json:"icon"`
	Commands []struct {
		Name string `json:"name"`
		Desc string `json:"desc"`
	} `json:"commands"`
}

type sendmessage struct {
	Content    string `json:"content"`
	FromUserID int    `json:"from_user_id"`
	SendAt     int64  `json:"send_at"`
	ObjectName int    `json:"object_name"`
	RoomID     int    `json:"room_id"`
	Nickname   string `json:"nickname"`
	MsgUID     string `json:"msg_uid"`
}

type joinVilla struct {
	JoinUID          uint64 `json:"join_uid"`
	JoinUserNickname string `json:"join_user_nickname"`
	JoinAt           int64  `json:"join_at"`
}

type changeRobot struct {
	VillaID int `json:"villa_id"`
}

type addQuickEmoticon struct {
	VillaID    int    `json:"villa_id"`
	RoomID     int    `json:"room_id"`
	UID        int    `json:"uid"`
	EmoticonID int    `json:"emoticon_id"`
	Emoticon   string `json:"emoticon"`
	MsgUID     string `json:"msg_uid"`
	BotMsgID   string `json:"bot_msg_id"`
	IsCancel   bool   `json:"is_cancel"`
}
type auditCallback struct {
	AuditID     string `json:"audit_id"`
	BotTplID    string `json:"bot_tpl_id"`
	VillaID     int    `json:"villa_id"`
	RoomID      int    `json:"room_id"`
	UserID      int    `json:"user_id"`
	PassThrough string `json:"pass_through"`
	AuditResult int    `json:"audit_result"`
}

type other struct{}

// 快捷消息结构
type CTX struct {
	Mess  *user
	Event *sendmessage
	Other *other
	Bot   *tem
	Being *being
}
type user struct {
	Trace struct {
		VisualRoomVersion string `json:"visual_room_version"`
		AppVersion        string `json:"app_version"`
		ActionType        int    `json:"action_type"`
		BotMsgID          string `json:"bot_msg_id"`
		Client            string `json:"client"`
		Env               string `json:"env"`
		RongSdkVersion    string `json:"rong_sdk_version"`
	} `json:"trace"`
	MentionedInfo struct {
		MentionedContent string   `json:"mentionedContent"`
		UserIDList       []string `json:"userIdList"`
		Type             int      `json:"type"`
	} `json:"mentionedInfo"`
	User struct {
		PortraitURI string `json:"portraitUri"`
		Extra       string `json:"extra"`
		Name        string `json:"name"`
		Alias       string `json:"alias"`
		ID          string `json:"id"`
		Portrait    string `json:"portrait"`
	} `json:"user"`
	Content struct {
		Images   []interface{} `json:"images"`
		Entities []struct {
			Offset int `json:"offset"`
			Length int `json:"length"`
			Entity struct {
				Type  string `json:"type"`
				BotID string `json:"bot_id"`
			} `json:"entity"`
		} `json:"entities"`
		Text string `json:"text"`
	} `json:"content"`
}

type sendState struct {
	Retcode int    `json:"retcode"`
	Message string `json:"message"`
	Data    struct {
		BotMsgID string `json:"bot_msg_id"`
	} `json:"data"`
}

type being struct {
	RoomID  int
	VillaID int
	Word    string
	Rex     []string
}
