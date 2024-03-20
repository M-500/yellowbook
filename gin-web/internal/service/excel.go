package service

import (
	"context"
	"errors"
	"fmt"
	"gin-web/internal/domain"
	"gin-web/internal/repository"
	"github.com/xuri/excelize/v2"
	"log"
	"os"
	"sync"
	"time"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-15 18:34

type IExcelParseService interface {
	ParserExcel(ctx context.Context, path string) error
	ParserExcelV1(ctx context.Context, path string) error
}

type ExcelParserService struct {
	repo repository.ICompany
}

func NewExcelParserService(repo repository.ICompany) IExcelParseService {
	return &ExcelParserService{
		repo: repo,
	}
}

func (e *ExcelParserService) parser(titleMap map[string]int, row []string) domain.Company {
	company := domain.Company{}
	val, ok := titleMap["企业名称"]
	if ok {
		company.EnterpriseName = row[val]
	}
	val, ok = titleMap["登记状态"]
	if ok {
		company.RegistrationStatus = row[val]
	}
	val, ok = titleMap["法定代表人"]
	if ok {
		company.LegalRepresentative = row[val]
	}
	val, ok = titleMap["注册资本"]
	if ok {
		company.RegisteredCapital = row[val]
	}
	//val, ok = titleMap["成立日期"]
	//if ok {
	//	company.EstablishmentDate = row[val]
	//}
	val, ok = titleMap["统一社会信用代码"]
	if ok {
		company.UnifiedSocialCreditCode = row[val]
	}
	val, ok = titleMap["企业注册地址"]
	if ok {
		company.EnterpriseRegistrationAddress = row[val]
	}
	val, ok = titleMap["电话"]
	if ok {
		company.Phone = row[val]
	}
	val, ok = titleMap["更多电话"]
	if ok {
		company.MorePhone = row[val]
	}
	val, ok = titleMap["邮箱"]
	if ok {
		company.Email = row[val]
	}
	val, ok = titleMap["更多邮箱"]
	if ok {
		company.MoreEmail = row[val]
	}
	return company
}

func (e *ExcelParserService) ParserExcel(ctx context.Context, path string) error {
	// 校验文件是否存在
	fileInfo, err := os.Stat(path)
	if os.IsNotExist(err) {
		return errors.New("文件路径不存在")
	}
	if err != nil {
		return errors.New("系统错误，校验文件错误")
	}
	parseLog := domain.ParserLogger{
		FileName: path,
		FileSize: float64(fileInfo.Size()) / 1024 / 1024, // 获取文件大小
	}
	//TODO 按理说这里还需要校验文件后缀名，文件大小，文件格式等信息，以及Excel模板是否合法等信息
	start := time.Now()
	// 解析excel
	f, err := excelize.OpenFile(path)
	defer f.Close()
	if err != nil {
		//println(err.Error())
		return errors.New("Excel 读取错误")
	}
	parseLog.IOTime = time.Since(start).Milliseconds() // IO读取时间
	sheetName := f.GetSheetName(0)
	// 获取第一个工作表的行数和列数
	rows, err := f.GetRows(sheetName)
	if err != nil {
		log.Fatal(err)
	}
	parseLog.Rows = uint32(len(rows)) // 记录文件行数
	titleMap := make(map[string]int, 20)
	start1 := time.Now()
	for i, row := range rows {
		if i == 0 {
			continue
		}
		if i == 1 {
			// 第一次进来遍历第一行，就可以获取map表结构
			for ik, cell := range row {
				titleMap[cell] = ik
			}
			continue
		}
		company := e.parser(titleMap, row)
		err = e.repo.Create(ctx, company)
		if err != nil {
			//return err
		}
	}
	parseLog.ParseTime = time.Since(start1).Milliseconds()
	parseLog.WaitTime = time.Since(start).Milliseconds()

	parseLog.Mark = "完全同步"
	err = e.repo.CreateLog(ctx, parseLog)

	if err != nil {
		fmt.Println("妈的 ，记录日志出错了", err)
	}
	return nil
}

func (e *ExcelParserService) ParserExcelV1(ctx context.Context, path string) error {
	// 校验文件是否存在
	fileInfo, err := os.Stat(path)
	if os.IsNotExist(err) {
		return errors.New("文件路径不存在")
	}
	if err != nil {
		return errors.New("系统错误，校验文件错误")
	}
	parseLog := domain.ParserLogger{
		FileName: path,
		FileSize: float64(fileInfo.Size()) / 1024 / 1024, // 获取文件大小
	}
	//TODO 按理说这里还需要校验文件后缀名，文件大小，文件格式等信息，以及Excel模板是否合法等信息
	start := time.Now()
	// 解析excel
	f, err := excelize.OpenFile(path)
	defer f.Close()
	if err != nil {
		//println(err.Error())
		return errors.New("Excel 读取错误")
	}
	parseLog.IOTime = time.Since(start).Milliseconds() // IO读取时间
	sheetName := f.GetSheetName(0)
	// 获取第一个工作表的行数和列数
	rows, err := f.GetRows(sheetName)
	if err != nil {
		log.Fatal(err)
	}
	parseLog.Rows = uint32(len(rows)) // 记录文件行数
	titleMap := make(map[string]int, 20)
	start1 := time.Now()
	var wg sync.WaitGroup
	for i, row := range rows {
		if i == 0 {
			continue
		}
		if i == 1 {
			// 第一次进来遍历第一行，就可以获取map表结构
			for ik, cell := range row {
				titleMap[cell] = ik
			}
			continue
		}
		wg.Add(1)
		go func(row []string) {
			defer wg.Done()
			company := e.parser(titleMap, row)
			err = e.repo.Create(ctx, company)
			if err != nil {
				fmt.Println("我丢", err)
				//return err
			}
		}(row)
	}
	wg.Wait()
	parseLog.ParseTime = time.Since(start1).Milliseconds()
	parseLog.WaitTime = time.Since(start).Milliseconds()
	parseLog.Mark = "完全异步"
	err = e.repo.CreateLog(ctx, parseLog)
	if err != nil {
		fmt.Println("妈的 ，记录日志出错了", err)
	}
	return nil
}
