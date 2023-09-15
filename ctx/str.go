package ctx

import "regexp"

type (
	// Rule filter the event
	Rule func(ctx *CTX) bool
	// Handler 事件处理函数
	Handler func(ctx *CTX)
)

// 回调的请求结构
type infoSTR struct {
	Event struct {
		Robot struct {
			Template tem `json:"template"` // 机器人模板信息
			VillaID  int `json:"villa_id"` // 事件所属的大别野 id
		} `json:"robot"`
		Type       int      `json:"type"`
		ExtendData struct { // 包含事件的具体数据
			EventData EventData `json:"EventData"`
		} `json:"extend_data"`
		CreatedAt int64  `json:"created_at"`
		ID        string `json:"id"`
		SendAt    int    `json:"send_at"`
	} `json:"event"`
}

// 所有事件
type EventData struct {
	SendMessage      sendmessage      `json:"SendMessage"`
	JoinVilla        joinVilla        `json:"JoinVilla"`
	CreateRobot      changeRobot      `json:"CreateRobot"`
	DeleteRobot      changeRobot      `json:"DeleteRobot"`
	AddQuickEmoticon addQuickEmoticon `json:"AddQuickEmoticon"`
	AuditCallback    auditCallback    `json:"AuditCallback"`
}

// 机器人相关信息
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

// 用户@机器人发送消息
type sendmessage struct {
	Content    string `json:"content"`
	FromUserID int    `json:"from_user_id"`
	SendAt     int64  `json:"send_at"`
	RoomID     int    `json:"room_id"`
	ObjectName int    `json:"object_name"`
	Nickname   string `json:"nickname"`
	MsgUID     string `json:"msg_uid"`
	BotMsgID   string `json:"bot_msg_id"`
	VillaID    int    `json:"villa_id"`
}

// 有新用户加入大别野
type joinVilla struct {
	JoinUID          int    `json:"join_uid"`
	JoinUserNickname string `json:"join_user_nickname"`
	JoinAt           int64  `json:"join_at"`
}

// 大别野添加机器人实例,大别野删除机器人实例
type changeRobot struct {
	VillaID int `json:"villa_id"`
}

// 用户使用表情回复消息表态
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

// 审核结果回调
type auditCallback struct {
	AuditID     string `json:"audit_id"`
	BotTplID    string `json:"bot_tpl_id"`
	VillaID     int    `json:"villa_id"`
	RoomID      int    `json:"room_id"`
	UserID      int    `json:"user_id"`
	PassThrough string `json:"pass_through"`
	AuditResult int    `json:"audit_result"`
}

// 快捷消息结构
type CTX struct {
	Mess  *mess
	Event *EventData
	Bot   *tem
	Being *being
}

// 常用数据
type being struct {
	RoomID  int
	VillaID int
	User    *user
	ATList  []string
	Word    string
	Rex     []string
}

// 接收的原始消息,解析
type mess struct {
	Trace struct {
		VisualRoomVersion string `json:"visual_room_version"`
		AppVersion        string `json:"app_version"`
		ActionType        int    `json:"action_type"`
		BotMsgID          string `json:"bot_msg_id"`
		Client            string `json:"client"`
		Env               string `json:"env"`
		RongSdkVersion    string `json:"rong_sdk_version"`
	} `json:"trace"`
	Quote struct {
		QuotedMessageSendTime   int64  `json:"quoted_message_send_time"`
		OriginalMessageID       string `json:"original_message_id"`
		OriginalMessageSendTime int64  `json:"original_message_send_time"`
		QuotedMessageID         string `json:"quoted_message_id"`
	} `json:"quote"`
	MentionedInfo struct {
		MentionedContent string   `json:"mentionedContent"`
		UserIDList       []string `json:"userIdList"`
		Type             int      `json:"type"`
	} `json:"mentionedInfo"`
	User    user    `json:"user"`
	Content content `json:"content"`
}

type content struct {
	Images   []any `json:"images"`
	Entities []struct {
		Offset int `json:"offset"`
		Length int `json:"length"`
		Entity struct {
			Type  string `json:"type"`
			BotID string `json:"bot_id"`
		} `json:"entity"`
	} `json:"entities"`
	Text string `json:"text"`
}
type user struct {
	PortraitURI string `json:"portraitUri"`
	Extra       string `json:"extra"`
	Name        string `json:"name"`
	Alias       string `json:"alias"`
	ID          string `json:"id"`
	Portrait    string `json:"portrait"`
}

// 消息发送回调
type SendState struct {
	ApiCode
	Data struct {
		BotMsgID string `json:"bot_msg_id"`
	} `json:"data"`
}

// api返回
type ApiCode struct {
	Retcode int    `json:"retcode"`
	Message string `json:"message"`
}

type Message []MessageSegment
type MessageSegment struct {
	Type string `json:"type"`
	Data H      `json:"data"`
}

// 消息模板
type Content struct {
	//图片
	ImageStr
	//文本
	Text     string     `json:"text,omitempty"`
	Entities []Entities `json:"entities,omitempty"`
	Images   []ImageStr `json:"images,omitempty"`
	//链接预览
	Preview *PreviewStr `json:"preview_link,omitempty"`
	//下标
	Badge *BadgeStr `json:"badge,omitempty"`
}

type ImageStr struct {
	URL      string `json:"url,omitempty"`
	FileSize int    `json:"file_size,omitempty"`
	Size     *Size  `json:"size,omitempty"`
}
type Size struct {
	Height int `json:"height,omitempty"`
	Width  int `json:"width,omitempty"`
}

type Entities struct {
	Entity H   `json:"entity,omitempty"`
	Length int `json:"length,omitempty"`
	Offset int `json:"offset,omitempty"`
}

type MentionedInfoStr struct {
	Type       int      `json:"type"`
	UserIDList []string `json:"userIdList"`
}

// 预览图片的结构
type PreviewStr struct {
	URL        string `json:"url,omitempty"`              //跳转
	ImageURL   string `json:"image_url,omitempty"`        //大图
	IsIntLink  bool   `json:"is_internal_link,omitempty"` //link
	Title      string `json:"title,omitempty"`            //右上角大标题
	Content    string `json:"content,omitempty"`          //中间小字
	Icon       string `json:"icon_url,omitempty"`         //左下角小图
	SourceName string `json:"source_name,omitempty"`      //左下角图标名
}

// 左下标的结构
type BadgeStr struct {
	Icon string `json:"icon_url,omitempty"` //图标
	Text string `json:"text,omitempty"`     //文字
	URL  string `json:"url,omitempty"`      //跳转
}

type PluginData struct {
	Help       string
	Name       string
	DataFolder string //"data/xxx/"+
	Matchers   []*Matcher
}

type Matcher struct {
	Word       []string
	Rex        []*regexp.Regexp
	Rules      []Rule
	Handler    Handler
	Block      bool        //阻断
	PluginNode *PluginData //溯源
}
