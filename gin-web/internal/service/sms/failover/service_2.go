package failover

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-27 19:40

import (
	"context"
	"errors"
	"gin-web/internal/service/sms"
	"sync/atomic"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-27 19:29

type FailOverSMSServiceV1 struct {
	sms []sms.ISMSService
	idx uint64 //
}

func NewFailOverSMSServiceV1(smsClient []sms.ISMSService) sms.ISMSService {
	return &FailOverSMSServiceV1{
		sms: smsClient,
	}
}

// 这么做的目的是不要让他从0 开始 而是从idx的下一个开始
func (f *FailOverSMSServiceV1) Send(ctx context.Context, tpl string, args []string, numbers ...string) error {
	// 取下一个节点来作为起始节点
	idx := atomic.AddUint64(&f.idx, 1) // 这是一个原子操作 是轻量级的并发控制工具
	length := uint64(len(f.sms))
	for i := idx; i < idx+length; i++ {
		svc := f.sms[int(i%length)] // 取模为了防止下标越界
		err := svc.Send(ctx, tpl, args, numbers...)
		switch err {
		case nil:
			return nil
		case context.DeadlineExceeded, context.Canceled: // 被取消或者被主动取消
			return nil
		default:
			return err
		}
	}
	return errors.New("全部服务发送都失败了")
}
