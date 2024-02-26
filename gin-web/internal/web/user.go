package web

import (
	"gin-web/internal/service"
	"github.com/gin-gonic/gin"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-26 15:49

type UserHandler struct {
	userSvc service.UserSvcInterface
}

func NewUserHandler(svc service.UserSvcInterface) *UserHandler {
	return &UserHandler{
		userSvc: svc,
	}
}

// SignUp
//
//	@Description: 注册接口
//	@receiver h
//	@param c
func (h *UserHandler) SignUp(c *gin.Context) {

}

// PwdLogin
//
//	@Description: 密码登录
//	@receiver h
//	@param c
func (h *UserHandler) PwdLogin(c *gin.Context) {

}

// PhoneLogin
//
//	@Description: 手机验证码登录
//	@receiver h
//	@param c
func (h *UserHandler) PhoneLogin(c *gin.Context) {

}
