package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"gin-web/internal/domain"
	"github.com/redis/go-redis/v9"
	"time"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-22 19:25
type UserCacheInterface interface {
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (domain.DMUser, error)
	Set(ctx context.Context, u domain.DMUser) error
}

type RedisCacheImpl struct {
	// redis客户端
	cmd redis.Cmdable
	// 过期时间
	expiration time.Duration
}

func NewRedisCacheImpl(cmd redis.Cmdable, expire time.Duration) *RedisCacheImpl {
	return &RedisCacheImpl{
		cmd:        cmd,
		expiration: expire,
	}
}

func (r *RedisCacheImpl) key(id int64) string {
	return fmt.Sprintf("user:info:%s", id)
}

func (r *RedisCacheImpl) Delete(ctx context.Context, id int64) error {
	return r.cmd.Del(ctx, r.key(id)).Err()
}

func (r *RedisCacheImpl) Get(ctx context.Context, id int64) (domain.DMUser, error) {
	key := r.key(id)
	data, err := r.cmd.Get(ctx, key).Result()
	if err != nil {
		return domain.DMUser{}, err
	}
	// 反序列化回来
	var u domain.DMUser
	err = json.Unmarshal([]byte(data), &u)
	return u, err
}

func (r *RedisCacheImpl) Set(ctx context.Context, u domain.DMUser) error {
	data, err := json.Marshal(u)
	if err != nil {
		return err
	}
	key := r.key(int64(u.Id))
	return r.cmd.Set(ctx, key, data, r.expiration).Err()
}
