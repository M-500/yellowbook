package ioc

import (
	"gin-web/internal/web"
	"gin-web/internal/web/middleware"
	"gin-web/pkg/ginx/middleware/metric"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-27 14:57

func InitGin(hdl *web.UserHandler, c *web.CaptchaHandler, e *web.ExcelHandler, mdls []gin.HandlerFunc) *gin.Engine {
	server := gin.Default()
	server.Use(mdls...)
	web.RegisterRouters(server, hdl, c, e)
	return server
}

// InitMiddlewares
//
//	@Description: 初始化中间件的注册
//	@param redisClient
//	@return []gin.HandlerFunc
func InitMiddlewares(redisClient redis.Cmdable) []gin.HandlerFunc {
	return []gin.HandlerFunc{
		//ratelimit.NewBuilder(redisClient, time.Minute, 3).Build(), // 限流组件
		middleware.CorsMiddleware(), // 跨域中间件
		(&metric.MiddlewareBuilder{
			NameSpace:  "test_gin_web",
			Name:       "gin_http",
			SubSystem:  "damn1",
			Help:       "统计gin_web的HTTP接口",
			InstanceId: "my_instance_id",
		}).Build(),
		otelgin.Middleware("webook"),
		//middleware.NewLoginMiddlewareBuilder().
		//	IgnorePath("/user/login").
		//	IgnorePath("/login-sms/code/send").
		//	IgnorePath("/user/signup").Build(), // Jwt中间件
	}
}
