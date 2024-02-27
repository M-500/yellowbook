package ralimie

import (
	"context"
	"fmt"
	"gin-web/internal/service/sms"
	"gin-web/pkg/ginx/retelimiter"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-27 18:22

// 好处是只需要实现你想要装饰的方法，坏处是别人可以绕开你的装饰
type RatelimitSMSServiceV1 struct {
	sms.ISMSService // 组合原有接口
	limiter         retelimiter.Limiter
}

func NewServiceV1(svc sms.ISMSService, limiter retelimiter.Limiter) sms.ISMSService {
	return &RatelimitSMSService{
		svc:     svc,
		limiter: limiter,
	}
}

func (s *RatelimitSMSServiceV1) Send(ctx context.Context, tpl string, args []string, numbers ...string) error {
	limited, err := s.limiter.Limited(ctx, "sms:tencent")
	if err != nil {
		// 系统错误
		// 可以限流：保守策略，你的下游很坑的时候，
		// 可以不限：你的下游很强，业务可用性要求很高，尽量容错策略
		// 包一下这个错误
		return fmt.Errorf("短信服务判断是否限流出现问题，%w", err)
	}
	if limited {
		return errLimited
	}
	// 你这里加一些代码，新特性
	err = s.ISMSService.Send(ctx, tpl, args, numbers...)
	// 你在这里也可以加一些代码，新特性
	return err
}
