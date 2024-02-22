package cache

import (
	"backend/internal/domain"
	"context"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-22 19:25

type CacheV1 interface {
	Get(ctx context.Context, key string) (any, error)
}

type UserCache struct {
	cache CacheV1
}

func NewUserCache() *UserCache {
	return &UserCache{}
}

func (u *UserCache) GetUser(ctx context.Context, id int64) (domain.User, error) {
	return domain.User{}, nil
}
