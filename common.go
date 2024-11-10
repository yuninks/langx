package langx

import "context"

const ctxLangKey string = "ctxLang"

func SetCtxLang(ctx context.Context, lang string) context.Context {
	return context.WithValue(ctx, ctxLangKey, lang)
}

func GetCtxLang(ctx context.Context) string {
	lang, ok := ctx.Value(ctxLangKey).(string)
	if !ok {
		return ""
	}
	return lang
}
