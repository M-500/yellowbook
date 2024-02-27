//@Author: wulinlin
//@Description:
//@File:  types
//@Version: 1.0.0
//@Date: 2024/02/26 22:00

package sms

import "context"

type ISMSService interface {
	Send(ctx context.Context, tpl string, args []string, numbers ...string) error
}
