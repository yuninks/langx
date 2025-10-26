package langx

import "context"

const ctxLangKey string = "ctxLang"

func SetCtxLang(ctx context.Context, lang string) context.Context {
	return context.WithValue(ctx, ctxLangKey, lang)
}

func GetCtxLang(ctx context.Context) string {
	// 指定的
	lang := ctx.Value(ctxLangKey)

	l := ""

	if lang != nil {
		l = lang.(string)
	}

	// HTTP头部的
	if l == "" {
		lang = ctx.Value("Accept-Language")
		if lang != nil {
			l = lang.(string)
		}
	}

	// 默认的
	if l == "" {
		l = GetDefaultLang()
	}

	return l
}
