package temp

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-27 18:25

type IDemoInterface interface {
	funcA()
}

// 原有的实现
type DemoImpl struct {
}

// 构造方法
func NewDemoImpl() IDemoInterface {
	return &DemoImpl{}
}

// 原有实现
func (s *DemoImpl) funcA() {
	// 这里是老代码，老功能
}
