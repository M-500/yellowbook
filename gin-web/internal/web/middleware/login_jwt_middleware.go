package middleware

import (
	"gin-web/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
	"time"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-27 12:01

type LoginMiddlewareBuilder struct {
	paths []string // 不需要JWT校验的URL
}

func NewLoginMiddlewareBuilder() *LoginMiddlewareBuilder {
	return &LoginMiddlewareBuilder{}
}

func (l *LoginMiddlewareBuilder) IgnorePath(path string) *LoginMiddlewareBuilder {
	l.paths = append(l.paths, path)
	return l
}

func (l *LoginMiddlewareBuilder) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 如果命中忽略 URL 直接放行
		for _, p := range l.paths {
			if ctx.Request.URL.Path == p {
				// 直接放行
				ctx.Next()
				return
			}
		}

		// 1. 从请求中获取token  TODO: 从其他任何位置获取token
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

		// 解析jwt
		claims := domain.UserClaims{}
		parse, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			// TODO 这里需要用配置文件来处理秘钥
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
	}
}
