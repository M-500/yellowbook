package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-21 18:55

type LoginMiddlewareBuilder struct {
	paths []string
}

func NewLoginMiddlewareBuilder() *LoginMiddlewareBuilder {
	return &LoginMiddlewareBuilder{
		paths: []string{},
	}
}
func (l *LoginMiddlewareBuilder) IgnorePaths(path string) *LoginMiddlewareBuilder {
	l.paths = append(l.paths, path)
	return l
}
func (l *LoginMiddlewareBuilder) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		for _, p := range l.paths {
			if ctx.Request.URL.Path == p {
				return
			}
		}
		//if ctx.Request.URL.Path == "/users/signup" || ctx.Request.URL.Path == "/users/login" {
		//	//ctx.Next()
		//	return
		//}
		sess := sessions.Default(ctx)
		if sess == nil {
			ctx.AbortWithStatus(http.StatusForbidden) // 不让访问
			return
		}
		id := sess.Get("userId")
		if id == nil {
			ctx.AbortWithStatus(http.StatusForbidden) // 不让访问
			return
		}
	}
}
