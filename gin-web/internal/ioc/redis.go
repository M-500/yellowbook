package ioc

import "github.com/redis/go-redis/v9"

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-27 13:47

func InitRedis() redis.Cmdable {
	// 初始化redis连接
	redisClient := redis.NewClient(&redis.Options{
		Addr: "192.168.1.52:6379",
		DB:   8,
	})
	return redisClient
}
