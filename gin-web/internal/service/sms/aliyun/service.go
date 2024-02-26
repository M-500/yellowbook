//@Author: wulinlin
//@Description:
//@File:  service
//@Version: 1.0.0
//@Date: 2024/02/26 22:15

package aliyun

import (
	"context"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

type ServiceAli struct {
	client *sms.Client
}

func NewServiceAli(client *sms.Client) *ServiceAli {
	return &ServiceAli{
		client: client,
	}
}

func (s *ServiceAli) Send(ctx context.Context, tpl string, args []string, numbers ...string) error {

	return nil
}
