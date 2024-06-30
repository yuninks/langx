package langx_test

import (
	"context"
	"testing"

	"code.yun.ink/pkg/langx"
)

func TestError(t *testing.T) {
	var err error

	ctx := context.Background()

	langx.InitLangx(
		langx.SetDefaultCode(0),
		langx.SetDefaultLanguage("zh"),
	)
	langx.RegisterCode(map[string]int{
		"login_success": 200,
		"error":         400,
	})
	langx.RegisterTrans("zh", map[string]string{
		"login_success": "成功",
		"error":         "错误",
		"username":      "你好 #name#", // 有占位符
	})
	langx.RegisterTrans("en", map[string]string{
		"login_success": "success",
		"error":         "error",
		"username":      "Hello #name#", // 有占位符
	})

	err = langx.NewError(ctx, "error")
	// fmt.Printf("err: %v\n", err)
	t.Log(err.Error())
	val, ok := err.(*langx.LangError)
	if ok {
		t.Log(val.Code())
	}

}
