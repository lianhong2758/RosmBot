package ctx

import (
	"os"
	"regexp"

	"github.com/FloatTech/floatbox/file"
	log "github.com/sirupsen/logrus"
)

const (
	Join   = "join"
	Create = "create"
	Delete = "delete"
	Quick  = "quick"
)

// 插件注册
var (
	plugins = map[string]*PluginData{}

	// 全匹配字典
	caseAllWord = map[string]*Matcher{}

	// 正则字典
	caseRegexp = map[*regexp.Regexp]*Matcher{}

	//事件触发
	caseOther = map[string][]*Matcher{
		Join:   {},
		Create: {},
		Delete: {},
		Quick:  {},
	}
)

// 注册插件
func Register(pluginName string, p *PluginData) *PluginData {
	log.Debugln("插件注册:", pluginName)
	plugins[pluginName] = p
	if file.IsNotExist(p.DataFolder) && p.DataFolder != "" {
		_ = os.MkdirAll("data/"+p.DataFolder, 0755)
	}
	plugins[pluginName].DataFolder = "data/" + p.DataFolder + "/"
	return plugins[pluginName]
}

// 完全词匹配
func (p *PluginData) AddWord(word ...string) *Matcher {
	m := new(Matcher)
	m.Block = true
	m.Word = append(m.Word, word...)
	for _, v := range word {
		caseAllWord[v] = m
	}
	p.Matchers = append(p.Matchers, m)
	m.PluginNode = p
	return m
}

// 正则匹配
func (p *PluginData) AddRex(rex string) *Matcher {
	m := new(Matcher)
	m.Block = true
	r := regexp.MustCompile(rex)
	m.Rex = append(m.Rex, r)
	caseRegexp[r] = m
	p.Matchers = append(p.Matchers, m)
	m.PluginNode = p
	return m
}

// 其他事件匹配器
func (p *PluginData) AddOther(types string) *Matcher {
	m := new(Matcher)
	m.Block = false
	if _, ok := caseOther[types]; ok {
		caseOther[types] = append(caseOther[types], m)
	} else {
		log.Errorln("插件载入失败: ", p.Name, "-", types, "#不存在的事件类型")
	}
	p.Matchers = append(p.Matchers, m)
	m.PluginNode = p
	return m
}

// 注册Handle
func (m *Matcher) Handle(h Handler) {
	m.Handler = h
}

// 阻断器
func (m *Matcher) SetBlock(ok bool) *Matcher {
	m.Block = ok
	return m
}

func (m *Matcher) Rule(r ...Rule) *Matcher {
	m.Rules = append(m.Rules, r...)
	return m
}
func Display() {
	log.Println(caseAllWord)
	log.Println(caseRegexp)
}
func GetPlugins() map[string]*PluginData {
	return plugins
}
