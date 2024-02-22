package middleware

import (
	"backend/internal/web"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"strings"
	"time"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-21 18:55

type LoginJWTMiddlewareBuilder struct {
	paths []string
}

func NewLoginJWTMiddlewareBuilder() *LoginJWTMiddlewareBuilder {
	return &LoginJWTMiddlewareBuilder{
		paths: []string{},
	}
}
func (l *LoginJWTMiddlewareBuilder) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		for _, p := range l.paths {
			if ctx.Request.URL.Path == p {
				return
			}
		}
		// 用JWT校验

		// 1. 从Header中获取JWT
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// 切割JWT
		segs := strings.SplitN(token, " ", 2)
		if len(segs) != 2 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenStr := segs[1]
		claims := &web.UserClaims{}
		parse, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("wulinlin"), nil
		})
		// 非要校验是否过期 (其实不用校验)
		if claims.ExpiresAt.Time.Before(time.Now()) {
			// token过期了
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if err != nil {
			// 用户没有登录
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if !parse.Valid || parse == nil || claims.UserId == 0 {
			// 用户没有登录
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// v2 Start 每10s 刷新依稀
		now := time.Now()
		if claims.ExpiresAt.Sub(now) < time.Second*50 {
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute * 10))
			tokenStr, err = parse.SignedString([]byte("wulinlin")) // 续约token
			if err != nil {
				// 记录续约失败的日志
				log.Printf("JWT续约失败,%s", err)
			}
			ctx.Header("x-jwt-token", tokenStr) // 续约的token 再次返回给前端
		}
		// V2 end

		// V1 Start
		claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * 1))
		tokenStr, err = parse.SignedString([]byte("wulinlin")) // 续约token
		if err != nil {
			// 记录续约失败的日志
			log.Printf("JWT续约失败,%s", err)
		}
		ctx.Header("x-jwt-token", tokenStr) // 续约的token 再次返回给前端
		// V1 end

		// 将JWT解析的数据 插入到gin的Context上下文中
		//ctx.Set("claims", claims)
		ctx.Set("userId", claims.UserId)
	}
}
