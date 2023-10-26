package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"strings"
)

func main() {
	numbers := "12\n77\n90\n92\n94\n100\n113\n126\n128\n141\n142\n147\n148\n149\n150\n153\n155\n156\n161\n164\n165\n173\n174\n181\n182\n188\n192\n199\n200\n201\n202\n203\n221\n229\n232\n234\n247\n254\n258\n267\n271\n277\n288\n289\n294\n298\n299\n309\n311\n314\n317\n318\n319\n326\n336\n341\n356\n363\n366\n368\n370\n378\n381\n388\n390\n391\n392\n393\n395\n398\n402\n403\n404\n405\n408\n412\n413\n415\n417\n420\n424"
	numbersArray := strings.Split(numbers, "\n")
	for i, number := range numbersArray {
		numbersArray[i] = number + ","
	}
	result := strings.Join(numbersArray, "\n")
	fmt.Println(result)

	return

	r := gin.Default()

	r.GET("/download", func(c *gin.Context) {
		// 创建一个新的Excel文件
		file := excelize.NewFile()

		// 向文件中写入数据
		file.SetCellValue("Sheet1", "A1", "Hello")
		file.SetCellValue("Sheet1", "A2", "Hello")
		file.SetCellValue("Sheet1", "B1", "World")

		// 将文件保存到本地
		file.SaveAs("sample.xlsx")

		// 将文件内容作为响应体发送给客户端
		c.File("sample.xlsx")

		// 设置响应头，指定文件名和文件类型
		c.Header("Content-Disposition", "attachment; filename=sample.xlsx")
		c.Header("Content-Type", "application/octet-stream")


	})

	r.Run(":9090")
}