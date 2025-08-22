package langx

import (
	"context"
)

type LangError interface {
	Copy() LangError                    // 复制错误信息
	Error() string                      // 实现error接口&获取翻译后的错误信息
	GetCode() int                       // 获取翻译后的Code
	GetKey() string                     // 获取原Key值
	GetFormat() map[string]string       // 获取附加数据
	SetFormat(format map[string]string) // 设置附加数据
	SetCtx(ctxhttp context.Context)     // 设置上下文
}

type langError struct {
	ctx    context.Context
	key    string
	format map[string]string
}

func (l *langError) Copy() LangError {
	return &langError{
		ctx:    l.ctx,
		key:    l.key,
		format: l.format,
	}
}

func (e *langError) Error() string {
	errLang := e.ctx.Value("Accept-Language")
	l := ""
	if errLang != nil {
		l = string(errLang.(string))
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

func (e *langError) SetFormat(format map[string]string) {
	e.format = format
}

func (e *langError) SetCtx(ctx context.Context) {
	e.ctx = ctx
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
	return context.WithValue(ctx, "Accept-Language", lang)
}
