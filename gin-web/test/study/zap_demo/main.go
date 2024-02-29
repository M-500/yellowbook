package main

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-28 20:32

import (
	"errors"
	"go.uber.org/zap"
)

type HaHaDemo struct {
	Name string
	Age  int32
}

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	zap.L().Info("这一条是不会显示的，因为没有替换")
	zap.ReplaceGlobals(logger) // 替换掉你全局的zap，之后在代码里就可以使用 zap.xx来使用日志了
	zap.L().Info("替换过后，所以你就可以使用啦")

	zap.L().Info("测试",
		zap.Error(errors.New("错误类型，error")),
		zap.Int64("int64类型的错误", 123),
		zap.Any("任意类型的错误", HaHaDemo{
			Name: "狗东西",
			Age:  100,
		}),
	)
}
