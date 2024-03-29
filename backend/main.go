//@Author: wulinlin
//@Description:
//@File:  main
//@Version: 1.0.0
//@Date: 2024/02/20 21:13

package main

import (
	"backend/internal/repository"
	cache2 "backend/internal/repository/cache"
	dao2 "backend/internal/repository/dao"
	"backend/internal/service"
	"backend/internal/service/captcha"
	"backend/internal/service/sms/tencent"
	"backend/internal/web"
	"backend/internal/web/middleware"
	"backend/pkg/ginx/middleware/ratelimit"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
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
	// 初始化一个redis
	redisClient := redis.NewClient(&redis.Options{
		Addr: "192.168.1.52:6379",
	})
	server.Use(ratelimit.NewBuilder(redisClient, time.Minute, 100).Build())
	server.Use(cors.New(cors.Config{
		//AllowOrigins:     []string{"https://foo.com"},
		//AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
		//ExposeHeaders:    []string{"Content-Length "},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				return true
			}
			return strings.Contains(origin, "yourcompany.com")
		},
		MaxAge: 12 * time.Hour,
	}))

	//store := cookie.NewStore([]byte("secret"))
	client := redis.NewClient(&redis.Options{
		Addr:    "192.16.1.52:6379",
		Network: "tcp"})
	// 处理session问题 使用gin session插件
	newStore := memstore.NewStore([]byte("wulinlin"), []byte("wulinlin"))
	//if err2 != nil {
	//	panic(err2)
	//}
	store := newStore
	server.Use(sessions.Sessions("mysession", store))

	// 登录拦截
	server.Use(middleware.NewLoginMiddlewareBuilder().
		IgnorePaths("/users/signup").
		IgnorePaths("/users/login").Build())

	db, err := gorm.Open(mysql.Open("admin:123456@tcp(192.168.1.52:3306)/xhs"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// 初始化表结构创建
	err = dao2.InitTables(db)
	if err != nil {
		panic(err)
	}
	// 创建一大堆变量
	dao := dao2.NewUserDAO(db)
	cache := cache2.NewUserCache(client, time.Duration(15))
	repo := repository.NewUserRepository(dao, cache)
	svc := service.NewUserService(repo)
	codeCache := cache2.NewCodeCache(client)
	codeRepo := repository.NewCodeRepository(codeCache)
	smsClient := tencent.NewService()
	codeSvc := captcha.NewCodeService(codeRepo, smsClient)
	u := web.NewUserHandler(svc, codeSvc)
	server = web.RegisterRouters(server, u)
	server.Run(":8080")
}
