package service

import (
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"os"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-15 18:34

type IExcelParse interface {
}

type ExcelParserService struct {
}

func NewExcelParserService() *ExcelParserService {
	return &ExcelParserService{}
}

func (e *ExcelParserService) ParserExcel(path string) error {
	// 校验文件是否存在
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return errors.New("文件路径不存在")
	}
	if err != nil {
		return errors.New("系统错误，校验文件错误")
	}

	//TODO 安总说这里还需要校验文件后缀名，文件大小，文件格式等信息，以及Excel模板是否合法等信息

	// 解析excel
	f, err := excelize.OpenFile(path)
	defer f.Close()
	if err != nil {
		//println(err.Error())
		return errors.New("Excel 读取错误")
	}
	sheetName := f.GetSheetName(0)
	// 获取第一个工作表的行数和列数
	rows, err := f.GetRows(sheetName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(rows))
	titleMap := make(map[string]int, 10)
	for i, row := range rows {
		if i == 0 {
			continue
		}
		if i == 1 {
			for ik, cell := range row {
				titleMap[cell] = ik
			}
		}

		//for _, cell := range row {
		//	fmt.Printf("%s\n", cell)
		//}
		if i > 1 {
			break
		}
		fmt.Println(titleMap)
	}
	// 异步处理，批量提交
	return nil
}
