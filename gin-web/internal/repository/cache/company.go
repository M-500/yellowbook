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
	Exist(ctx context.Context, key string) (bool, error)
}

type RedisCompanyCache struct {
	redis redis.Cmdable
}

func (r *RedisCompanyCache) Add(ctx context.Context, key string, val any) error {
	//TODO implement me
	panic("implement me")
}

func (r *RedisCompanyCache) Exist(ctx context.Context, key string) (bool, error) {

	//TODO implement me
	panic("implement me")
}
