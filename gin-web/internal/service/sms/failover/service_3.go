package failover

import (
	"context"
	"gin-web/internal/service/sms"
	"sync/atomic"
)

// @Description 连续N个请求超时的实现
// @Author 代码小学生王木木
// @Date 2024-02-27 20:39

type TimeoutFailOverService struct {
	cnt  int32             // 连续超时的个数
	svcs []sms.ISMSService // 服务商礼拜哦
	idx  int32
	// 阈值 连续超时超过这个数字，就需要切换
	threshold int32
}

func NewTimeoutFailOverService() sms.ISMSService {
	return &TimeoutFailOverService{}
}

func (f *TimeoutFailOverService) Send(ctx context.Context, tpl string, args []string, numbers ...string) error {
	idx := atomic.LoadInt32(&f.idx) // 这是一个原子 +1 操作 是轻量级的并发控制工具
	cnt := atomic.LoadInt32(&f.cnt) // 这是一个原子 +1 操作 是轻量级的并发控制工具
	//
	if cnt > f.threshold {
		// 切换
		// 重新计算新的下标 往后挪
		newIdx := (idx + 1) % int32(len(f.svcs))
		if atomic.CompareAndSwapInt32(&f.idx, idx, newIdx) {
			// 这意味着 我成功往后挪了一位
			atomic.StoreInt32(&f.cnt, 0)
		} else {
			// 出现并发了，别人换成功了
			//idx = newIdx // 写法1
			idx = atomic.LoadInt32(&f.idx) // 两种写法都可以
		}
	}
	svc := f.svcs[idx]
	err := svc.Send(ctx, tpl, args, numbers...)
	switch err {
	case context.DeadlineExceeded, context.Canceled: // 超时
		atomic.AddInt32(&f.cnt, 1)
	case nil:
		atomic.StoreInt32(&f.cnt, 0) // 你的连续状态被打断了
	default:
		// 不知道是什么错误  你可以考虑换下一个
		return err
	}
	return nil
}
