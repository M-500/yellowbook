//@Author: wulinlin
//@Description:
//@File:  service
//@Version: 1.0.0
//@Date: 2024/02/26 22:38

package ralimie

import (
	"context"
	"fmt"
	"gin-web/internal/service/sms"
	"gin-web/pkg/ginx/retelimiter"
)

var errLimited = fmt.Errorf("触发了限流")

type RatelimitSMSService struct {
	svc     sms.ISMSService // 组合原有接口
	limiter retelimiter.Limiter
}

func NewService(svc sms.ISMSService, limiter retelimiter.Limiter) sms.ISMSService {
	return &RatelimitSMSService{
		svc:     svc,
		limiter: limiter,
	}
}

// Send
//
//	@Description: 药医不死病 ，佛都有缘人
//	@receiver s
//	@param ctx
//	@param tpl
//	@param args
//	@param numbers
//	@return error
func (s *RatelimitSMSService) Send(ctx context.Context, tpl string, args []string, numbers ...string) error {
	// 前面加功能
	// 应该用装饰器模式
	limited, err := s.limiter.Limited(ctx, "sms:tencent")
	if err != nil {
		// 系统错误
		// 如果下游很垃圾，那就保守策略，一定要限流，
		// 如果下游很强大，或者业务的可用性要求很高，那么就容错策略，不限流
		return fmt.Errorf("短信判断是否限流出现问题:%w", err) // 包一下这个错误
	}
	if limited {
		// 被限流了
		return errLimited
	}
	err = s.svc.Send(ctx, tpl, args, numbers...)
	// 后面加功能
	return err
}
