//@Author: wulinlin
//@Description:
//@File:  main
//@Version: 1.0.0
//@Date: 2024/02/20 21:13

package main

import (
	"backend/internal/repository"
	dao2 "backend/internal/repository/dao"
	"backend/internal/service"
	"backend/internal/web"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

func main() {
	server := gin.Default()
	server.Use(func(context *gin.Context) {
		fmt.Println("这里是第1个middleware")
	})
	server.Use(func(context *gin.Context) {
		fmt.Println("这里是第2个middleware")
	})

	server.Use(cors.New(cors.Config{
		//AllowOrigins:     []string{"https://foo.com"},
		//AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
		//ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				return true
			}
			return strings.Contains(origin, "yourcompany.com")
		},
		MaxAge: 12 * time.Hour,
	}))

	db, err := gorm.Open(mysql.Open("admin:123456@tcp(192.168.1.52:3306)/xhs"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// 创建一大堆变量
	dao := dao2.NewUserDAO(db)
	repo := repository.NewUserRepository(dao)
	svc := service.NewUserService(repo)
	u := web.NewUserHandler(svc)
	server = web.RegisterRouters(server, u)
	server.Run(":8080")
}
