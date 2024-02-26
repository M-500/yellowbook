package gin_demo

import (
	"github.com/gin-gonic/gin"
	"testing"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-26 15:25

func TestUserHandler_Login(t *testing.T) {
	testCase := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			// 我在这里怎么拿到HTTP的响应呢？
			handler := NewUserHandler(nil)
			ctx := &gin.Context{}
			handler.Login(ctx)
		})
	}
}
