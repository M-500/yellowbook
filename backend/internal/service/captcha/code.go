package captcha

import (
	"backend/internal/repository"
	"backend/internal/service/sms"
	"context"
	"fmt"
	"math/rand"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-23 11:03

type CodeService struct {
	repo  *repository.CodeRepository
	sms   sms.Service
	tplId string
}

func NewCodeService(r *repository.CodeRepository, smsClient sms.Service) *CodeService {
	return &CodeService{
		repo: r,
		sms:  smsClient,
	}
}

// Send
//
//	思考一个问题，这个验证码的生成是归调用者生成，还是在这里生成？
//	@Description: 发送验证码
//	@receiver c
//	@param ctx
//	@param code
//	@param phone
//	@param biz  区别哪一条业务线
func (svc *CodeService) Send(ctx context.Context, phone string, biz string) error {
	// 1. 生成验证码 2.塞进Redis 3.发送数据
	captcha := svc.generateCode()
	err := svc.repo.Store(ctx, biz, phone, captcha)
	if err != nil {
		return err
	}
	err = svc.sms.Send(ctx, svc.tplId, []string{captcha}, phone)
	if err != nil {
		// 这个地方怎么办？ 这意味着 Redis有验证码，但是调用三方短信验证码出错 你也不知道有没有真的发送出去，所以你不能删redis中的数据
		// 1. 重发，如果重发的话，需要svc内嵌一个重试 的实现过来，让他去重试 RetryAble  会涉及到错误扩展的问题
		// 2.
	}
	return err
}

func (svc *CodeService) Verify(ctx context.Context, code string, phone string, biz string) (bool, error) {
	return svc.repo.Verify(ctx, biz, phone, code)
}

func (svc *CodeService) generateCode() string {
	num := rand.Intn(1000000)
	// 不够6位的，加上前导0
	return fmt.Sprintf("%6d", num)
}
