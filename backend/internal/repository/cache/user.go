package cache

import (
	"backend/internal/domain"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-22 19:25

// 最佳实践 但是没有CacheV1的实现
//type CacheV1 interface {
//	Get(ctx context.Context, key string) (any, error)
//}
//
//type UserCache struct {
//	cache CacheV1
//}
//
//func NewUserCache() *UserCache {
//	return &UserCache{}
//}
//
//func (u *UserCache) GetUser(ctx context.Context, id int64) (domain.User, error) {
//	return domain.User{}, nil
//}

var ErrKeyNotExist = errors.New("Key 不存在")

type UserCache struct {
	// 传单击Redis可以 传Redis集群也可以
	client redis.Cmdable

	// 超时控制
	expiration time.Duration
}

func NewUserCache(client redis.Cmdable, expiration time.Duration) *UserCache {
	return &UserCache{
		client:     client,
		expiration: expiration,
	}
}

func (cache *UserCache) GetUser(ctx context.Context, id int64) (domain.User, error) {
	key := cache.key(id)
	bytes, err := cache.client.Get(ctx, key).Bytes()
	if err != nil {
		return domain.User{}, ErrKeyNotExist
	}
	var u domain.User
	err = json.Unmarshal(bytes, &u)
	return u, err
}

func (cache *UserCache) SetUser(ctx context.Context, u domain.User) error {
	val, err := json.Marshal(u)
	if err != nil {
		return err
	}
	return cache.client.Set(ctx, cache.key(u.Id), val, cache.expiration).Err()
}

func (cache *UserCache) key(id int64) string {
	return fmt.Sprintf("user:info:%d", id)
}
