package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
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
		// 我怎么知道1分钟已经过了？ 或者说我怎么定时刷新我的session？

		// 获取到上一次的更新时间
		updateTime := sess.Get("update_time")
		now := time.Now().UnixMilli()
		if updateTime == nil {
			// 还没有刷新过 那就刷新一下
			sess.Set("update_time", now)
			sess.Save()
			return
		}
		// 刷新时间过期？
		updateTimeVal, ok := updateTime.(int64)
		if !ok {

		}
		if now-updateTimeVal > 60*1000 {
			sess.Set("update_time", now)
			sess.Save()
			return
		}
	}
}

var IgnorePath []string

// CheckLoginV1
//
//	@Description: 中间件的V1版本
//	@return gin.HandlerFunc
func CheckLoginV1() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		for _, p := range IgnorePath {
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

// CheckLoginV2
//
//	@Description: session中间件的第二种方法
//	@return gin.HandlerFunc
func CheckLoginV2(paths []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		for _, p := range paths {
			if ctx.Request.URL.Path == p {
				return
			}
		}
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
