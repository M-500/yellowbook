package web

import (
	"gin-web/internal/domain"
	"gin-web/internal/service"
	"github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-26 15:49

type UserHandler struct {
	userSvc      service.UserSvcInterface
	emailCompile *regexp2.Regexp
	pwdCompile   *regexp2.Regexp
}

// 邮箱校验验证码
const (
	emailRegexPattern = "^[a-zA-Z0-9_.-]+@[a-zA-Z0-9-]+(\\.[a-zA-Z0-9-]+)*\\.[a-zA-Z0-9]{2,6}$" // 邮箱
	pwdRegexPattern   = "^(?![a-zA-Z]+$)(?!\\d+$)(?![^\\da-zA-Z\\s]+$).{6,32}$"                 // 由字母、数字、特殊字符，任意2种组成，6-32位
)

func NewUserHandler(svc service.UserSvcInterface) *UserHandler {
	emailCompile := regexp2.MustCompile(emailRegexPattern, regexp2.Debug) // 预编译正则表达式
	pwdCompile := regexp2.MustCompile(pwdRegexPattern, regexp2.Debug)     // 预编译正则
	return &UserHandler{
		userSvc:      svc,
		emailCompile: emailCompile,
		pwdCompile:   pwdCompile,
	}
}

// SignUp
//
//	@Description: 注册接口
//	@receiver h
//	@param c
func (h *UserHandler) SignUp(c *gin.Context) {
	type SignUpForm struct {
		Email           string `json:"email"`
		ConfirmPassword string `json:"confirmPassword"`
		Password        string `json:"password"`
	}
	var signForm SignUpForm
	if err := c.Bind(&signForm); err != nil {
		c.String(http.StatusUnauthorized, "数据不合法")
		return
	}
	// 校验数据类型
	// 校验两个密码是否相同
	if signForm.ConfirmPassword != signForm.Password {
		c.JSON(http.StatusOK, gin.H{"msg": "两次密码输入不一致！"})
		return
	}
	ok, err := h.emailCompile.MatchString(signForm.Email)
	if err != nil {
		// 正则匹配失败会返回Error
		c.JSON(http.StatusOK, gin.H{"msg": "系统内部错误！"})
		return
	}
	if !ok {
		c.JSON(http.StatusOK, gin.H{"msg": "邮箱格式不合法！"})
		return
	}

	//校验密码是否合法
	ok, err = h.pwdCompile.MatchString(signForm.Password)
	if err != nil {
		// 正则匹配失败会返回Error
		c.JSON(http.StatusOK, gin.H{"msg": "系统内部错误！"})
		return
	}
	if !ok {
		c.JSON(http.StatusOK, gin.H{"msg": "密码必须包含字母，数字或者特殊字符的任意两种，且长度不能低于6位！"})
		return
	}
	err = h.userSvc.Signup(c, domain.DMUser{
		Email: signForm.Email,
		Pwd:   signForm.Password,
	})
	if err != nil {

		return
	}
	//响应前端 注册用户成功
	c.JSON(http.StatusOK, gin.H{"msg": "注册成功！"})
	return
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