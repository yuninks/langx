package langx

import "context"

// 定义错误常量

type ErrorLanguage string

// 生成错误常量
func NewLanguage(uniKey string, code int, defaultValue string) ErrorLanguage {
	AppendCode(map[string]int{uniKey: code})
	AppendTrans("zh_hans", map[string]string{uniKey: defaultValue})
	return ErrorLanguage(uniKey)
}

// 获取原key
func (l ErrorLanguage) String() string {
	return string(l)
}

// Key生成错误信息
func (l ErrorLanguage) Err() error {
	return NewError(context.Background(), l.String())
}

// Key生成错误信息
func (l ErrorLanguage) Errf(format map[string]string) error {
	return NewErrorf(context.Background(), l.String(), format)
}

// 获取翻译后的Code
func (l ErrorLanguage) Code() int {
	return GetCode(l.String())
}

// 获取翻译后的错误信息
func (l ErrorLanguage) Msg(ctx context.Context) string {
	return GetFormatCtx(ctx, l.String(), nil)
}

// 获取翻译后的错误信息
func (l ErrorLanguage) Msgf(ctx context.Context, format map[string]string) string {
	return GetFormatCtx(ctx, l.String(), format)
}

var (
	Success    ErrorLanguage = NewLanguage("success", 200, "操作成功")
	Error      ErrorLanguage = NewLanguage("error", 400, "操作失败")
	ErrWithMsg ErrorLanguage = NewLanguage("error_with_msg", 400, "操作失败: #msg#")
)
