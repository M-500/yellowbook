//go:build wireinject

package ioc

import (
	"gin-web/internal/repository"
	"gin-web/internal/repository/dao"
	"gin-web/internal/service"
	"gin-web/internal/web"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-27 13:43

func InitWebServer() *gin.Engine {
	wire.Build(
		InitDB,
		InitRedis,

		dao.NewUserDAO,

		repository.NewUserRepository,

		service.NewUserService,

		web.NewUserHandler,

		// 你中间件呢？注册路由呢
		InitGin,
		InitMiddlewares,
	)

	return new(gin.Engine)
}
