package main

import (
	"fmt"
	"unicode"
)

var stuList = []int{13, 12, 22, 21, 44}

func GetWxUserByUserID(id int) []int {
	var arr = make([]int, 0)
	if id == 1 {
		arr = []int{11, 21, 31, 41}
	}
	if id == 2 {
		arr = []int{12, 22, 32, 42}
	}
	if id == 3 {
		arr = []int{13, 23, 33, 43}
	}
	if id == 4 {
		arr = []int{14, 24, 34, 44}
	}
	return arr
}

func main() {
	var wcList = []int{1, 2, 3, 4}
	for _, v := range wcList {
		userFollowList := GetWxUserByUserID(v)
		for _, vv := range userFollowList {
			for _, sv := range stuList {
				if sv == vv {
					fmt.Println(sv)
				}
			}
		}
	}

	//var stuBindList = make(map[string][]string, 0)
	//for _, v := range workCodeList.List {
	//	userFollowList, err := component.GetWxUserByUserID(ctx, v.WxUserID)
	//	if err != nil {
	//		logger.Ex(ctx, tag, "GetWxUserByUserID is err:%+v, req:%+v", err, v)
	//		continue
	//	}
	//	tag := false
	//	for _, fv := range userFollowList {
	//		for _, stu := range req.StuInfo {
	//			if fv.ExternalUserID == stu.ExternalUserId {
	//				stuBindList[stu.ExternalUserId] = append(stuBindList[stu.ExternalUserId], fv.ExternalUserID)
	//				tag = true
	//			}
	//		}
	//		//判断学员是否加该老师多个微信号
	//		// 一个学员加了该老师全部微信的好友,调用该员工最新添加学员微信的企微号并执行发送
	//		if tag {
	//			if len(stuBindList[fv.ExternalUserID]) == len(userFollowList) {
	//				if maxId == fv.ID {
	//					wxUserid = v.WxUserID
	//					wxNickName = v.NickName
	//				}
	//			}
	//		}
	//	}
	//}

	////fmt.Println(numTest())
	//
	//// 1=>[1,2,3,4,5],2=>[1,2,3,4,5],3=>[1,2,3,4,5]
	//// [2,3,4,5,8]
	//
	//var arr = [][]int{{1, 2, 3, 4, 5}, {10, 2, 3, 9, 5}, {1, 2, 3, 7, 9,8}}
	//var info = []int{2, 3, 4, 5, 8}
	//var newArr = make(map[int][]int, 0)
	//for ik, iv := range info {
	//
	//	for k, v := range arr {
	//		tag := 0
	//		for _, vv := range v {
	//			if iv == vv {
	//				tag = 1
	//			}
	//		}
	//		if tag == 1 {
	//			newArr[ik] = append(newArr[ik], k)
	//		}
	//	}
	//}
	//fmt.Println(newArr)
	//println("-----------")

	//var a = make(map[string]interface{}, 0)
	//a["a"] = 123
	//var aa = a["a"]
	//fmt.Println(aa.(int) - 12)
	//
	//fmt.Println(reflect.TypeOf(a["a"]))
	//time.Now().Unix()
	//str := "中国 牛逼"
	//if Sp(str) {
	//	fmt.Println("ok")
	//} else {
	//	fmt.Println("false")
	//}
}

func Sp(str string) bool {

	for _, v := range str {
		if unicode.IsSpace(v) {
			return true
		}
	}
	return false
}

func numTest() []int {
	var arr = []int{1, 2, 3, 4, 5, 6}
	var newArr = make([]int, 0)
	var count int
	for _, v := range arr {
		if v == 3 {

			continue
		}
		newArr = append(newArr, v)
		count++
	}
	return newArr
}
