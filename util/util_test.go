/*
@Time : 2020-01-01 19:29 
@Author : Tenlu
@File : util_test
@Software: GoLand
*/
package util

import (
	"testing"
)

// 单元测试,功能测试
func TestSum(t *testing.T) {
	Sum(20, 30)
	t.Log("sum(20, 30) exec success")
}

// 当需要终止当前测试用例时，可以使用 FailNow
func TestFailNow(t *testing.T) {
	t.FailNow()
}

// 只标记错误不终止测试的方法
func TestFail(t *testing.T) {
	t.Log("before fail")
	t.Fail()
	t.Log("after fail")
}

func TestDivision(t *testing.T) {
	if i, e := Division(6, 0); i != 3 || e != nil { //try a unit test on function
		t.Error("除法函数测试没通过") // 如果不是如预期的那么就报错
	} else {
		t.Log("第一个测试通过了") //记录一些你期望记录的信息
	}
	if i, e := Division(6, 2); i != 3 || e != nil { //try a unit test on function
		t.Error("除法函数测试没通过") // 如果不是如预期的那么就报错
	} else {
		t.Log("第一个测试通过了") //记录一些你期望记录的信息
	}
}

// go test -v util_test.go util.go

// 参数-run对应一个正则表达式，只有测试函数名被它正确匹配的测试函数才会被go test测试命令运行
// go test -v -run="TestSum"

// 测试覆盖率
// go test 命令接受 -covermode 标志将覆盖模式设置为三种设置之一：
// set: 每个语句是否执行？
// count: 每个语句执行了几次？
// atomic: 类似于 count, 但表示的是并行程序中的精确计数

// go test -coverprofile=size_coverage.out

// go test -covermode=count  -coverprofile=size_coverage.out

// go tool cover -html=size_coverage.out

/**
单元测试框架提供的日志方法
Log	打印日志，同时结束测试
Logf	格式化打印日志，同时结束测试
Error	打印错误日志，同时结束测试
Errorf	格式化打印错误日志，同时结束测试
Fatal	打印致命日志，同时结束测试
Fatalf	格式化打印致命日志，同时结束测试
**/


// go 单元测试
// http://objcoding.com/2018/09/14/go-test/