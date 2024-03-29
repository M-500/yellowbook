package repository

import (
	"backend/internal/repository/cache"
	"context"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-23 14:57
var (
	ErrSetSendTooManyRepo    = cache.ErrSetSendTooMany
	ErrCodeVerifyTooManyRepo = cache.ErrCodeVerifyTooMany
)

type CodeRepository struct {
	cache *cache.CodeCache
}

func NewCodeRepository(c *cache.CodeCache) *CodeRepository {
	return &CodeRepository{
		cache: c,
	}
}

func (repo *CodeRepository) Store(ctx context.Context, biz, phone, code string) error {
	return repo.cache.Set(ctx, biz, phone, code)
}

func (repo *CodeRepository) Verify(ctx context.Context, biz, phone, inputCode string) (bool, error) {
	return repo.cache.Verify(ctx, biz, phone, inputCode)
}
