package tools

import "time"

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-20 14:16

// DateTimeUtils 是一个包含常用日期时间转换功能的工具类
type DateTimeUtils struct{}

// ParseDateTime 将字符串解析为时间，并返回 time.Time 对象
func (d *DateTimeUtils) ParseDateTime(datetimeStr string) (time.Time, error) {
	return time.Parse(time.RFC3339, datetimeStr)
}

func (d *DateTimeUtils) ParseDate(datetimeStr string) (time.Time, error) {
	return time.Parse("2006-01-02", datetimeStr)
}

// FormatDateTime 将时间格式化为指定格式的字符串，并返回字符串
func (d *DateTimeUtils) FormatDateTime(datetime time.Time, layout string) string {
	return datetime.Format(layout)
}

// GetCurrentDateTime 获取当前时间，并返回 time.Time 对象
func (d *DateTimeUtils) GetCurrentDateTime() time.Time {
	return time.Now()
}

// GetCurrentDateTimeStr 获取当前时间，并返回格式化后的字符串
func (d *DateTimeUtils) GetCurrentDateTimeStr(layout string) string {
	return d.GetCurrentDateTime().Format(layout)
}
