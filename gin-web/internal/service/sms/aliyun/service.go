//@Author: wulinlin
//@Description:
//@File:  service
//@Version: 1.0.0
//@Date: 2024/02/26 22:15

package aliyun

import (
	"context"
	"gin-web/internal/service/sms"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	"github.com/alibabacloud-go/tea/tea"
)

type ServiceAli struct {
	signName string
	client   *dysmsapi20170525.Client
}

func NewServiceAli(accessKeyId *string, accessKeySecret *string, signName string) sms.ISMSService {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	result, err := dysmsapi20170525.NewClient(config)
	if err != nil {
		panic(err)
	}
	return &ServiceAli{
		signName: signName,
		client:   result,
	}
}

func (s *ServiceAli) Send(ctx context.Context, tpl string, args []string, numbers ...string) error {
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		SignName:      tea.String(s.signName),
		TemplateCode:  tea.String(tpl),
		PhoneNumbers:  tea.String("18574945291"),
		TemplateParam: tea.String("{\"code\":\"1234\"}"),
	}
	_, err := s.client.SendSms(sendSmsRequest)
	if err != nil {
		return err
	}
	return nil
}
