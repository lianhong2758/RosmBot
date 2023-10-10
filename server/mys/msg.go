package mys

import (
	"github.com/lianhong2758/RosmBot-MUL/message"
	"github.com/lianhong2758/RosmBot-MUL/rosm"
)

func init() {
	var cfg = &Config{}
	var runner rosm.Boter
	runner = cfg
	if err := runner.Run(); err != nil {
		panic(err)
	}
}

type Config struct{}

func (c *Config) Run() error {
	return nil
}
func (c *Config) Name() string {
	return ""
}
func (c *Config) Message() any {
	return ""
}
func (c *Config) Config() any {
	return ""
}
func (c *Config) Send(msg ...message.MessageSegment) any {
	return ""
}
func (c *Config) MakeCTX(Message any) *rosm.CTX {
	return nil
}
