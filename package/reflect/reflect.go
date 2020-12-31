package main

import (
	"fmt"
	"reflect"
)

func init()  {
	fmt.Println("1111")
}
func init()  {
	fmt.Println("222")
}
type NextDate struct {
	Year int `json:"year" id:"100"`
	Moon int `json:"moon"`
	Day  int `json:"day"`
}

//反射是指在程序运行期对程序本身进行访问和修改的能力，程序在编译时变量被转换为内存地址，变量名不会被编译器写入到可执行部分，在运行程序时程序无法获取自身的信息。
//支持反射的语言可以在程序编译期将变量的反射信息，如字段名称、类型信息、结构体信息等整合到可执行文件中，并给程序提供接口访问反射信息，这样就可以在程序运行期获取类型的反射信息，并且有能力修改它们。
func main() {
	var nd = NextDate{
		Year: 2021,
		Moon: 9,
	}
	typeOfNd := reflect.TypeOf(nd)
	fmt.Println(typeOfNd.Name(), typeOfNd.Kind())
	valueOfNd := reflect.ValueOf(nd)
	fmt.Println("reflct-value:", valueOfNd.Kind())

	// 遍历结构体所有成员
	for i := 0; i < typeOfNd.NumField(); i++ {
		// 获取每个成员的结构体字段类型
		fieldType := typeOfNd.Field(i)
		// 输出成员名和tag
		fmt.Printf("name: %v  tag: '%v'\n", fieldType.Name, fieldType.Tag)
	}
	// 通过字段名, 找到字段类型信息
	if catType, ok := typeOfNd.FieldByName("Year"); ok {
		// 从tag中取出需要的tag
		fmt.Println(catType.Tag.Get("json"), catType.Tag.Get("id"))
	}
	var a int
	typeOfA := reflect.TypeOf(a)
	fmt.Println(typeOfA.Name(), typeOfA.Kind())

	// 零值判断
	var b *int
	fmt.Println("var b *int:", reflect.ValueOf(b).IsNil())

	// 声明整型变量a并赋初值
	var c int = 1024
	// 获取变量a的反射值对象(a的地址)
	valueOfC := reflect.ValueOf(&c)
	// 取出c地址的元素(c的值)
	valueOfC = valueOfC.Elem()
	if valueOfC.CanAddr() {
		fmt.Println("c", valueOfC.Kind())
	}
	fmt.Println(valueOfC.CanSet())
	// 修改c的值为1
	valueOfC.SetInt(1)
	// 打印c的值
	fmt.Println(valueOfC.Int())
	// 首先保证对象是可写,可被寻址的才能被修改
	// 反射对象 f 是不可写的，但是我们也不想修改 f，事实上我们要修改的是 *f。为了得到 f 指向的数据，可以调用 Value 类型的 Elem 方法。Elem 方法能够对指针进行“解引用”，然后将结果存储到反射 Value 类型对象 f 中：
	var f float64 = 3.1415
	v := reflect.ValueOf(&f) // 传了f的指针，&f 隐式地被转成了interface{}
	v = v.Elem() // 必须使用其元素
	v.SetFloat(2.873) // 成功修改
	fmt.Println(v.Float())
	fmt.Println(reflect.DeepEqual(1,"1"))
}
