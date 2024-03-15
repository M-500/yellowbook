package service

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-15 18:34

type IExcelParse interface {
}

type ExcelParser struct {
}

func NewExcelParser() *ExcelParser {
	return &ExcelParser{}
}

func (e *ExcelParser) ParserExcel() {

}
