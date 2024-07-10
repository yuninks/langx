package main

import (
	"fmt"

	"github.com/yuninks/langx"
)

func main() {
	langx.RegisterDir("./lang")

	code, msg := langx.GetTransFormat("zh", "success", map[string]string{})
	fmt.Println(code, msg)
	code, msg = langx.GetTransFormat("en", "error", map[string]string{
		"msg": "这是失败的原因",
	})
	fmt.Println(code, msg)
}
