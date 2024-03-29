//@Author: wulinlin
//@Description:
//@File:  user
//@Version: 1.0.0
//@Date: 2024/02/20 21:15

package web

import (
	"backend/internal/domain"
	"backend/internal/repository"
	"backend/internal/service"
	"backend/internal/service/captcha"
	"fmt"
	"github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

const (
	emailRegexPattern = "^[a-zA-Z0-9_.-]+@[a-zA-Z0-9-]+(\\.[a-zA-Z0-9-]+)*\\.[a-zA-Z0-9]{2,6}$" // 邮箱
	pwdRegexPattern   = "^(?![a-zA-Z]+$)(?!\\d+$)(?![^\\da-zA-Z\\s]+$).{6,32}$"                 // 由字母、数字、特殊字符，任意2种组成，6-32位
)

type UserHandler struct {
	svc          *service.UserService
	emailCompile *regexp2.Regexp
	pwdCompile   *regexp2.Regexp
	codeSvc      *captcha.CodeService
}

func NewUserHandler(svc *service.UserService, codeSvc *captcha.CodeService) *UserHandler {
	emailCompile := regexp2.MustCompile(emailRegexPattern, regexp2.Debug) // 预编译正则表达式
	pwdCompile := regexp2.MustCompile(pwdRegexPattern, regexp2.Debug)     // 预编译正则
	return &UserHandler{
		svc:          svc,
		emailCompile: emailCompile,
		pwdCompile:   pwdCompile,
		codeSvc:      codeSvc,
	}
}
func (u *UserHandler) SignUp(ctx *gin.Context) {
	// 创建一个匿名结构体 用来映射前端的body
	type SignUpReq struct {
		Email           string `json:"email"`
		ConfirmPassword string `json:"confirmPassword"`
		Password        string `json:"password"`
	}
	var req SignUpReq
	// 如果前端传递的JSON信息和上面结构体不匹配，会报错的
	if err := ctx.Bind(&req); err != nil {
		return
	}

	// 校验数据类型
	ok, err := u.emailCompile.MatchString(req.Email)
	if err != nil {
		// 正则匹配失败会返回Error
		ctx.JSON(http.StatusOK, gin.H{"msg": "系统内部错误！"})
		return
	}
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{"msg": "邮箱格式不合法！"})
		return
	}
	// 校验两个密码是否相同
	if req.ConfirmPassword != req.Password {
		ctx.JSON(http.StatusOK, gin.H{"msg": "两次密码输入不一致！"})
		return
	}
	//校验密码是否合法
	ok, err = u.pwdCompile.MatchString(req.Password)
	if err != nil {
		// 正则匹配失败会返回Error
		ctx.JSON(http.StatusOK, gin.H{"msg": "系统内部错误！"})
		return
	}
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{"msg": "密码必须包含字母，数字或者特殊字符的任意两种，且长度不能低于6位！"})
		return
	}

	// 这里应该去操作数据库 完成注册的逻辑
	err = u.svc.SignUp(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if err == repository.RepoErrUserDuplicateEmail {
		ctx.JSON(http.StatusOK, gin.H{"msg": "邮箱已经存在"})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"msg": "系统内部错误"})
		return
	}
	//响应前端
	ctx.JSON(http.StatusOK, gin.H{"msg": "注册成功！"})
}

func (u *UserHandler) LoginJWT(ctx *gin.Context) {
	type LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req LoginReq
	if err := ctx.Bind(&req); err != nil {
		return
	}
	domainU, err := u.svc.Login(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"msg": "系统内部错误"})
		return
	}

	// 这里生成一个JWT
	//token := jwt.New(jwt.SigningMethodES512)
	//token := jwt.NewWithClaims(jwt.SigningMethodES512, jwt.MapClaims{"userId": domainU.Id}) // 不建议用map来传递
	claims := UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 6)), // 设置过期时间
		},
		UserId: domainU.Id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES512, claims)
	tokenStr, err := token.SignedString([]byte("wulinlin"))
	if err != nil {
		ctx.String(http.StatusInternalServerError, "系统错误")
		return
	}
	ctx.Header("x-jwt-token", tokenStr)
	fmt.Println(tokenStr, "哈哈")
	ctx.JSON(http.StatusOK, gin.H{"msg": "登录成功！"})
	return
}

func (u *UserHandler) Login(ctx *gin.Context) {
	type LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req LoginReq
	if err := ctx.Bind(&req); err != nil {
		return
	}
	doUser, err := u.svc.Login(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"msg": "系统内部错误"})
		return
	}
	//登录成功后 将session拿出来
	sess := sessions.Default(ctx)
	// 设置值到session中 比如设置用户id
	sess.Set("userId", doUser.Id)
	sess.Options(sessions.Options{
		//Path:     "",
		//Domain:   "",
		//MaxAge:   0,
		Secure: true,
		//HttpOnly: false,
		//SameSite: 0,
	})
	sess.Save()
	ctx.JSON(http.StatusOK, gin.H{"msg": "登录成功！"})
	return
}

func (u *UserHandler) Logout(ctx *gin.Context) {
	sess := sessions.Default(ctx)
	// 设置值到session中 比如设置用户id
	sess.Options(sessions.Options{
		MaxAge: -1,
	})
	sess.Save()
	ctx.JSON(http.StatusOK, gin.H{"msg": "退出登录！"})
}
func (u *UserHandler) ProfileJWT(ctx *gin.Context) {
	// 获取jwt的内容
	value, ok := ctx.Get("claims") // 这里拿到的是一个any类型
	if !ok {
		// 这部分的逻辑你可以省略，因为必然会有 claims
		// 当然也可以考虑监控住这里
		// 最眼睛的写法
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	claims, ok := value.(*UserClaims)
	if !ok {
		// ok代表是否成功
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	println(claims.UserId)
	// 写其他逻辑
}

func (u *UserHandler) Profile(ctx *gin.Context) {
	// 获取jwt的内容

}

func (u *UserHandler) Edit(ctx *gin.Context) {
}

func (u *UserHandler) SendSMSLoginCode(ctx *gin.Context) {
}

func (u *UserHandler) LoginSMS(ctx *gin.Context) {
}

type UserClaims struct {
	jwt.RegisteredClaims
	// 声明自己需要放入token里的数据
	UserId int64
}
