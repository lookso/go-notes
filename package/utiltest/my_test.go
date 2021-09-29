package utiltest

import (
	"flag"
	"github.com/spf13/cast"
	"os"
	"testing"
)

/**
go test -v -run TestArgs -json -args "cloud" 1 2
-json
-json 参数用于指示go test将结果输出转换成json格式，以方便自动化测试解析使用。

-o
-o 参数指定生成的二进制可执行程序，并执行测试，测试结束不会删除该程序。

-count=1
禁止从缓存中获取数据

**/

// TestMain 用于主动执行各种测试，可以测试前后做setup和tear-down操作
func TestMain(m *testing.M) {
	println("TestMain setup.")

	retCode := m.Run() // 执行测试，包括单元测试、性能测试和示例测试

	println("TestMain tear-down.")

	os.Exit(retCode)
}

func TestArgs(t *testing.T) {
	if !flag.Parsed() {
		flag.Parse()
	}

	argList := flag.Args() // flag.Args() 返回 -args 后面的所有参数，以切片表示，每个元素代表一个参数
	t.Logf("argList：%+v", argList)
	if len(argList) < 3 {
		t.Fatal("arge not enough") // fatal 错误级别会中断程序运行
	}
	sum := Add(cast.ToInt(argList[1]), cast.ToInt(argList[2]))
	t.Logf("sum:%v", sum)

	for _, arg := range argList {
		if arg == "cloud" {
			t.Log("Running in cloud.")
		} else {
			t.Log("Running in other mode.")
		}
	}
	t.SkipNow()
}
