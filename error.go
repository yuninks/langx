package langx

import (
	"context"
)

type LangError interface {
	NewError(ctx context.Context, key string) error                                 // New错误
	NewErrorFormat(ctx context.Context, key string, format map[string]string) error // New错误
	Error() string                                                                  // 实现error接口&获取翻译后的错误信息
	GetCode() int                                                                   // 获取翻译后的Code
	GetKey() string                                                                 // 获取原Key值
	GetFormat() map[string]string                                                   // 获取附加数据
}

type langError struct {
	ctx    context.Context
	key    string
	format map[string]string
}

type errorConst string

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

func (e *langError) GetKey() string {
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

const errorLang errorConst = "errorLang"

func SetCtxLang(ctx context.Context, lang string) context.Context {
	return context.WithValue(ctx, errorLang, lang)
}
