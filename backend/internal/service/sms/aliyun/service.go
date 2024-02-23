package aliyun

import (
	"context"
	"github.com/alibabacloud-go/tea/tea"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-23 17:13

const endPoint = "dysmsapi.aliyuncs.com"

type Service struct {
	appId    *string
	signName *string
	endPoint *string
	client   *sms.Client
}

func NewService() *Service {
	return &Service{
		endPoint: tea.String(endPoint),
	}
}
func (s *Service) Send(ctx context.Context, tpl string, args []string, numbers ...string) error {
	return nil
}
