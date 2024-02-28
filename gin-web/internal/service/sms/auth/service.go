package auth

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-27 21:01
import (
	"context"
	"errors"
	"gin-web/internal/service/sms"
	"github.com/golang-jwt/jwt/v5"
)

type AuthSmsService struct {
	svc sms.ISMSService
	key string
}

func NewAuthSmsService(svc sms.ISMSService) sms.ISMSService {
	return &AuthSmsService{
		svc: svc,
	}
}

func (s *AuthSmsService) Send(ctx context.Context, tpl string, args []string, numbers ...string) error {
	// 这里做权限校验

	var tc TokenClaims
	// 如果我这里能解析成功，就说明是对应的业务方
	token, err := jwt.ParseWithClaims(tpl, &tc, func(tk *jwt.Token) (interface{}, error) {
		return s.key, nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return errors.New("token非法")
	}
	return s.svc.Send(ctx, tpl, args, numbers...)
}

type TokenClaims struct {
	jwt.RegisteredClaims
	TemplateId string
}
