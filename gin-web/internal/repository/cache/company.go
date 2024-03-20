//@Author: wulinlin
//@Description:
//@File:  company
//@Version: 1.0.0
//@Date: 2024/03/19 23:56

package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type CompanyCache interface {
	Add(ctx context.Context, key string, val any) error
	Exist(ctx context.Context, key string, val any) (bool, error)
}

type RedisCompanyCache struct {
	redis redis.Cmdable
}

func NewRedisCompanyCache(redis redis.Cmdable) CompanyCache {
	return &RedisCompanyCache{
		redis: redis,
	}
}

func (r *RedisCompanyCache) Add(ctx context.Context, key string, val any) error {
	return r.redis.SAdd(ctx, key, val).Err()
}

func (r *RedisCompanyCache) Exist(ctx context.Context, key string, val any) (bool, error) {
	//result, err := r.redis.SIsMember(ctx, key,val).Result()
	//if err != nil {
	//	return false, err
	//}
	return r.redis.SIsMember(ctx, key, val).Result()
}
