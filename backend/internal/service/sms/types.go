package sms

import "context"

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-22 21:07

type Service interface {
	Send(ctx context.Context, tpl string, args []string, numbers ...string) error
}
