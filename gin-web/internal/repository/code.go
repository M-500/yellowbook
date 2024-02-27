package repository

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-27 16:10

import (
	"context"
	"gin-web/internal/repository/cache"
)

var (
	ErrCodeSendTooMany        = cache.ErrCodeSendTooMany
	ErrCodeVerifyTooManyTimes = cache.ErrCodeVerifyTooManyTimes
)

type ICodeRepo interface {
	Store(ctx context.Context, biz string, phone string, code string) error
	Verify(ctx context.Context, biz, phone, inputCode string) (bool, error)
}

type CodeRepository struct {
	cache cache.ICodeCache
}

func NewCodeRepository(client cache.ICodeCache) ICodeRepo {
	return &CodeRepository{
		cache: client,
	}
}

func (c *CodeRepository) Store(ctx context.Context, biz string, phone string, code string) error {

	return c.cache.Set(ctx, biz, phone, code)
}

func (c *CodeRepository) Verify(ctx context.Context, biz string, phone string, inputCode string) (bool, error) {
	return c.cache.Verify(ctx, biz, phone, inputCode)
}
