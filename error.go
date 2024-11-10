package langx

import (
	"context"
)

type LangError interface {
	Error() string                // 实现error接口&获取翻译后的错误信息
	GetCode() int                 // 获取翻译后的Code
	GetKey() string               // 获取原Key值
	GetFormat() map[string]string // 获取附加数据
	SetLang(lang string)          // 设置语言
}

type langError struct {
	ctx    context.Context
	key    string
	format map[string]string
}

func (l *langError) Error() string {
	return GetFormat(GetCtxLang(l.ctx), l.key, l.format)
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

func (e *langError) SetLang(lang string) {
	e.ctx = SetCtxLang(e.ctx, lang)
}

func NewErrorf(ctx context.Context, key string, format map[string]string) error {
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
