package langx

import (
	"context"
)

// Key生成错误信息

type LangError interface {
	Copy() LangError                    // 复制一个新的错误信息
	Error() string                      // 实现error接口&获取翻译后的错误信息
	GetCode() int                       // 获取翻译后的Code
	GetKey() string                     // 获取原Key值
	GetFormat() map[string]string       // 获取附加数据
	SetFormat(format map[string]string) // 设置附加数据
	SetCtx(ctxHttp context.Context)     // 设置上下文
	SetLang(lang string)                // 设置语言
	SetFormatKV(key, value string)      // 设置附加数据键值对
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

func (e *langError) SetFormat(format map[string]string) {
	e.format = format
}

func (l *langError) SetFormatKV(key, value string) {
	if l.format == nil {
		l.format = make(map[string]string)
	}
	l.format[key] = value
}

func (e *langError) SetCtx(ctxHttp context.Context) {
	e.ctx = ctxHttp
}

func (l *langError) Error() string {
	errLang := l.ctx.Value("Accept-Language")
	lang := ""
	if errLang != nil {
		lang = errLang.(string)
	}
	if lang == "" {
		lang = GetDefaultLang()
	}

	return GetFormat(lang, l.key, l.format)
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

func NewErrorStruct(ctx context.Context, key string, format map[string]string) LangError {
	return &langError{
		ctx:    ctx,
		key:    key,
		format: format,
	}
}
