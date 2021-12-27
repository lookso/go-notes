package main

import "fmt"

type Res struct {
	uid interface{} `json:"uid"`

}
func main() {
	var res=&Res{uid: 0}
	res.TS()
}
func (res *Res)TS()  {
	fmt.Println(res.uid)
	res.T()
}
func (res *Res)T()  {

	//但是这里可以判断，为什么不同的类型可以判断相等？？？
	fmt.Println(res.uid)
	if res.uid == 0 {
		fmt.Println("相等", res.uid)
	} else {
		fmt.Println("不相等", res.uid)
	}
}
