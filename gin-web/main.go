package main

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-22 10:12

import (
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	//userDao := dao.NewUserDAO()
	server.Run(":8081")
}
