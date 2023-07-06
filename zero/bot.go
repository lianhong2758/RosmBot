package zero

import ()

// 默认值
var MYSconfig = MYSCFG{
	BotToken: Token{
		VillaID: []string{"123456"},
		Master:  []string{"123456"},
	}}

type Token struct {
	VillaID   []string `json:"villa_id"`
	Master    []string `json:"master_id"`
	BotID     string   `json:"bot_id"`
	BotSecret string   `json:"bot_secret"`
	BotName   string   `json:"bot_name"`
}
type MYSCFG struct {
	BotToken  Token  `json:"token"`
	EventPath string `json:"eventpath,omitempty"`
	Port      string `json:"port,omitempty"`
	Host      string `json:"host,omitempty"`
	Types     int    `json:"types"`
	Key       string `json:"key,omitempty"`
}
