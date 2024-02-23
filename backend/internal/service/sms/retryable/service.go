package retryable

import (
	"backend/internal/service/sms"
	"context"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-23 15:48

// Service
// @Description: 这个要小心并发问题，不能做成单例的形态
type Service struct {
	svc      sms.Service
	retryCnt int
}

func NewService(svc sms.Service, cnt int) *Service {
	return &Service{
		svc:      svc,
		retryCnt: cnt,
	}
}

func (s *Service) Send(ctx context.Context, tpl string, args []string, numbers ...string) error {
	err := s.svc.Send(ctx, tpl, args, numbers...)
	for err != nil && s.retryCnt < 10 {
		err = s.svc.Send(ctx, tpl, args, numbers...)
		s.retryCnt++
		return err
	}
	return nil
}
