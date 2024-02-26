//@Author: wulinlin
//@Description: redis 滑动窗口限流器的实现
//@File:  redis_slid_window
//@Version: 1.0.0
//@Date: 2024/02/26 21:41

package retelimiter

import (
	"context"
	_ "embed"
	"github.com/redis/go-redis/v9"
	"time"
)

// go:embed slide_window.lua
var luaSlideWindow string

type RedisSlidingWindowLimiter struct {
	cmd redis.Cmdable
	// 窗口大小
	interval time.Duration
	// 阈值
	rate int
	// interval 内 允许 rate个请求
}

func NewRedisSlidingWindowLimiter(client redis.Cmdable, interval time.Duration, rate int) Limiter {
	return &RedisSlidingWindowLimiter{
		cmd:      client,
		interval: interval,
		rate:     rate,
	}
}

func (r *RedisSlidingWindowLimiter) Limited(ctx context.Context, key string) (bool, error) {
	return r.cmd.Eval(ctx, luaSlideWindow, []string{key}, r.interval.Milliseconds(), r.rate, time.Now().UnixMilli()).Bool()
}
