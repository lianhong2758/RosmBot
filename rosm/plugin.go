package rosm

import (
	"os"
	"regexp"

	"github.com/FloatTech/floatbox/file"
	"github.com/lianhong2758/RosmBot-MUL/tool/rate"
	log "github.com/sirupsen/logrus"
)

const (
	Join           = "join"
	Create         = "create"
	Delete         = "delete"
	Quick          = "quick"
	AllMessage     = "all"
	SurplusMessage = "surplus"
)

type (
	// Rule filter the event
	Rule func(ctx *CTX) bool
	// Handler 事件处理函数
	Handler func(ctx *CTX)
)

type PluginData struct {
	Help       string
	Name       string
	DataFolder string //"data/xxx/"+
	Matchers   []*Matcher
}
type Matcher struct {
	Word       []string
	Rex        []*regexp.Regexp
	rules      []Rule
	handler    Handler
	block      bool        //阻断
	PluginNode *PluginData //溯源
}

// 插件注册
var (
	plugins = map[string]*PluginData{}

	// 全匹配字典
	caseAllWord = map[string]*Matcher{}

	// 正则字典
	caseRegexp = map[*regexp.Regexp]*Matcher{}

	//事件触发
	caseOther = map[string][]*Matcher{
		Join:           {},
		Create:         {},
		Delete:         {},
		Quick:          {},
		AllMessage:     {},
		SurplusMessage: {},
	}
)

// 注册插件
func Register(p *PluginData) *PluginData {
	pluginName := p.Name
	log.Debugln("插件注册:", pluginName)
	plugins[pluginName] = p
	if p.DataFolder != "" && file.IsNotExist(p.DataFolder) {
		_ = os.MkdirAll("data/"+p.DataFolder, 0755)
	}
	plugins[pluginName].DataFolder = "data/" + p.DataFolder + "/"
	return plugins[pluginName]
}

// 创建插件对象信息
func NewRegist(name, help, dataFolder string) *PluginData {
	return &PluginData{Name: name, Help: help, DataFolder: dataFolder}
}

// 完全词匹配
func (p *PluginData) AddWord(word ...string) *Matcher {
	m := new(Matcher)
	m.block = true
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
	m.block = true
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
	m.block = false
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
	m.handler = h
}

// 阻断器
func (m *Matcher) SetBlock(ok bool) *Matcher {
	m.block = ok
	return m
}

func (m *Matcher) Rule(r ...Rule) *Matcher {
	m.rules = append(m.rules, r...)
	return m
}

// Limit 限速器
// postfn 当请求被拒绝时的操作
func (m *Matcher) Limit(limiterfn func(*CTX) *rate.Limiter, postfn ...func(*CTX)) *Matcher {
	m.rules = append(m.rules, func(ctx *CTX) bool {
		if limiterfn(ctx).Acquire() {
			return true
		}
		if len(postfn) > 0 {
			for _, fn := range postfn {
				fn(ctx)
			}
		}
		return false
	})
	return m
}

func Display() {
	log.Println(caseAllWord)
	log.Println(caseRegexp)
}
func GetPlugins() map[string]*PluginData {
	return plugins
}
