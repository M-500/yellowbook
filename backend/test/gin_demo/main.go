package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-26 15:17

func main() {
	server := gin.Default()
	db, err := gorm.Open(mysql.Open("admin:123456@tcp(192.168.1.52:3306)/test?charset=utf8&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})
	dao := NewUserDAO(db)
	svc := NewUserService(dao)
	handler := NewUserHandler(svc)

	server.POST("/user/pwd-login", handler.Login)
	server.POST("/user/sign", handler.SignUp)

	server.Run(":8001")
}
