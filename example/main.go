package main

import (
	"embed"
	"fmt"

	"github.com/yuninks/langx"
)

//go:embed lang
var assetsFs embed.FS

func main() {
	demoEmbed()
}

func demoEmbed() {
	err := langx.RegisterEmbed(assetsFs)
	fmt.Println(err)

	code, msg := langx.GetTransFormat("zh", "success", map[string]string{}) 
	fmt.Println(code, msg)
	code, msg = langx.GetTransFormat("en", "error", map[string]string{
		"msg": "这是失败的原因",
	})
	fmt.Println(code, msg)
}

func demo1() {
	langx.RegisterDir("./lang")

	code, msg := langx.GetTransFormat("zh", "success", map[string]string{})
	fmt.Println(code, msg)
	code, msg = langx.GetTransFormat("en", "error", map[string]string{
		"msg": "这是失败的原因",
	})
	fmt.Println(code, msg)
}
