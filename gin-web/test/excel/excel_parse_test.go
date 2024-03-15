package excel

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"testing"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-15 18:11

func TestExcelParse(t *testing.T) {
	f, err := excelize.OpenFile("1.xlsx")
	defer f.Close()
	if err != nil {
		println(err.Error())
		return
	}
	sheetName := f.GetSheetName(0)

	// 获取第一个工作表的行数和列数
	rows, err := f.GetRows(sheetName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(rows))
	// 遍历每一行并打印
	for i, row := range rows {
		fmt.Println(i, row)
		//if i == 0 {
		//	continue
		//}
		//for _, cell := range row {
		//	fmt.Printf("%s\t", cell)
		//}
		//fmt.Println()
	}
}
