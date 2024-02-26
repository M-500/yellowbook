package demo1

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-26 11:44

type Link struct {
	Data int
	Next *Link
}

func Range(head *Link) []int {
	var ans []int
	for head != nil {
		ans = append(ans, head.Data)
		head = head.Next
	}
	return ans
}
