package langx

import (
	"context"
)

type LangError struct {
	ctx    context.Context
	key    string
	format map[string]string
}

type errorConst string

const errorLang errorConst = "errorLang"

func (e *LangError) Error() string {
	errLang := e.ctx.Value(errorLang)
	l := ""
	if errLang != nil {
		l = string(errLang.(errorConst))
	}
	return GetFormat(l, e.key, e.format)
}

func (e *LangError) Code() int {
	return GetCode(e.key)
}

func NewErrorFormat(ctx context.Context, key string, format map[string]string) error {
	return &LangError{
		ctx:    ctx,
		key:    key,
		format: format,
	}
}

func NewError(ctx context.Context, key string) error {
	return &LangError{
		ctx: ctx,
		key: key,
	}
}

func SetCtxLang(ctx context.Context, lang string) context.Context {
	return context.WithValue(ctx, errorLang, lang)
}
