package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-26 16:49
func TestUserHandler_Login(t *testing.T) {
	testCase := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	req, err := http.NewRequest(http.MethodPost, "/user/pwd-login", bytes.NewBuffer([]byte(`
{
	"username":"wulinlin",
	"pwd":"123465"
}
`)))
	require.NoError(t, err)
	//if err != nil {
	//
	//}
	//这里可以继续使用req了
	// 模拟响应数据
	resp := httptest.NewRecorder()
	resp.Code = http.StatusOK
	resp.Body = bytes.NewBuffer([]byte(``))

	// 不可能走下面的流程的
	//// 调用handler 这一定会遇到初始化第三方的问题  一连MySQL就 插入数据 测试数据等等等等
	//db, err := gorm.Open(mysql.Open(""))
	//dao := NewUserDAO(db)
	//svc := NewUserService(dao)
	//h := NewUserHandler(svc)

	// 使用mock工具

	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			// 我在这里怎么拿到HTTP的响应呢？
			handler := NewUserHandler(nil)
			ctx := &gin.Context{}
			handler.Login(ctx)
		})
	}
}
