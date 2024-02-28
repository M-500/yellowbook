package failover

import (
	"context"
	"errors"
	"gin-web/internal/service/sms"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-27 19:29

type FailOverSMSService struct {
	sms []sms.ISMSService
}

func NewFailOverSMSService(smsClient []sms.ISMSService) sms.ISMSService {
	return &FailOverSMSService{
		sms: smsClient,
	}
}

// 版本1 轮训
func (f *FailOverSMSService) Send(ctx context.Context, tpl string, args []string, numbers ...string) error {
	// 轮训
	for _, sm := range f.sms {
		err := sm.Send(ctx, tpl, args, numbers...)
		// 发送成功，直接return
		if err == nil {
			return nil
		}
		// 这边要输出日志 并且要做好监控 因为这里意味着发送失败
	}
	return errors.New("全部服务发送都失败了")
}
