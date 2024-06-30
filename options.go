package langx

type options struct {
	defaultCode int
	defaultLang string
	replaceKey  string
}

func defaultOptions() *options {
	return &options{
		defaultCode: 200,
		defaultLang: "zh",
		replaceKey:  "#%s#",
	}
}

type Option func(*options)

func SetDefaultCode(code int) Option {
	return func(o *options) {
		o.defaultCode = code
	}
}

func SetDefaultLanguage(lang string) Option {
	return func(o *options) {
		o.defaultLang = lang
	}
}

func SetReplaceKey(key string) Option {
	return func(o *options) {
		o.replaceKey = key
	}
}
