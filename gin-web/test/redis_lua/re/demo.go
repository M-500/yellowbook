package main

import (
	"fmt"
	"github.com/dlclark/regexp2"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-27 17:42

const phoneRegexPattern = "^1[345789]{1}\\d{9}$"

func main() {
	phoneCompile := regexp2.MustCompile(phoneRegexPattern, regexp2.Debug)
	ok, err := phoneCompile.MatchString("18574945221")
	if err != nil {
		fmt.Println("内部错误")
	}
	fmt.Println(ok)
}
