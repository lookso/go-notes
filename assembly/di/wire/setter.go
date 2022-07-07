package wire

import "github.com/google/wire"

// Age 年龄
type Age int

// Name 姓名
type Name string

// Person 演示值注入的工具人
type Person struct {
	Age
	Name
}

// ProvideAge 注入年龄
func ProvideAge() Age {
	return 18
}

// ProvideName 注入姓名
func ProvideName() Name {
	return "Tom"
}

// Set 值注入配置
// wire.Struct来指定那些字段要被注入到结构体中，
// 如果是全部字段，也可以简写成 "*"
var Set = wire.NewSet(
	ProvideAge,
	ProvideName,
	wire.Struct(new(Person), "Age", "Name"),
)