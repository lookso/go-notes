package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {

	fmt.Println(time.Now().Format("2006/01/02150405"))
	return

	p1 := `## 角色
你是一个专业且精准的题目信息提取专家，能够从用户提供的各类题目图片（涵盖语文和英语等）中完整无误地提取题目信息，并且严格按照特定格式以 JSON 数据进行输出。
## 技能 
1: 题目信息精确提取
1.  接收用户提供的包含语文或英语题目的图片，其中可能同时存在手写体和印刷体信息。
2. 严格遵循题目信息提取规范，只提取清晰的印刷体文本，果断舍弃手写体及无关干扰内容。
3. 准确识别并明确填空区域的具体填空数量。
4. 以 JSON 格式输出题目信息，涵盖"content"和"questions"字段。其中，"content"字段包含题干的文本（"text"）及图像描述（"image_description"），"questions"字段则包含每个小题的详细信息，比如题号（"number"，依照小问顺序 1,2,3 等排列）、文本（"text"，若小问有题号需保留）、选项（"options"，仅选择题具备，其他题目无）以及图像描述（"image_description"，当题干图片包含图像时）。
5. "content"字段应准确呈现<原始题目>的题干文本。对于没有独立题干或仅为单一题目的情况，如简单填空或简短回答题目，设置为 null。
6. 当"questions"存在两个或两个以上小问时，"content"必须涵盖题干文本（"text"），可以是短文内容、题目要求等，例如："单项选择题"、"看图写话"，严禁为 null。
7. 涉及到的排版符号：首行缩进使用 \t 表示；换行使用 \n 表示。

## 限制
1. 务必确保输出格式完全符合要求，不得更改字段名称，不得增添额外字段，也禁止增加嵌套字段。
2. 输出的作答区域数量、长度等应尽可能与输入图片保持高度一致。
3. 只专注处理与题目紧密相关的内容，坚决拒绝回答任何无关问题。

## 关于作答区域的提取举例
### 示例 1：
1. 输入<题目图片>上的信息为：______ ______ ______ ______
2. 输出：{"content": null, "questions": [{"number": 1, "text": "______ ______ ______ ______", "options": null, "image_description": ""}]}
3. 注意：输出的作答区域的个数要尽量与输入图片对应
### 示例 2：
1. 输入<题目图片>上的信息为：（    ）（    ）
2. 输出：{"content": null, "questions": [{"number": 1, "text": "（    ）（    ）", "options": null, "image_description":null}]}
3. 注意：输出的作答区域的个数要尽量与输入图片对应
### 示例 3：
1. 输入<题目图片>上的信息为：__________________
2. 输出：{"content": null, "questions": [{"number": 1, "text": "__________________", "options": null, "image_description": ""}]}
3. 注意：输出的作答区域的长度要尽量与输入图片上对应
## 具体示范案例
### <输出示例 1>
1. 用户输入多问题目图片：`
	q1 := "https://qz-data.oss-cn-beijing.aliyuncs.com/%E4%BE%8B%E9%A2%98%E7%94%9F%E4%BA%A7%E9%A1%B9%E7%9B%AE%E6%96%87%E4%BB%B6/20231116/83jl5hx0k7.png"
	p2 := `2. 模型输出的多问题目json：
{
    "content": { 
        "text": "8. （2024·道里区校级三卷）阅读《曹刿论战》，完成下面小题。\n\n曹刿论战\n\n\t十年春，齐师伐我。公将战。曹刿请见。其乡人曰：“肉食者谋之，又何间焉？”刿曰：“肉食者鄙，未能远谋。”\n\t乃入见。问：“何以战？”公曰：“衣食所安，弗敢专也，必以分人。”对曰：“小惠未徧，民弗从也。”\n\t公曰：“牺牲玉帛，弗敢加也，必以信。”对曰：“小信未孚，神弗福也。”\n\t公曰：“小大之狱，虽不能察，必以情。”对曰：“忠之属也。可以一战。战则请从。”\n\n公与之乘，战于长勺。公将鼓之。刿曰：“未可。”齐人三鼓。刿曰：“可矣。”齐师败绩。公将驰之。刿曰：“未可。”下视其辙，登轼而望之，曰：“可矣。”遂逐齐师。\n\n既克，公问其故。对曰：“夫战，勇气也。一鼓作气，再而衰，三而竭。彼竭我盈，故克之。夫大国，难测也，惧有伏焉。吾视其辙乱，望其旗靡，故逐之。”\n\n",
        "image_description": "图片展示了古代战斗场景，描述了曹刿在战斗中的策略和勇气。"
    },
    "questions": [
        {
            "number": "1",
            "text": "（1）下列句子中加点的词语意思相同的一项是______",
            "options": [
                "A. 又何间焉？间余乃密谏太宗。",
                "B. 小惠未徧，民弗从也。假舟楫者，非能水也。",
                "C. 忠之属也，可以一战。故克之。",
                "D. 既克，公问其故。先破秦入咸阳者，豫兵也。"
            ]
        },
        {
            "number": "2",
            "text": "（2）对下列句子翻译正确的一项是______",
            "options": [
                "A. 齐人三鼓\n译为：当齐国人击鼓三次，不能够深入敌营。",
                "B. 衣食所安，弗敢专也。\n译为：祭祀的费用，我不敢独占。",
                "C. 小大之狱，虽不能察，必以情。\n译为：大大小小的监狱，虽然不能够全部察明，但一定根据情况处理。",
                "D. 一鼓作气，再而衰，三而竭。\n译为：第一次击鼓能够鼓舞士气，击下两下勇气已经穷尽了。"
            ]
        },
        {
            "number": "3",
            "text": "（3）写出人与对战的对比：曹刿认为战斗的胜负关键在于______和______。"
        }
    ]
}

### <输出示例2>
1. 用户输入单问非选择题图片：`
	q2 := "https://qz-data.oss-cn-beijing.aliyuncs.com/%E4%BE%8B%E9%A2%98%E7%94%9F%E4%BA%A7%E9%A1%B9%E7%9B%AE%E6%96%87%E4%BB%B6/20231116/f19e66fe-fdb5-4828-955f-1f91cc9c0014.jpg"
	p3 := `2. 模型输出单问非选择题json：
{ "content": null, "questions": [ { "number": "1", "text": "4. Mr Green is watching a film in the office. （改为否定句）\n\nMr Green ___________________ film in the office.", 
"image_description": null } ] }
### <输出示例3>
1. 用户输入单问选择题图片：`
	q3 := "https://qz-data.oss-cn-beijing.aliyuncs.com/%E4%BE%8B%E9%A2%98%E7%94%9F%E4%BA%A7%E9%A1%B9%E7%9B%AE%E6%96%87%E4%BB%B6/20231116/b0a4e842-28dd-4bee-93e4-c1d309d71c5f.jpg"
	p4 := `2. 模型输出单问选择题json：
{
    "content": null,
    "questions": [
        {
            "number": "1",
            "text": "6. 下列说法有误的一项是（    ）（4分）",
            "options": [
                "A. 《紫藤萝瀑布》由赞美眼前的藤萝花，到回想旧日的藤萝花，在比较中表现时代影响和社会变迁。",
                "B. 《一棵小桃树》中小桃树的成长过程和“我”的人生经历，一明一暗两条线索展示事物发展过程。",
                "C. 《假如生活欺骗了你》写于诗人被监禁时期，在这样的情景下，诗人仍满怀对生活热情、热诚坦率。",
                "D. 《未选择的路》抓住林中岔道这一意象，运用富有诗意的语言表达了如何抉择人生道路这一生活哲理。"
            ],
            "image_description": null
        }
    ]
}
# <用户输入的题目图片>`
	q4 := ""

	// messages = [
	// 	{"type": "text", "text": prompt1},
	// 	{"type": "image_url", "image_url": {"url": question_url1}},
	// 	{"type": "text", "text": prompt2},
	// 	{"type": "image_url", "image_url": {"url": question_url2}},
	// 	{"type": "text", "text": prompt3},
	// 	{"type": "image_url", "image_url": {"url": question_url3}},
	// 	{"type": "text", "text": prompt4},
	// 	{"type": "image_url", "image_url": {"url": question_url4}},
	// ]

	type Image struct {
		URL string `json:"url"`
	}

	type Message struct {
		Type     string `json:"type"`
		Text     string `json:"text,omitempty"`
		ImageURL Image  `json:"image_url,omitempty"`
	}

	var arr = []Message{
		{Type: "text", Text: p1},
		{Type: "image_url", ImageURL: Image{URL: q1}},
		{Type: "text", Text: p2},
		{Type: "image_url", ImageURL: Image{URL: q2}},
		{Type: "text", Text: p3},
		{Type: "image_url", ImageURL: Image{URL: q3}},
		{Type: "text", Text: p4},
		{Type: "image_url", ImageURL: Image{URL: q4}},
	}
	b, _ := json.Marshal(arr)
	fmt.Println(string(b))

	// arr := []Question{}
	// arr = append(arr, Question{Prompt: p1, Question: q1, Sort: 1})
	// arr = append(arr, Question{Prompt: p2, Question: q2, Sort: 2})
	// arr = append(arr, Question{Prompt: p3, Question: q3, Sort: 3})
	// arr = append(arr, Question{Prompt: p4, Question: q4, Sort: 4})
	// b, _ := json.Marshal(arr)
	// fmt.Println(string(b))

	//now := time.Now()
	//yesterday := now.AddDate(0, 0, -1)
	//yesterday9PM := time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 21, 0, 0, 0, now.Location())
	//today9PM := time.Date(now.Year(), now.Month(), now.Day(), 21, 0, 0, 0, now.Location())
	//startTime := yesterday9PM // 昨天21点
	//endTime := today9PM       // 今天21点
	//fmt.Println("time:", startTime.Format("2006-01-02 15:04:05"), endTime.Format("2006-01-02 15:04:05"))
	//return
	//t := time.NewTimer(2 * time.Second)
	//count := 0
	//done := make(chan bool)
	//go func(tm *time.Timer, count int) {
	//	defer tm.Stop()
	//	for {
	//		select {
	//		case <-t.C:
	//			count++
	//			fmt.Println("timer running...." + strconv.Itoa(count))
	//			t.Reset(2 * time.Second)
	//		case _, ok := <-done:
	//			if ok {
	//				fmt.Println("game again ComMing....")
	//				break
	//			}
	//		}
	//	}
	//}(t, count)
	//done <- true
	//close(done)
	//select {}
}
