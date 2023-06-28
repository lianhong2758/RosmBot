package zero

import ()

// 默认值
var MYSconfig = MYSCFG{
	BotToken: Token{
		VillaID: []string{"123456"},
		Master:  []string{"123456"},
	},
	EventPath: "/",
	Port:      "0.0.0.0:80",
	Host:      "127.0.0.1:80",
}

type Token struct {
	VillaID   []string `json:"villa_id"`
	Master    []string `json:"master_id"`
	BotID     string   `json:"bot_id"`
	BotSecret string   `json:"bot_secret"`
	BotName   string   `json:"bot_name"`
}
type MYSCFG struct {
	BotToken  Token  `json:"token"`
	EventPath string `json:"eventpath"`
	Port      string `json:"port"`
	Host      string `json:"host"`
	Types     int    `json:"types"`
	Key       string `json:"key,omitempty"`
}
