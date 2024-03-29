package service

import (
	"context"
	"fmt"
	"gin-web/internal/repository"
	"gin-web/internal/service/sms"
	"go.uber.org/atomic"
	"math/rand"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-26 19:39

var codeTplId atomic.String = atomic.String{}
var (
	ErrCodeVerifyTooManyTimes = repository.ErrCodeVerifyTooManyTimes
	ErrCodeSendTooMany        = repository.ErrCodeSendTooMany
)

type ICodeService interface {
	Send(ctx context.Context,
		// 区别业务场景
		biz string, phone string) error
	Verify(ctx context.Context, biz string,
		phone string, inputCode string) (bool, error)
}
type codeService struct {
	sms  sms.ISMSService
	repo repository.ICodeRepo
}

func NewCodeService(sms sms.ISMSService, repo repository.ICodeRepo) ICodeService {
	return &codeService{
		sms:  sms,
		repo: repo,
	}
}

func (svc *codeService) Send(ctx context.Context,
	// 区别业务场景
	biz string, phone string) error {
	// 1. 生成一个验证码
	code := svc.generateCode()
	// 2. 将验证码存入到redis中
	err := svc.repo.Store(ctx, biz, phone, code)
	if err != nil {
		// 有问题
		return err
	}

	// 3. 将验证码发出去
	err = svc.sms.Send(ctx, "", []string{code}, phone)
	if err != nil {
		err = fmt.Errorf("发送短信出现异常 %w", err) // 将错误包一层
	}
	//if err != nil {
	// 这个地方怎么办？
	// 这意味着，Redis 有这个验证码，但是不好意思，
	// 我能不能删掉这个验证码？
	// 你这个 err 可能是超时的 err，你都不知道，发出了没
	// 在这里重试
	// 要重试的话，初始化的时候，传入一个自己就会重试的 smsSvc
	//}
	return err
}

func (svc *codeService) Verify(ctx context.Context, biz string, phone string, inputCode string) (bool, error) {
	return svc.repo.Verify(ctx, biz, phone, inputCode)

}

func (svc *codeService) generateCode() string {
	// 六位数，num 在 0, 999999 之间，包含 0 和 999999
	num := rand.Intn(1000000)
	// 不够六位的，加上前导 0
	// 000001
	return fmt.Sprintf("%06d", num)
}
