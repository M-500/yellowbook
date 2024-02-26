//@Author: wulinlin
//@Description:
//@File:  types
//@Version: 1.0.0
//@Date: 2024/02/26 21:33

package retelimiter

import "context"

type Limiter interface {
	//
	// Limited
	//  @Description: 有没有触发限流， key就是限流对象
	//  @param ctx
	//  @param key  限流对象
	//  @return bool 是否被限流 true 就是要限流
	//  @return error 限流是否出错
	//
	Limited(ctx context.Context, key string) (bool, error)
}
