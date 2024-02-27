package web

import "github.com/gin-gonic/gin"

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-22 18:15

func RegisterRouters(engine *gin.Engine, u *UserHandler) *gin.Engine {
	//u := &UserHandler{}  这个已经不能用啦              // 改成这个
	engine.POST("/users/signup", u.SignUp)      // 用户注册
	engine.POST("/users/login", u.PwdLogin)     // 用户登录
	engine.POST("/users/login-jwt", u.PwdLogin) // 用户登录
	//engine.GET("/users/profile", u.Profile)                 // 上传头像
	//engine.POST("/users/edit", u.Edit)                      // 编辑信息
	engine.POST("/login-sms/code/send", u.SMSSender) // 编辑信息
	//engine.POST("/login-sms", u.LoginSMS)                   // 编辑信息

	return engine
}
