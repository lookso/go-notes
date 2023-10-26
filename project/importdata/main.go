package main

import (
	"errors"
	"fmt"
	"go-notes/project/importdata/es"

	"github.com/spf13/cast"
	"github.com/xuri/excelize/v2"
)

var QuestionTypeMap = map[string]int{
	"邮件":    1,
	"看图写故事": 2,
	"明信片":   3,
	"便条":    4,
	"议论文":   5,
	"续写故事":  6,
}

func getExcel(excel string) (file *excelize.File) {
	f, err := excelize.OpenFile(excel)
	defer f.Close()
	if err != nil {
		fmt.Println("打开文件失败")
		panic(err)
	}
	list := f.GetSheetList()
	if len(list) == 0 {
		fmt.Println("文件格式不正确")
		panic(err)
	}
	//fmt.Println("sheet list", list)
	return f
}

func ImportEnglishToEs(file *excelize.File) {
	list := file.GetSheetList()
	rows, err := file.GetRows(list[0])
	if err != nil {
		panic(err)
	}
	if len(rows) < 1 {
		panic(errors.New("sheet 行数缺少"))
	}
	rows = rows[1:]

	var questionList = make([]es.Question, 0, 10)
	var failQuestionIdArr = make([]int, 0, 10)
	for _, row := range rows {
		// 只导入文本
		if row[7] == "是" {
			continue
		}
		if row[1] == "" {
			failQuestionIdArr = append(failQuestionIdArr, cast.ToInt(row[6]))
			fmt.Print("222")
			break
		}
		if row[4] == "" && row[7] == "是" {
			fmt.Print("333")
			failQuestionIdArr = append(failQuestionIdArr, cast.ToInt(row[6]))
			break
		}
		if row[4] != "" && row[7] == "否" {
			fmt.Print("4444")
			failQuestionIdArr = append(failQuestionIdArr, cast.ToInt(row[6]))
			break
		}
		var relationPicture = 0
		if row[7] == "是" {
			relationPicture = 1
		}
		if row[7] == "否" {
			relationPicture = 2
		}
		if relationPicture == 0 {
			fmt.Print("5555")
			failQuestionIdArr = append(failQuestionIdArr, cast.ToInt(row[6]))
			break
		}
		correntType, ok := QuestionTypeMap[row[2]]
		if !ok {
			fmt.Println("questionType:", row[2])
			failQuestionIdArr = append(failQuestionIdArr, cast.ToInt(row[6]))
			break
		}
		questionList = append(questionList, es.Question{
			Description:     row[8],
			Detail:          row[5],
			Order:           row[0],
			QuestionID:      cast.ToInt(row[10]),
			QuestionType:    cast.ToInt(row[2]),
			RelationPicture: relationPicture,
			Source:          row[3],
			Type:            cast.ToString(correntType),
			URL:             row[4],
		})
	}
	println("questionList", len(questionList))

	if len(failQuestionIdArr) > 0 {
		fmt.Println("failArr", failQuestionIdArr)
		return
	}

	esClient := es.NewEs()
	//if err := esClient.CreateIndex(); err != nil {
	//	fmt.Println("createIndex err", err)
	//	return
	//}
	//if err = esClient.Insert(questionList); err != nil {
	//	fmt.Println("insert err", err)
	//}
}

func main() {
	file := ""
	excel := getExcel(file)
	defer excel.Close()

	ImportEnglishToEs(excel)
}
