package ctx

import (
	"log"
	"os"
	"regexp"

	"github.com/FloatTech/floatbox/file"
)

// 插件注册
var (
	plugins = map[string]*PluginData{}

	// 全匹配字典
	caseAllWord = map[string]func(ctx *CTX){}

	// 正则字典
	caseRegexp = map[*regexp.Regexp]func(ctx *CTX){}

	//事件触发
	caseOther = map[string][]func(ctx *CTX){
		"join":   {},
		"create": {},
		"delete": {},
		"quick":  {},
	}
)

// 注册插件
func Register(pluginName string, p *PluginData) *PluginData {
	plugins[pluginName] = p
	if file.IsNotExist(p.DataFolder) && p.DataFolder != "" {
		_ = os.MkdirAll("data/"+p.DataFolder, 0755)
	}
	plugins[pluginName].DataFolder = "data/" + p.DataFolder + "/"
	return plugins[pluginName]
}

// 完全词匹配
func (p *PluginData) AddWord(f func(ctx *CTX), word ...string) {
	p.Word = append(p.Word, word...)
	for _, v := range word {
		caseAllWord[v] = f
	}
}

// 正则匹配
func (p *PluginData) AddRex(f func(ctx *CTX), rex string) {
	r := regexp.MustCompile(rex)
	p.Rex = append(p.Rex, r)
	caseRegexp[r] = f
}

// 其他事件匹配器
func (p *PluginData) AddOther(f func(ctx *CTX), types string) {
	if _, ok := caseOther[types]; ok {
		caseOther[types] = append(caseOther[types], f)
	} else {
		log.Panicln("插件载入失败: ", p.Name, "-", types, "#不存在的事件类型")
	}
}
func Display() {
	log.Println(caseAllWord)
	log.Println(caseRegexp)
}
func GetPlugins() map[string]*PluginData {
	return plugins
}
