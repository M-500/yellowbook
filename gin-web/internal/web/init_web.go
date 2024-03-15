package web

import "github.com/gin-gonic/gin"

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-22 18:15

func RegisterRouters(engine *gin.Engine, u *UserHandler, c *CaptchaHandler) *gin.Engine {
	//u := &UserHandler{}  这个已经不能用啦              // 改成这个
	engine.POST("/api/v1/users/signup", u.SignUp)      // 用户注册
	engine.POST("/api/v1/users/pwd-login", u.PwdLogin) // 用户登录
	engine.POST("/api/v1/users/login-jwt", u.PwdLogin) // 用户登录
	//engine.GET("/users/profile", u.Profile)                 // 上传头像
	//engine.POST("/users/edit", u.Edit)                      // 编辑信息
	engine.POST("/api/v1/login-sms/code/send", u.SMSSender) // 编辑信息
	engine.POST("/api/v1/login-sms", u.PhoneLogin)          // 编辑信息
	engine.GET("/api/v1/captcha", c.CaptchaImgAPI)          // 编辑信息

	return engine
}
