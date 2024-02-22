package tencent

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-22 21:12
import (
	"context"
	"fmt"
	"github.com/ecodeclub/ekit"
	"github.com/ecodeclub/ekit/slice"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111" // 引入sms
)

type Service struct {
	appId    *string
	signName *string
	client   *sms.Client
}

func NewService(appId string, signName string, client *sms.Client) *Service {
	return &Service{
		appId:    ekit.ToPtr[string](appId),
		signName: ekit.ToPtr[string](signName),
		client:   client,
	}
}

func (s *Service) Send(ctx context.Context, tpl string, args []string, numbers ...string) error {
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
