package zero

import ()

var BotToken *Token

type Token struct {
	VillaID   []string `json:"villa_id"`
	BotID     string   `json:"bot_id"`
	BotSecret string   `json:"bot_secret"`
	BotName   string   `json:"bot_name"`
}
