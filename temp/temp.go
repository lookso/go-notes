package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/cast"
	"math"
	"sort"
	"time"
)

func main() {
	fmt.Println(bytes.Compare([]byte("a"), []byte("b")))

	arr := []int{1,2,3,4,5,6,8,9,10}
	d := sort.Search(len(arr), func(i int) bool { return arr[i]>=70})
	fmt.Println(d)
	return

	classIDsArr := []int{}
	num := 100

	fmt.Println(cast.ToFloat64(len(classIDsArr)) / cast.ToFloat64(num))
	fmt.Println(cast.ToFloat64(len(classIDsArr) / num))
	reqNum := cast.ToInt(math.Ceil(cast.ToFloat64(len(classIDsArr)) / cast.ToFloat64(num)))
	fmt.Println("reqNum",reqNum)
	for i := 1; i <= reqNum; i++ {
		dl := i*num
		if i> len(classIDsArr)/num {
			dl = len(classIDsArr)
		}
		fmt.Println(i,dl,len(classIDsArr)/num)

		fmt.Println(classIDsArr[(i-1)*num: dl])
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
		//发送知音楼报警
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
