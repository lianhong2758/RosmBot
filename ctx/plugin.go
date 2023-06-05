package ctx

import (
	"log"
	"regexp"
)

// 插件注册
var plugins = map[string]*PluginData{}

// 全匹配字典
var caseAllWord = map[string]func(c *CTX){}

// 正则字典
var caseRegexp = map[*regexp.Regexp]func(c *CTX){}

func Register(pluginName string, p *PluginData) *PluginData {
	plugins[pluginName] = p
	return plugins[pluginName]
}

func (p *PluginData) AddWord(f func(c *CTX), word ...string) {
	p.Word = append(p.Word, word...)
	for _, v := range word {
		caseAllWord[v] = f
	}
}

func (p *PluginData) AddRex(f func(c *CTX), rex string) {
	r := regexp.MustCompile(rex)
	p.Rex = append(p.Rex, r)
	caseRegexp[r] = f
}

func Display() {
	log.Println(caseAllWord)
}

type PluginData struct {
	Word []string
	Rex  []*regexp.Regexp
	Help string
	Name string
}
