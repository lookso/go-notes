package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"sync"
	"time"
)

func getTableName(id int) string {
	// 假设你有10张表，每张表有500万条数据
	tableIndex := (id-1) / 5000000
	fmt.Println("tableIndex:",tableIndex)
	return fmt.Sprintf("question_%03d", tableIndex)
}
func processQuestionMappings(qm []int)  error{
	fmt.Println("qm",qm)
	return nil
}

var wg sync.WaitGroup

func producer(ch chan int, id int) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		ch <- i + id*5
	}
}

func mainGo() {
	ch := make(chan int)
	wg.Add(2)
	go producer(ch, 1)
	go producer(ch, 2)

	go func() {
		wg.Wait()
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}
}


func main() {
	//mainGo()
	//return

	// 更新题目状态
	hour := 4
	mainTaskTimeoutTime := time.Now().Add((time.Duration(hour) * time.Hour))
	fmt.Println(mainTaskTimeoutTime)
	return




	fmt.Println(getTableName(12459445))
	return
	questionMappings := make([]int, 10000000)
	questionMappings=[]int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28}
	pn:=5
	partSize := len(questionMappings) /pn
	var eg errgroup.Group
	//eg.SetLimit(pn)
	fmt.Println("partSize:",partSize)
	for i := 0; i < pn; i++ {
		start := i * partSize
		end := start + partSize
		if i == pn-1 {
			// 最后一个协程处理剩余的所有元素
			end = len(questionMappings)
		}
		part := questionMappings[start:end]
		eg.TryGo(func() error {
			return processQuestionMappings(part)
		})
		fmt.Println("------------------")
	}
	if err := eg.Wait(); err != nil {
		fmt.Println("Encountered error:", err)

	}


	//partSize := len(questionMappings) / 3
	//var eg errgroup.Group
	//eg.SetLimit(3)
	//for i := 0; i < 3; i++ {
	//	start := i * partSize
	//	end := start + partSize
	//	if i == 2 {
	//		// 最后一个协程处理剩余的所有元素
	//		end = len(questionMappings)
	//	}
	//	part := questionMappings[start:end]
	//	eg.TryGo(func() error {
	//		return processQuestionMappings(circeCount, sid, part)
	//	})
	//}
	//if err := eg.Wait(); err != nil {
	//	fmt.Println("Encountered error:", err)
	//	return err
	//}
	//

	//text := `<p class="ocr_text_invisible">答案：</p><div class="root-container android">                        <p><img data-type=""{"x":75,"y":1771,"width":440,"height":120,"rotate":0,"scaleX":1,"scaleY":1,"imgno":"2"}"" id="mosaic-image" src="http://qz-question.oss-cn-beijing.aliyuncs.com/016%2Fa6acd410356c8e524a0b6978a967ee1d.jpg"/></p>                                  </div>`
	//text := `<p class="ocr_text_invisible">答案：</p><div class=""root-container android"">                        <p><img data-type=""""{""x"":20,""y"":1409,""width"":411,""height"":56,""rotate"":0,""scaleX"":1,""scaleY"":1,""imgno"":""1""}"""" id=""mosaic-image"" src=""http://qz-question.oss-cn-beijing.aliyuncs.com/qing%2F2805077a06338e64acb1ecb021150cf1.jpg"" style=""visibility: visible;""/></p>                                                <div class=""collection"">              <div class=""reduceIcon""></div>              <span>已加入错题本</span>            </div>          </div>`
	//re1 := regexp.MustCompile(`<img[^>]+src="([^"]+)"`)
	//
	//questionMatches := re1.FindAllStringSubmatch(text2, -1)
	//if len(questionMatches) > 0 {
	//	fmt.Println(questionMatches[0][1])
	//}else{
	//	re2 := regexp.MustCompile(`<img[^>]+src=""([^"]+)""`)
	//	questionMatchesall := re2.FindAllStringSubmatch(text2, -1)
	//	if len(questionMatchesall) > 0 {
	//		fmt.Println(questionMatchesall[0][1])
	//	}
	//}





	//text := `<p class="ocr_text_invisible">答案：</p><div class=""root-container android"">                        <p><img data-type=""""{""x"":20,""y"":1409,""width"":411,""height"":56,""rotate"":0,""scaleX"":1,""scaleY"":1,""imgno"":""1""}"""" id=""mosaic-image"" src=""http://qz-question.oss-cn-beijing.aliyuncs.com/qing%2F2805077a06338e64acb1ecb021150cf1.jpg"" style=""visibility: visible;""/></p>                                                <div class=""collection"">              <div class=""reduceIcon""></div>              <span>已加入错题本</span>            </div>          </div>`

	//re := regexp.MustCompile(`<img[^>]+src="*([^"]+)"*`)
	//matches := re.FindAllStringSubmatch(text, -1)
	//for _, match := range matches {
	//	if len(match) > 1 {
	//		fmt.Println(match[1]) // 输出图片链接
	//	}
	//}
}