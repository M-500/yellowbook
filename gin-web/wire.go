//go:build wireinject

package main

import (
	"gin-web/internal/ioc"
	"gin-web/internal/repository"
	"gin-web/internal/repository/cache"
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
		ioc.InitDB,
		ioc.InitRedis,

		cache.NewCodeCache,
		cache.NewRedisCompanyCache,
		dao.NewUserDAO,
		dao.NewEnterpriseDao,
		dao.NewParseLogDao,

		repository.NewUserRepository,
		repository.NewCodeRepository,
		repository.NewCompanyRepo,

		service.NewCodeService,
		service.NewUserService,
		service.NewCaptchaService,
		service.NewExcelParserService,

		web.NewUserHandler,
		web.NewCaptchaHandler,
		web.NewExcelHandler,

		// 你中间件呢？注册路由呢
		ioc.InitGin,
		ioc.InitMiddlewares,
		ioc.InitSmsService,
	)

	return new(gin.Engine)
}
