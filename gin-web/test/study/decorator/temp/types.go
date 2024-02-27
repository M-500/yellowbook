package temp

import "fmt"

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-27 18:27

// DemoPlusImpl
// @Description: 扩展原有实现
type DemoPlusImpl struct {
	oldImpl IDemoInterface
}

// 构造方法
func NewDemoPlusImpl() IDemoInterface {
	return &DemoPlusImpl{}
}

// 装饰者也实现了接口的所有方法，实现扩展
func (s *DemoPlusImpl) funcA() {
	fmt.Println("这里随便增加功能")
	s.oldImpl.funcA() // 这里还是老功能，老方法
	fmt.Println("这里随便增加功能")
}
