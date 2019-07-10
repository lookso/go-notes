/*
@Time : 2019-07-10 22:13 
@Author : Tenlu
@File : slce3.go
@Software: GoLand
*/
package main

import "fmt"

func main(){
	var sliceTest []int
	if sliceTest==nil{
		fmt.Printf("未初始化的slice 结果是%v,len是%d\n","nil",len(sliceTest))
	}
	sliceTest=make([]int,1)
	sliceTest[0]=10
	fmt.Println(sliceTest)
	name :=[]int{1,2,3,4}
	fmt.Printf("%p\n",name)


	var osa = make ([]string,0)
	sa:=&osa
	fmt.Printf("addr of osa:%p,\tsa addr:%p \t sa point to:%p \t cap:%v content:%v\n",osa,sa,*sa,cap(*sa),sa)
	for i:=0;i<10;i++{
		*sa=append(*sa,fmt.Sprintf("%v",i))
		fmt.Printf("addr of osa:%p,\tsa addr:%p \t sa point to:%p \t cap:%v content:%v\n",osa,sa,*sa,cap(*sa),sa)
	}
	fmt.Printf("addr of osa:%p,\tsa addr:%p \t sa point to:%p \t cap:%v content:%v\n",osa,sa,*sa,cap(*sa),sa)

	
}
