package main

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-22 10:12

import (
	"gin-web/internal/ioc"
)

func main() {
	server := ioc.InitWebServer()
	server.Run(":8081")
}
