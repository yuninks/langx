package main

import (
	"context"

	"github.com/yuninks/langx"
)

func main() {

	err := ErrorWithMsg.Error()

	// 输出：错误
	println(err.Error())

}

type Language string

// 添加key+默认语言
func newLanguage(uniKey string, code int, defaultValue string) Language {
	langx.AppendCode(map[string]int{uniKey: code})
	langx.AppendTrans("zh_hans", map[string]string{uniKey: defaultValue})
	return Language(uniKey)
}

func (l Language) String() string {
	return string(l)
}

func (l Language) Error() error {
	return langx.NewError(context.Background(), l.String())
}

func (l Language) Errorf(format map[string]string) error {
	return langx.NewErrorFormat(context.Background(), l.String(), format)
}

var (
	Success Language = newLanguage("success", 200, "成功")

	Error        Language = newLanguage("error", 400, "错误")
	ErrorWithMsg Language = newLanguage("error_with_msg", 400, "错误 #msg#")
)
