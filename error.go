package langx

import (
	"context"
)

type LangError interface {
	NewError(ctx context.Context, key string) error
	NewErrorFormat(ctx context.Context, key string, format map[string]string) error
	Error() string
	GetCode() int
	GetMsg() string
	GetFormat() map[string]string
}

type langError struct {
	ctx    context.Context
	key    string
	format map[string]string
}

type errorConst string

const errorLang errorConst = "errorLang"

func (e *langError) Error() string {
	errLang := e.ctx.Value(errorLang)
	l := ""
	if errLang != nil {
		l = string(errLang.(errorConst))
	}
	return GetFormat(l, e.key, e.format)
}

func (e *langError) GetCode() int {
	return GetCode(e.key)
}

func (e *langError) GetMsg() string {
	return e.key
}

func (e *langError) GetFormat() map[string]string {
	if e.format == nil {
		e.format = make(map[string]string)
	}
	return e.format
}

func NewErrorFormat(ctx context.Context, key string, format map[string]string) error {
	return &langError{
		ctx:    ctx,
		key:    key,
		format: format,
	}
}

func NewError(ctx context.Context, key string) error {
	return &langError{
		ctx: ctx,
		key: key,
	}
}

func SetCtxLang(ctx context.Context, lang string) context.Context {
	return context.WithValue(ctx, errorLang, lang)
}
