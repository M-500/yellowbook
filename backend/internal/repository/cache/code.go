package cache

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-23 14:25

var (
	ErrSetSendTooMany    = errors.New("验证码发送太频繁")
	ErrCodeVerifyTooMany = errors.New("验证次数太多")
	ErrUnknownForCode    = errors.New("未知错误")
)

//go:embed lua/set_code.lua
var luaSetCode string

//go:embed lua/set_code.lua
var luaVerifyCode string

type CodeCache struct {
	client redis.Cmdable
}

func NewCodeCache(c redis.Cmdable) *CodeCache {
	return &CodeCache{
		client: c,
	}
}

func (c *CodeCache) Set(ctx context.Context, biz, phone, code string) error {
	i, err := c.client.Eval(ctx, luaSetCode, []string{c.key(biz, phone), code}).Int()
	if err != nil {
		return err
	}
	switch i {
	case 0:
		// 毫无问题
		return nil
	case 1:
		// 发送太频繁
		return ErrSetSendTooMany
	default:
		// 系统错误
		return errors.New("系统内部错误")
	}
	return nil
}

func (c *CodeCache) Verify(ctx context.Context, biz, phone, inputCode string) (bool, error) {
	val, err := c.client.Eval(ctx, luaVerifyCode, []string{c.key(biz, phone), inputCode}).Int()
	if err != nil {
		return false, err
	}
	switch val {
	case 0:
		return true, nil
	case -1:
		//有人搞你
		return false, nil
	case -2:
		// 验证码出错
		return false, ErrCodeVerifyTooMany
	default:
		return false, ErrUnknownForCode
	}
}

func (c *CodeCache) key(biz, phone string) string {
	return fmt.Sprintf("phone_code:%s:%s", biz, phone)
}
