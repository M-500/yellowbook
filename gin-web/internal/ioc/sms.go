package ioc

import (
	"gin-web/internal/service/sms"
	"gin-web/internal/service/sms/localsms"
	"github.com/redis/go-redis/v9"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-27 16:49

func InitSmsService(cmd redis.Cmdable) sms.ISMSService {
	// 这里选择用内存还是用 别的
	return localsms.NewService()
}
