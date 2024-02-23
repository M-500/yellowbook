package aliyun

import (
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-23 17:13

type Service struct {
	appId    *string
	signName *string
	endPoint *string
	client   *sms.Client
}
