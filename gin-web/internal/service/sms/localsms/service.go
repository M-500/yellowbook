package localsms

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-27 15:56
import (
	"context"
	"fmt"
	"gin-web/internal/service/sms"
)

type Service struct {
}

func NewService() sms.ISMSService {
	return &Service{}
}

func (s *Service) Send(ctx context.Context, tpl string, args []string, numbers ...string) error {
	fmt.Println("验证码是", args)
	return nil
}
