package langx

import "context"

// 定义错误常量

type ErrorLanguage struct {
	LangError
}

// 生成错误常量
func NewLanguage(uniKey string, code int, defaultValue string) ErrorLanguage {
	AppendCode(map[string]int{uniKey: code})
	AppendTrans(GetDefaultLang(), map[string]string{uniKey: defaultValue})

	l := NewStruct(context.Background(), uniKey, nil)

	return ErrorLanguage{l}
}

// Key生成错误信息
func (l ErrorLanguage) Err() error {
	return l
}

// Key生成错误信息
func (l ErrorLanguage) Errf(format map[string]string) error {
	newLang := l.Copy()
	newLang.SetFormat(format)
	return newLang
}

func (l ErrorLanguage) ErrfKV(key, value string) error {
	newLang := l.Copy()
	newLang.SetFormatKV(key, value)
	return newLang
}

// 获取翻译后的错误信息
func (l ErrorLanguage) Msg(ctx context.Context) string {
	newLang := l.Copy()
	newLang.SetCtx(ctx)
	return newLang.Error()
}

// 获取翻译后的错误信息
func (l ErrorLanguage) Msgf(ctx context.Context, format map[string]string) string {
	newLang := l.Copy()
	newLang.SetCtx(ctx)
	newLang.SetFormat(format)
	return newLang.Error()
}

var (
	Success    ErrorLanguage = NewLanguage("success", 200, "操作成功")
	Error      ErrorLanguage = NewLanguage("error", 400, "操作失败")
	ErrWithMsg ErrorLanguage = NewLanguage("error_with_msg", 400, "操作失败: #msg#")
)
