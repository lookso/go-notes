package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/cast"
	"math"
	"sort"
	"strings"
	"time"
)

var Names struct {
	NewStr string `json:"new_str"`
	Name   string `json:"name"`
}

func Tf() (page int ,err error) {
	page=1
	for{
		tag:=0
		for _,vv:=range []int{1,2,3,4,5,6}{
			if vv==5{
				tag=1
				break
			}
			fmt.Println("vv",vv)
		}
		if tag==1{
			break
		}
	}
	fmt.Println("page",page)
	return page,nil
}
func main() {
	time.Sleep(time.Second*2)
	fmt.Println(11)
	time.Sleep(time.Second*2)
	fmt.Println(22)
	if 1>10 && 1<10{
		fmt.Println("ok")
	}
	return
	fmt.Println(Tf())
	return

//	var s = `{"msg_content":[{"content_type":1,"content":"🔥咱们人文美育体验课已经结束了，您还可以0元免费领取3节国际象棋课程❗
//
//👉世界第一大智力运动，助力宝贝锻炼思维能力、提升宝贝的专注力
//👉变化莫测的棋局，磨练孩子的抗挫力和意志力，改掉拖拉马虎的习惯
//👉全国冠军教宝贝get国际象棋新技能，为孩子多规划一条升学路
//
//点击链接即可报名：
//https://www.xueersi.com/wx50/635554139?source=635554139u0026site_id=8316u0026adsite_id=16584648","url":"","image_url":"","title":"","desc":""}]}`
//

	var s=`"{\"msg_content\":[{\"content_type\":1,\"content\":\"🔥咱们人文美育体验课已经结束了，您还可以0元免费领取3节国际象棋课程❗\\n\n👉世界第一大智力运动，助力宝贝锻炼思维能力、提升宝贝的专注力\n👉变化莫测的棋局，磨练孩子的抗挫力和意志力，改掉拖拉马虎的习惯\n👉全国冠军教宝贝get国际象棋新技能，为孩子多规划一条升学路\n\n点击链接即可报名：\nhttps://www.xueersi.com/wx50/635554139?source=635554139u0026site_id=8316u0026adsite_id=16584648\",\"url\":\"\",\"image_url\":\"\",\"title\":\"\",\"desc\":\"\"}]}"`


	//	var s = `🔥咱们人文美
	//育体验课已经结束了，您还可以0元免费领取3节国际象棋课程❗`
	fmt.Println(strings.Replace(s, "\\n", "\\\\n", -1))
	//narr := strings.Split(s, "\n")
	//fmt.Println(narr)
	//fmt.Println(len(narr))
	//newc := ""
	//for _, v := range narr {
	//	newc += v + "\\n"
	//}
	//newc=strings.Trim(newc,"\\n")
	//fmt.Println(newc)

	return

	fmt.Println(strings.Replace(s, "\n", "\\\\n", -1))

	fmt.Println(bytes.Compare([]byte("a"), []byte("b")))

	arr := []int{1, 2, 3, 4, 5, 6, 8, 9, 10}
	d := sort.Search(len(arr), func(i int) bool { return arr[i] >= 70 })
	fmt.Println(d)
	return

	classIDsArr := []int{}
	num := 100

	fmt.Println(cast.ToFloat64(len(classIDsArr)) / cast.ToFloat64(num))
	fmt.Println(cast.ToFloat64(len(classIDsArr) / num))
	reqNum := cast.ToInt(math.Ceil(cast.ToFloat64(len(classIDsArr)) / cast.ToFloat64(num)))
	fmt.Println("reqNum", reqNum)
	for i := 1; i <= reqNum; i++ {
		dl := i * num
		if i > len(classIDsArr)/num {
			dl = len(classIDsArr)
		}
		fmt.Println(i, dl, len(classIDsArr)/num)

		fmt.Println(classIDsArr[(i-1)*num : dl])
	}

	return
	fmt.Println(cast.ToInt(math.Ceil(201 / 200)))
	return
	for i := 0; i < 100; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i)
	}
	return
	a := 1
	if a != 1 && a != 20 {
		fmt.Println("a-111", a)
	}

	for i := 0; i < 10; i++ {
		if i == 6 {
			return
		}
		fmt.Println(i)

	}
	type ErrInfo struct {
		err error
	}
	var errInfo *ErrInfo
	defer func() {
		//获取traceId
		msgByte := []byte("同步成功-")
		if errInfo != nil && errInfo.err != nil {
			msgByte, _ = json.Marshal(errInfo.err.Error())
		}
		fmt.Println(string(msgByte))
	}()
	var err error
	err = MyErr()
	if err != nil {
		errInfo = &ErrInfo{
			err: err,
		}
		//	fmt.Println(err)
		return
	}
}
func MyErr() error {
	return nil
	err := errors.New("wc")
	return err
}
