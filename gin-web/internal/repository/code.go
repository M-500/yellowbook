package repository

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-27 16:10

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type ICodeRepo interface {
	Store(ctx context.Context, biz string, phone string, code string) error
	Verify(ctx context.Context, biz, phone, inputCode string) (bool, error)
}

type CodeRepository struct {
	client redis.Cmdable
}

func NewCodeRepository(client redis.Cmdable) ICodeRepo {
	return &CodeRepository{
		client: client,
	}
}

func (c CodeRepository) Store(ctx context.Context, biz string, phone string, code string) error {
	//TODO implement me
	panic("implement me")
}

func (c CodeRepository) Verify(ctx context.Context, biz string, phone string, inputCode string) (bool, error) {
	//TODO implement me
	panic("implement me")
}
