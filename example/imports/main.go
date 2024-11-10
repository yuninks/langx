package main

import (
	"embed"
	"fmt"

	"github.com/yuninks/langx"
)

//go:embed lang
var assetsFs embed.FS

func main() {
	regByAppend()
}

// 导入语言包 基于Append
func regByAppend() {
	langx.AppendCode(map[string]int{
		"success": 200,
	})
	langx.AppendTrans("zh-CN", map[string]string{
		"success": "成功",
	})

	code, msg := langx.GetTransFormat("zh-CN", "success", map[string]string{})
	fmt.Println(code, msg)

}

// 导入语言包 基于Embed
func regByEmbed() {
	err := langx.RegisterEmbed(assetsFs)
	fmt.Println(err)

	code, msg := langx.GetTransFormat("zh", "success", map[string]string{})
	fmt.Println(code, msg)
	code, msg = langx.GetTransFormat("en", "error", map[string]string{
		"msg": "这是失败的原因",
	})
	fmt.Println(code, msg)
}

// 导入语言包 基于文件
func regByDir() {
	langx.RegisterDir("./lang")

	code, msg := langx.GetTransFormat("zh", "success", map[string]string{})
	fmt.Println(code, msg)
	code, msg = langx.GetTransFormat("en", "error", map[string]string{
		"msg": "这是失败的原因",
	})
	fmt.Println(code, msg)
}
