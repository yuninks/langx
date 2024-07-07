package langx

import (
	"fmt"
	"strings"
)

type langx struct {
	ops      *options
	codeMap  map[string]int
	transMap map[string]map[string]string
}

var l *langx = &langx{}

func init() {

	o := defaultOptions()

	l = &langx{
		ops:      o,
		codeMap:  make(map[string]int),
		transMap: make(map[string]map[string]string),
	}
}

func InitLangx(ops ...Option) {
	for _, opt := range ops {
		opt(l.ops)
	}
}

// 这是语言的Code
func RegisterCode(datas map[string]int) {
	l.codeMap = datas
}

// 注册语言翻译
func RegisterTrans(langName string, trans map[string]string) {
	l.transMap[langName] = trans
}

// 获取翻译
// 包含Code和Message
func GetTrans(lang string, key string, format map[string]string) (code int, str string) {
	code = GetCode(key)
	str = GetFormat(lang, key, format)
	return
}

// 根据Key获取code
func GetCode(key string) int {
	code, ok := l.codeMap[key]
	if !ok {
		return l.ops.defaultCode
	}
	return code
}

func GetMsg(lang string, key string) string {
	// 找指定语言
	str, ok := l.transMap[lang]
	if ok {
		val, ok := str[key]
		if ok {
			return val
		}
	}
	// 找默认语言
	str, ok = l.transMap[l.ops.defaultLang]
	if ok {
		val, ok := str[key]
		if ok {
			return val
		}
	}

	return key
}

// 拼接回复
func GetFormat(lang string, key string, arr map[string]string) string {
	str := GetMsg(lang, key)
	for k, v := range arr {
		str = strings.ReplaceAll(str, fmt.Sprintf(l.ops.replaceKey, k), v)
	}
	return str
}

func GetDefaultCode() int {
	return l.ops.defaultCode
}
