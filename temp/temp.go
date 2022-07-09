package main

import (
	"fmt"
	"math"
)

func round(x float64) int {
	return int(math.Floor(x + 0.5))
}

//var t = "{\"msg_content\":[{\"content_type\":1,\"content\": \"#辅导姓名#,#主讲姓名#您好，咱们班级学习群已提前建好了，拉您进群啦\\n之后的【课程重要通知】和【重要的学习资料】，我都会发在学习群里面，一定要进群噢~\\n/爱心 温馨小建议：\\n/太阳 您可以将咱们的学习群置顶\\n/太阳 进群之后改备注【XX妈妈/爸爸】\\n/太阳 重要的通知我会发群公告，您注意查看哦～\",\"title\":\"\",\"desc\":\"\",\"image_url\":\"\",\"url\":\"\",\"at_ext_userids\":\"\"}]}"
//var t ="{\"msg_content\":[{\"content_type\":1,\"content\": \"#辅导姓名#,#主讲姓名#您好，咱们班级学习群已提前建好了，拉您进群啦\\n之后的【课程重要通知】和【重要的学习资料】，我都会发在学习群里面，一定要进群噢~\\n/爱心 温馨小建议：\\n/太阳 您可以将咱们的学习群置顶\\n/太阳 进群之后改备注【XX妈妈/爸爸】\\n/太阳 重要的通知我会发群公告，您注意查看哦～\",\"title\":\"\",\"desc\":\"\",\"image_url\":\"\",\"url\":\"\",\"at_ext_userids\":\"\"}]}"
var t = "{\"msg_content\":[{\"content_type\":1,\"content\": \"#辅导姓名#,#主讲姓名#您好，咱们班级学习群已提前建好了，拉您进群啦\\n之后的【课程重要通知】和【重要的学习资料】，我都会发在学习群里面，一定要进群噢~\\n/爱心 温馨小建议：\\n/太阳 您可以将咱们的学习群置顶\\n/太阳 进群之后改备注【XX妈妈/爸爸】\\n/太阳 重要的通知我会发群公告，您注意查看哦～\",\"title\":\"\",\"desc\":\"\",\"image_url\":\"\",\"url\":\"\",\"at_ext_userids\":\"\"}]}"

type WxSendMsgList struct {
	MsgContent []MsgContent `json:"msg_content"`
}
type MsgContent struct {
	ContentType  int    `json:"content_type"`
	Content      string `json:"content"`
	Title        string `json:"title"`
	Desc         string `json:"desc"`
	ImageURL     string `json:"image_url"`
	URL          string `json:"url"`
	AtExtUserids string `json:"at_ext_userids"`
}
type WxInfo struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

type Req struct {
	Messages []MessageItem `json:"messages"`
}
type MessageItem struct {
	ExtUserID []int  `json:"ext_userid"`
	Content   string `json:"content"`
}

//func callfunNumBak()  {
//	//qt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10,11,12,13}
//	//for i := 0; i < len(qt); i += 4 {
//	//	if len(qt[i:]) >= 4 {
//	//		fmt.Println(qt[i:][:4])
//	//	} else {
//	//		fmt.Println(qt[i:])
//	//	}
//	//}
//	//fmt.Println("-----")
//
//	itemNum := 4
//	messages := make([]MessageItem, 0)
//	messages = append(messages, MessageItem{
//		ExtUserID: []int{1, 2, 3, 4, 5},
//		Content:   "aaaaa",
//	})
//	messages = append(messages, MessageItem{
//		ExtUserID: []int{100, 200, 300, 400, 500, 600, 700, 800, 900, 1000},
//	})
//	messages = append(messages, MessageItem{
//		ExtUserID: []int{10, 20, 30},
//		Content:   "bbbb",
//	})
//	messages = append(messages, MessageItem{
//		ExtUserID: []int{1},
//	})
//
//	max := 0
//	sn := 0
//	for _, v := range messages {
//		temp := 0
//		//newExtUserID := make([]int, 0)
//		for i := 0; i < len(v.ExtUserID); i += itemNum {
//			sn++
//			temp++
//			//fmt.Printf("k:%d,i:%d,v:%+v\n", k, i, v.ExtUserID[i])
//			//
//			//if len(v.ExtUserID[i:]) >= itemNum {
//			//	fmt.Println(v.ExtUserID[i:][:itemNum])
//			//} else {
//			//	fmt.Println(v.ExtUserID[i:])
//			//}
//		}
//		if max < temp {
//			max = temp
//		}
//	}
//
//	for i := 0; i < max; i++ {
//		newMessages := make([]MessageItem, 0)
//		for _, v := range messages {
//			newExtUserID := make([]int, 0)
//			if i*itemNum > len(v.ExtUserID) {
//				continue
//			} else {
//				if len(v.ExtUserID[i*itemNum:]) > itemNum {
//					newExtUserID = v.ExtUserID[i*itemNum:][:itemNum]
//				} else {
//					newExtUserID = v.ExtUserID[i*itemNum:]
//				}
//			}
//			if len(newExtUserID) == 0 {
//				continue
//			}
//			fmt.Println("newExtUserID", newExtUserID)
//			newMessages = append(newMessages, MessageItem{
//				ExtUserID: newExtUserID,
//			})
//		}
//		fmt.Printf("call:%+d\n", i)
//		fmt.Println("newMessages:", newMessages)
//		fmt.Println("----------------")
//	}
//	fmt.Println("sn", sn)
//	fmt.Println("max", max)
//
//	fmt.Println("------------------")
//
//	/**
//		[1, 2, 3, 4, 5],[100, 200, 300, 400, 500, 600, 700, 800, 900, 1000],[10, 20, 30]
//
//		a:
//		1,2,3,4
//		100,200,300,400
//		10,20,30
//		b:
//		5
//		500, 600, 700, 800
//		0
//		c:
//		0
//		900,1000
//		0
//	**/
//}
func callfunNum() {
	itemNum := 1
	messages := make([]MessageItem, 0)
	//messages = append(messages, MessageItem{
	//	ExtUserID: []int{1, 2, 3, 4, 5},
	//	Content:   "aaaaa",
	//})
	messages = append(messages, MessageItem{
		ExtUserID: []int{100},
		Content:   "bbbb",
	})
	messages = append(messages, MessageItem{
		ExtUserID: []int{1,2},
		Content:   "bbbb",
	})
	messages = append(messages, MessageItem{
		ExtUserID: []int{1,2,3,4,5,6,7,8,8,9,10},
		Content:   "bbbb",
	})
	//messages = append(messages, MessageItem{
	//	ExtUserID: []int{10, 20, 30, 40},
	//	Content:   "ccc",
	//})
	//messages = append(messages, MessageItem{
	//	ExtUserID: []int{1},
	//	Content:   "dddd",
	//})

	max := 0
	for _, v := range messages {
		temp := 0
		for i := 0; i < len(v.ExtUserID); i += itemNum {
			temp++
		}
		if max < temp {
			max = temp
		}
	}

	for i := 0; i < max; i++ {
		newMessages := make([]MessageItem, 0)
		for _, v := range messages {
			newExtUserID := make([]int, 0)
			if i*itemNum > len(v.ExtUserID) {
				continue
			} else {
				if len(v.ExtUserID[i*itemNum:]) > itemNum {
					newExtUserID = v.ExtUserID[i*itemNum:][:itemNum]
				} else {
					newExtUserID = v.ExtUserID[i*itemNum:]
				}
			}
			if len(newExtUserID) == 0 {
				continue
			}
			newMessages = append(newMessages, MessageItem{
				ExtUserID: newExtUserID,
				Content:   v.Content,
			})
		}
		fmt.Printf("call:%+d\n", i)
		fmt.Println("newMessages:", newMessages)
		fmt.Println("----------------")
	}

	/**
		[1, 2, 3, 4, 5],[100, 200, 300, 400, 500, 600, 700, 800, 900, 1000],[10, 20, 30]

		a:
		1,2,3,4
		100,200,300,400
		10,20,30
		b:
		5
		500, 600, 700, 800
		0
		c:
		0
		900,1000
		0
	**/
}

func main() {

	contentTextTypeCount:=0
	for i:=0;i<10;i++{
		if i%2==0 {
			contentTextTypeCount++
		}
	}
	fmt.Println("contentTextTypeCount",contentTextTypeCount)

	callfunNum()

	//var wx = WxInfo{
	//	//Name: "123",
	//	//ID: 12,
	//}
	//fmt.Println(wx.Name, wx.ID)
	//
	//fmt.Println(fmt.Sprintf("%.2f", float64(411)/float64(200)))
	//lf, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(400)/float64(200)), 64)
	//l := math.Ceil(lf)
	//gn := cast.ToInt(l)
	//if gn <= 0 {
	//	gn = 1
	//}
	//fmt.Printf("lf:%+v,l:%+v,gn:%+v\n", lf, l, gn)
	//
	//var n = []int{1, 2, 3}
	//var s = []int{5, 6}
	//var q = 0
	//for k, _ := range n {
	//
	//	for kk, _ := range s {
	//		q = q + 1
	//		fmt.Println(k, kk, q)
	//	}
	//}
	//
	//var c = "{\"msg_content\":[{\"content_type\":1,\"content\": \"多个群的群公告\\n#辅导姓名#,#主讲姓名#您好，咱们班级学习群已提前建好了，拉您进群啦\\n之后的【课程重要通知】和【重要的学习资料】，我都会发在学习群里面，一定要进群噢~\\n/爱心 温馨小建议：\\n/太阳 您可以将咱们的学习群置顶\\n/太阳 进群之后改备注【XX妈妈/爸爸】\\n/太阳 重要的通知我会发群公告，您注意查看哦～\",\"title\":\"\",\"desc\":\"\",\"image_url\":\"\",\"url\":\"\",\"at_ext_userids\":\"\"}]}"
	//newContent := strings.Replace(c, "\\n", "\\\\n", -1)
	//fmt.Println("newContent", newContent)
	//
	//var wxMsg WxSendMsgList
	//err := json.Unmarshal([]byte(t), &wxMsg)
	//fmt.Println("err", err)

	//var s=map[string]int{}
	//s["name"]=10
	//fmt.Println(s["class"])
	//fmt.Println("------")
	//fmt.Println(strings.Split(" ", ","))
	//
	//fmt.Println(strings.Split("", ","))
	//
	//classIDsStr := ""
	//for _, v := range []int{1,2,3,4} {
	//	classIDsStr += strconv.Itoa(v) + ","
	//}
	//classIDsStr=strings.Trim(classIDsStr,",")
	//fmt.Println(classIDsStr)
	//
	//a, b,c := 1, 2,3
	//if c>0 ||!(a > 10 && b > 0) {
	//	fmt.Println("1111111111")
	//}else {
	//	fmt.Println("222222222222222")
	//}
	//var chatName = ""
	//if chatName == "" {
	//	fmt.Println("chatName111", chatName)
	//
	//	for i := 1; i <= 2; i++ {
	//		fmt.Println("chatName222")
	//		// 等待10s,请求青鸟自己的群详情接口
	//		time.Sleep(10 * time.Second)
	//		chatName = "fuck"
	//		if chatName != "" {
	//			break
	//		}
	//	}
	//}
	//if chatName == "" {
	//	fmt.Println("fuck this empty too")
	//}
	//fmt.Println("chatName44", chatName)
	//return

	//	fmt.Println("time.Now().Unix()",time.Now().Unix())
	//	//now := time.Now()
	//	//// 获取10分钟后时间
	//	//t := now.Add(time.Minute * 10)
	//	//fmt.Println(t.Unix())
	//	//lastM := t.Unix() - int64(now.Second())
	//	//fmt.Println(lastM)
	//
	//	x := 0.3
	//	//fmt.Println(round(x))
	//	ceil := math.Ceil(x)
	//	fmt.Println(ceil)
	//	fmt.Println(math.Floor(ceil / 2)) // 2
	//	//fmt.Println(math.Floor(x))        // 1
	//
	//	var arrs = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//	fmt.Println(arrs[0:10])
	//	a := 30
	//	num1, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", float64(a)/float64(60)), 64) // 保留2位小数
	//
	//	fmt.Println("time", num1)
	//	// "2022-05-12 15:35:45"
	//currentTime, _ := time.ParseInLocation("2006-01-02 15:04:05", "2022-05-12 15:36:59", time.Local) //获取当前时间，类型是Go的时间类型Time
	//a := currentTime.Add(-time.Minute * 1)
	//
	//t1 := a.Year()   //年
	//t2 := a.Month()  //月
	//t3 := a.Day()    //日
	//t4 := a.Hour()   //小时
	//t5 := a.Minute() //分钟
	//t6 := 59
	//t7 := 0
	//
	//currentTimeData := time.Date(t1, t2, t3, t4, t5, t6, t7, time.Local) //获取当前时间，返回当前时间Time
	//fmt.Println(a)                                                       //打印结果：2017-04-11 12:52:52.794351777 +0800 CST
	//fmt.Println(t1, t2, t3, t4, t5, t6)                                  //打印结果：2017 April 11 12 52 52
	//fmt.Println(currentTimeData.Unix())
	//p()
}
func p() {
	fmt.Println(88888)
	return
	var isErr *int
	fmt.Println(123123213324)
	defer func() {
		if *isErr == 2 {
			fmt.Println("123")
		}
		fmt.Println("窝草")
	}()

	num := 1
	isErr = &num
	if 1 > 0 {
		num = 2
		isErr = &num
		return
	}

	//打印结果：2017-04-11 12:52:52.794411287 +0800 CST
	//fmt.Println("add", Add(1, 2))
	return
}
