package wechat

import (
	"context"
	"gin-web/internal/domain"
	"gin-web/internal/service/oauth2"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-28 10:08

type WechatService struct {
}

func NewWechatService() oauth2.IOAuth2 {
	return &WechatService{}
}

func (w WechatService) AuthURL(ctx context.Context, state string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (w WechatService) VerifyCode(ctx context.Context, code string) (domain.WechatInfo, error) {
	//TODO implement me
	panic("implement me")
}
