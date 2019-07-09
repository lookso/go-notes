
# 关于nil

```
1.slice和channel和map 以及interface 使用make初始化,如果没有初始化,只是声明,就是nil,因为其本身就是引用类型
func main() {
	var theMap map[string]string
    if theMap == nil {
        fmt.Println("map未初始化前其值是nil")
    }
    var theChannel chan int
    if theChannel == nil {
        fmt.Println("channel未初始化前其值是nil")
    }
    var theInterface interface{}
    if theInterface == nil {
        fmt.Println("interface未初始化前其值是nil")
    }
    var theSlice []int
    if theSlice == nil {
        fmt.Println("slice未初始化前其值是nil")
    }
    if theSlice == nil {
        fmt.Println("ok")
    }
	//if theSlice == theMap {
	//	fmt.Println("ok") // 报错,两个类型属性值不一样,两个不同类型的nil值可能不能相互比较
	//}
}
我们可以使用前置解引用操作符*来访问存储在一个指针所表示的地址处的值（即此指针所引用着的值）。
比如，对于基类型为T的指针类型的一个指针值p，我们可以用*p来表示地址p处的值。
此值的类型为T。*p称为指针p的解引用。解引用是取地址的逆过程
解引用一个nil指针将产生一个恐慌。

我们可以用内置函数new来为任何类型的值开辟一块内存并将此内存块的起始地址做为此值的地址返回。
假设T是任一类型，则函数调用new(T)返回一个类型为*T的指针值。 存储在返回指针值所表示的地址处的值（可被看作是一个匿名变量）为T的零值。
我们也可以使用前置取地址操作符&来获取一个可寻址的值的地址。
对于一个类型为T的可寻址的值t，我们可以用&t来取得它的地址。&t的类型为*T。

// 返回指针类型的数据
func newInt() *int {
    intPoint:=new(int)
	fmt.Println("类型指针地址:",intPoint)

	a := 3
	fmt.Println(a)
	var ip *int /* 声明指针变量 */
	ip=&a
	fmt.Println(ip)
	return &a
}

```

