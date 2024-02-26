//@Author: wulinlin
//@Description:
//@File:  builder
//@Version: 1.0.0
//@Date: 2024/02/26 21:48

package ratelimitV2

import (
	"fmt"
	"gin-web/pkg/ginx/retelimiter"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Builder struct {
	prefix  string
	limiter retelimiter.Limiter
}

func NewBuilder(limiter retelimiter.Limiter) *Builder {
	return &Builder{
		prefix:  "",
		limiter: limiter,
	}
}

func (b *Builder) Prefix(prefix string) *Builder {
	b.prefix = prefix
	return b
}

func (b *Builder) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		limited, err := b.limit(ctx)
		if err != nil {
			log.Println(err)
			// 这一步很有意思，就是如果这边出错了
			// 要怎么办？
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if limited {
			log.Println(err)
			ctx.AbortWithStatus(http.StatusTooManyRequests)
			return
		}
		ctx.Next()
	}
}

func (b *Builder) limit(ctx *gin.Context) (bool, error) {
	key := fmt.Sprintf("%s:%s", b.prefix, ctx.ClientIP())
	return b.limiter.Limited(ctx, key)
}
