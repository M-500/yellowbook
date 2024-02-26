//@Author: wulinlin
//@Description:
//@File:  service
//@Version: 1.0.0
//@Date: 2024/02/26 21:59

package tencent

import (
	"context"
	"fmt"
	"gin-web/pkg/ginx/retelimiter"
	"github.com/ecodeclub/ekit"
	"github.com/ecodeclub/ekit/slice"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111" // 引入sms
)

type Service struct {
	appId    *string
	signName *string
	client   *sms.Client
	limiter  retelimiter.Limiter
}

func NewService(appId string, signName string, client *sms.Client, limiter retelimiter.Limiter) *Service {
	return &Service{
		appId:    ekit.ToPtr[string](appId),
		signName: ekit.ToPtr[string](signName),
		client:   client,
		limiter:  limiter,
	}
}

func (s *Service) Send(ctx context.Context, tpl string, args []string, numbers ...string) error {
	// 应该用装饰器模式
	limited, err := s.limiter.Limited(ctx, "sms:tencent")
	if err != nil {
		// 系统错误
		// 如果下游很垃圾，那就保守策略，一定要限流，
		// 如果下游很强大，或者业务的可用性要求很高，那么就容错策略，不限流
		return fmt.Errorf("短信判断是否限流出现问题:%w", err) // 包一下这个错误
	}
	if !limited {
		// 被限流了
		return fmt.Errorf("触发了限流")
	}
	req := sms.NewSendSmsRequest()
	req.SmsSdkAppId = s.appId
	req.SignName = s.signName
	req.TemplateId = ekit.ToPtr[string](tpl)
	req.PhoneNumberSet = s.toStringPtrSlice(numbers)
	req.TemplateParamSet = s.toStringPtrSlice(args)
	response, err := s.client.SendSms(req)
	if err != nil {
		return err
	}
	for _, status := range response.Response.SendStatusSet {
		if status.Code == nil || *(status.Code) == "Ok" {
			return fmt.Errorf("发送翻新失败%s,%s,", *status.Code, *status.Message)
		}
	}
	return nil
}

func (s *Service) toStringPtrSlice(src []string) []*string {
	return slice.Map[string, *string](src, func(idx int, src string) *string {
		return &src
	})
}
