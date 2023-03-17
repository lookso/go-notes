package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

func UnicodeIndex(str, substr string) int {
	// 子串在字符串的字节位置
	result := strings.Index(str, substr)
	if result >= 0 {
		// 获得子串之前的字符串并转换成[]byte
		prefix := []byte(str)[0:result]
		// 将子串之前的字符串转换成[]rune
		rs := []rune(string(prefix))
		// 获得子串之前的字符串的长度，便是子串在字符串的字符位置
		result = len(rs)
	}

	return result
}
/*
if err != nil {
		for i := 1; i <= 2; i++ {
			var retryErr error
			if strings.Contains(err.Error(), "接口并发调用超过限制") {
				time.Sleep(200 * time.Millisecond)
				resp, retryErr = cli.SingleRemark(ctx, req)
				if retryErr != nil {
					return resp, retryErr
				}
				return resp, nil
			}
			if retryErr == nil {
				break
			}
		}
		return resp, err
	}
*/
func errFun() error {
	var err=errors.New("hello")
	for i := 1; i <= 2; i++ {
		fmt.Println("i",i)
		//var retryErr error
		if strings.Contains(err.Error(), "hello") {
			time.Sleep(200 * time.Millisecond)
			err =errors.New("hello")
			if err != nil {
				return err
			}

			return  nil
		}
		if err == nil {
			break
		}
	}
	return nil

}



func main() {
	removeCh:=make(chan int64)
	go func() {
		for {
			//time.Sleep(100*time.Millisecond)
			removeCh<-time.Now().Unix()
		}
	}()
	for {
		select {
		case taskID := <-removeCh:
			fmt.Println("remove taskid:%+v", taskID)
		default:
			time.Sleep(10*time.Second)
			fmt.Println(time.Now().Unix())
		}
	}

	var content="大家好推荐你快购买快过年了"
	start:=strings.Index(content,"推荐")
	end:=strings.Index(content,"购买")
	fmt.Println(start,end)
	if start>=0 && end>0 {
		fmt.Println(content[0:start] + content[end+6:len(content)])
	}

	return
	fmt.Println(errFun())
	return

	path := "/zhgjm/chenwei.htm"
	p := strings.LastIndex(path, "/")
	p2 := strings.LastIndex(path, ".")
	fmt.Println("p", p)
	fmt.Println("p2", p2)
	fmt.Println(path[p:p2])
	fmt.Println(strings.Replace(path, path[p:p2], path[p:p2]+"_2", -1))

	category := "周公解梦人物篇"
	lastpos := strings.Index(category, "篇")
	fmt.Println(lastpos)
	pos := strings.Index(category, "梦")
	s := strings.Replace(category, category, category[pos+3:lastpos], -1)
	fmt.Println(pos, s)
	return
	type ArrStu struct {
		ID int `json:"id"`
	}
	var arr = make([]ArrStu, 0, 10)
	arr = append(arr, ArrStu{
		ID: 10,
	})
	arr = append(arr, ArrStu{
		ID: 11,
	})
	arr = append(arr, ArrStu{
		ID: 12,
	})

	for i := 0; i < len(arr); i++ {
		for j := 0; j < 3; j++ {
			if j == 1 {
				//	fmt.Println("j-j",j,i)
				continue
			}
			//	fmt.Println("j",j,i)
			//	fmt.Println(arr[i].ID)
			fmt.Println(arr[i].ID, j)
		}
	}

	fmt.Println("          /\\         ")
	fmt.Println("         /  \\        ")
	fmt.Println("        /    \\       ")
	fmt.Println("       /      \\      ")
	fmt.Println("      /        \\     ")
	fmt.Println("     /          \\    ")
	fmt.Println("    /            \\   ")
	fmt.Println("   /              \\  ")
	fmt.Println("  /                \\ ")
	fmt.Println(" /                  \\")
	fmt.Println("---------------------")
}
