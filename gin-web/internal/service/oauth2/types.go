package oauth2

import (
	"context"
	"gin-web/internal/domain"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-28 10:09

type IOAuth2 interface {
	AuthURL(ctx context.Context, state string) (string, error)
	VerifyCode(ctx context.Context, code string) (domain.WechatInfo, error)
}
