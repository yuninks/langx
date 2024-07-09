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

type Option func(*options)

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

// 替换规则 %s 为占位符
func SetReplaceKey(key string) Option {
	return func(o *options) {
		o.replaceKey = key
	}
}
