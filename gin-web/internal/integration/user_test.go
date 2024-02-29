package integration

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-29 19:42

// UserTestSuit
// @Description: 测试套件  单元测试的一种组织方式
type UserTestSuit struct {
	suite.Suite
}
type Article struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Result[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

func (s *UserTestSuit) TestHelloWorld() {
	s.T().Log("这里介绍一下测试套件")
}

func (s *UserTestSuit) TestEdit() {
	t := s.T()
	testCase := []struct {
		name string

		// 集成测试验证数据
		before func(t *testing.T)
		after  func(t *testing.T)
		art    Article

		wantCode int // 对应HTTP的响应码
		wantRes  Result[int64]
	}{}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			// 构造请求
			// 执行
			// 验证结果
			tc.before(t)
		})
	}
}

func TestUser(t *testing.T) {
	suite.Run(t, &UserTestSuit{})
}
