package langx_test

import (
	"testing"

	"github.com/yuninks/langx"
)

const(
	Lang string = "s"
)

var MapCode = map[string]int{
	Lang:200,
}


func TestLangx(t *testing.T) {
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
	// 获取翻译码
	code, msg := langx.GetTrans("zh", "login_success", nil)
	t.Log(code, msg)
	code, msg = langx.GetTrans("en", "error", nil)
	t.Log(code, msg)

	// 获取翻译码，有占位符
	code, msg = langx.GetTrans("zh", "username", map[string]string{
		"name": "张三",
	})
	t.Log(code, msg)

}
