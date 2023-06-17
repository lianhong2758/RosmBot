package zero

import ()

var MYSconfig = MYSCFG{
	BotToken:  Token{},
	EventPath: "/",
	Port:      "0.0.0.0:80",
	Host:      "127.0.0.1:80",
}

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
