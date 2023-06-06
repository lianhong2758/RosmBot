package zero

import ()

var MYSconfig MYSCFG

type Token struct {
	VillaID   []int  `json:"villa_id"`
	BotID     string `json:"bot_id"`
	BotSecret string `json:"bot_secret"`
	BotName   string `json:"bot_name"`
}
type MYSCFG struct {
	BotToken  Token  `json:"token"`
	EventPath string `json:"eventpath"`
	Port      string `json:"port"`
	Host      string `json:"host"`
}
