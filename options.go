package langx

type options struct {
	defaultCode int
	defaultLang string
	replaceKey  string
	ctxLangKey  string
}

func defaultOptions() *options {
	return &options{
		defaultCode: 200,
		defaultLang: "zh",
		replaceKey:  "#%s#",
		ctxLangKey:  "language",
	}
}

// 设置
func InitLangx(ops ...Option) {
	for _, opt := range ops {
		opt(l.ops)
	}
}

type Option func(*options)

// 设置默认的code
func SetDefaultCode(code int) Option {
	return func(o *options) {
		o.defaultCode = code
	}
}

// 从ctx里面获取语言的key
func SetCtxLangKey(key string) Option {
	return func(o *options) {
		o.ctxLangKey = key
	}
}

// 默认语言
func SetDefaultLanguage(lang string) Option {
	return func(o *options) {
		o.defaultLang = lang
	}
}

// 替换规则 %s 为占位符 需要填进去
func SetReplaceKey(key string) Option {
	return func(o *options) {
		o.replaceKey = key
	}
}
