package main

import "fmt"

//思路:两个数相加,计算机中实际就是二进制的相加,分为 0+0 =0,0+1=1,1+1=0 的情况。
//
//这样看来计算分两部:
//1、不考虑进位的情况。其实就是按位异或 
//2、考虑到进位情况,就是与,同时要左移一位。
//3、计算结束返回结果条件,进位为0时,只需要按位异或计算即可

//第一步：两数进行异或操作,相当于求出两个数二进制不算进位的和,记为sum,如:5^7-->101^111=010-->2
//第二步：两数进行与操作并向左移一位,相当于求出两个数二进制进位的和,记为carry，如：(5&7)<<1  -->  (101&111)<<1010  -->  10
//第三步：将sum赋给num1,将carry赋给num2,重复一、二步操作，直到num2等于0,此时说明没有再进位了,返回num1.

// 考察位运算

func main(){
	num1:=11
	num2:=10
	fmt.Println(num1^num2)
	numA:=sumA(123,123)
	fmt.Println("sumA:",numA)

	numB:=sumB(345,345)
	fmt.Println("sumB:",numB)
}
func sumA(num1,num2 int) int{
	var res1 = num1^num2
	var res2 = (num1&num2)<<1
	for res2!=0{
		num1 = res1^res2
		num2 = (res1&res2)<<1
		res1 = num1
		res2 = num2
	}
	return res1
}

// 递归方式
func sumB(num1,num2 int) int{
	if num1 == 0{
		return num2
	}
	if num2 == 0{
		return num1
	}
	var resA = num1 ^ num2           //不进位情况
	var resB = (num1 & num2)<<1      //进位情况，当不为0需要左移一位。为0时，递归结束
	return sumB(resA,resB)
}