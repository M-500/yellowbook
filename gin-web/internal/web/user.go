package web

import (
	"fmt"
	"gin-web/internal/domain"
	"gin-web/internal/service"
	"github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-26 15:49
const biz = "login"

type UserHandler struct {
	userSvc      service.IUserService
	codeSvc      service.ICodeService
	emailCompile *regexp2.Regexp
	pwdCompile   *regexp2.Regexp
	phoneCompile *regexp2.Regexp
}

// 邮箱校验验证码
const (
	emailRegexPattern = "^[a-zA-Z0-9_.-]+@[a-zA-Z0-9-]+(\\.[a-zA-Z0-9-]+)*\\.[a-zA-Z0-9]{2,6}$"         // 邮箱
	pwdRegexPattern   = "^(?![a-zA-Z]+$)(?!\\d+$)(?![^\\da-zA-Z\\s]+$).{6,32}$"                         // 由字母、数字、特殊字符，任意2种组成，6-32位
	phoneRegexPattern = "^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\\d{8}$" // 由字母、数字、特殊字符，任意2种组成，6-32位
	userIdKey         = "userId"
	bizLogin          = "login"
)

func NewUserHandler(svc service.IUserService, codeSvc service.ICodeService) *UserHandler {
	emailCompile := regexp2.MustCompile(emailRegexPattern, regexp2.Debug) // 预编译正则表达式
	pwdCompile := regexp2.MustCompile(pwdRegexPattern, regexp2.Debug)     // 预编译正则
	phoneCompile := regexp2.MustCompile(phoneRegexPattern, regexp2.Debug) // 预编译正则
	return &UserHandler{
		userSvc:      svc,
		codeSvc:      codeSvc,
		emailCompile: emailCompile,
		phoneCompile: phoneCompile,
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
		UserName        string `json:"username"`
		ConfirmPassword string `json:"rpassword"`
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
	//ok, err := h.emailCompile.MatchString(signForm.UserName)
	//if err != nil {
	//	// 正则匹配失败会返回Error
	//	c.JSON(http.StatusOK, gin.H{"msg": "系统内部错误！"})
	//	return
	//}
	//if !ok {
	//	c.JSON(http.StatusOK, gin.H{"msg": "邮箱格式不合法！"})
	//	return
	//}

	////校验密码是否合法
	//ok, err := h.pwdCompile.MatchString(signForm.Password)
	//if err != nil {
	//	// 正则匹配失败会返回Error
	//	c.JSON(http.StatusOK, gin.H{"msg": "系统内部错误！"})
	//	return
	//}
	//if !ok {
	//	c.JSON(http.StatusOK, gin.H{"msg": "密码必须包含字母，数字或者特殊字符的任意两种，且长度不能低于6位！"})
	//	return
	//}
	err := h.userSvc.Signup(c, domain.DMUser{
		Username: signForm.UserName,
		Pwd:      signForm.Password,
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
	type PwdLoginReq struct {
		UserName    string `json:"username"`
		Password    string `json:"password"`
		CaptchaId   string `json:"captcha_id"`
		CaptchaCode string `json:"captcha_code"`
	}
	var pwdForm PwdLoginReq
	if err := c.Bind(&pwdForm); err != nil {
		c.String(http.StatusUnauthorized, "数据不合法")
		return
	}
	//domainU, err := h.userSvc.Login(c, pwdForm.UserName, pwdForm.Password)
	_, err := h.userSvc.Login(c, pwdForm.UserName, pwdForm.Password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": err.Error()})
		return
	}
	// 登录成功后 生成一个JWT

	//claims := domain.UserClaims{
	//	RegisteredClaims: jwt.RegisteredClaims{},
	//	UserId:           int64(domainU.Id),
	//}
	//token := jwt.NewWithClaims(jwt.SigningMethodES512, claims)
	//tokenStr, err := token.SignedString([]byte("wulinlin"))
	//if err != nil {
	//	c.String(http.StatusInternalServerError, "系统错误")
	//	return
	//}
	//c.Header("x-jwt-token", tokenStr)
	//fmt.Println(tokenStr, "哈哈")
	c.JSON(http.StatusOK, Result{
		Data: map[string]string{
			"token": "1231",
		},
	})
	return
}

// PhoneLogin
//
//	@Description: 手机验证码登录
//	@receiver h
//	@param c
func (h *UserHandler) PhoneLogin(ctx *gin.Context) {
	type Req struct {
		Phone string `json:"phone"`
		Code  string `json:"code"`
	}
	var req Req
	if err := ctx.Bind(&req); err != nil {
		return
	}
	// 要不要查询数据库是否存在该手机号
	// 1. 校验手机号
	verify, err := h.codeSvc.Verify(ctx, biz, req.Phone, req.Code)
	if err != nil {
		ctx.JSON(http.StatusOK, Result{
			Code: 5,
			Msg:  "系统错误",
		})
		return
	}
	if !verify {
		// 校验不通过
		ctx.JSON(http.StatusOK, Result{
			Code: 4,
			Msg:  "验证码有误",
		})
		return
	}
	// 我这个手机号，会不会是一个新用户呢？
	// 这样子
	user, err := h.userSvc.FindOrCreate(ctx, req.Phone)
	if err != nil {
		ctx.JSON(http.StatusOK, Result{
			Code: 5,
			Msg:  "系统错误",
		})
		return
	}
	fmt.Println(user)
	// 3. 组装JWT

	ctx.JSON(http.StatusOK, Result{
		Msg: "验证码校验通过",
	})
}

func (h *UserHandler) SMSSender(c *gin.Context) {
	type SmsSendForm struct {
		Phone string `json:"phone"`
	}
	var smsForm SmsSendForm
	if err := c.Bind(&smsForm); err != nil {
		c.String(http.StatusUnauthorized, "数据不合法")
		return
	}
	// 正则匹配手机号是否合法
	ok, err := h.phoneCompile.MatchString(smsForm.Phone)
	if err != nil {
		// 正则匹配失败会返回Error
		c.JSON(http.StatusOK, gin.H{"msg": "系统内部错误"})
		return
	}
	if !ok {
		// 正则校验不通过，直接返回
		c.JSON(http.StatusOK, gin.H{"msg": "手机号码不合法"})
		return
	}
	err = h.codeSvc.Send(c, bizLogin, smsForm.Phone)
	switch err {
	case nil:
		c.JSON(http.StatusOK, Result{Msg: "发送成功"})
	case service.ErrCodeSendTooMany:
		c.JSON(http.StatusOK, Result{Msg: "短信发送太频繁，请稍后再试"})
	default:
		c.JSON(http.StatusOK, Result{Code: 5, Msg: "系统错误"})
		// 要打印日志
		return
	}
}
