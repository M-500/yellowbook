package domain

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-20 11:41

type ParserLogger struct {
	FileName  string  `json:"file_name"`
	FileSize  float64 `json:"file_size"`
	Rows      uint32  `json:"rows"`
	IOTime    int64   `json:"io_time"`
	ParseTime int64   `json:"parse_time"`
	WaitTime  int64   `json:"wait_time"`
	Mark      string  `json:"mark"`
}
