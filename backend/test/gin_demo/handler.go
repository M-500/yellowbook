package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-26 15:18

type UserHandler struct {
	svc *UserService
}

func NewUserHandler(svc *UserService) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}
func (u *UserHandler) SignUp(ctx *gin.Context) {
	type SignForm struct {
		Username string `json:"username"`
		Pwd      string `json:"pwd"`
	}
	var user SignForm
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusPaymentRequired, gin.H{
			"msg": "数据不合法",
		})
		return
	}
	domainU := User{
		Username: user.Username,
		Pwd:      user.Pwd,
	}
	name := u.svc.FindByUserName(user.Username)
	if name == nil {
		_ = u.svc.CreateUser(domainU)
	}
	ctx.JSON(http.StatusPaymentRequired, gin.H{
		"msg": "用户注册成功",
	})
	return
}
func (u *UserHandler) Login(ctx *gin.Context) {
	type LoginForm struct {
		Username string `json:"username"`
		Pwd      string `json:"pwd"`
	}
	var user LoginForm
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusPaymentRequired, gin.H{
			"msg": "数据不合法",
		})
		return
	}
	w := u.svc.FindByUserName(user.Username)
	if w == nil {
		ctx.JSON(http.StatusPaymentRequired, gin.H{
			"msg": "用户不存在",
		})
		return
	}
	fmt.Println(w)

}
