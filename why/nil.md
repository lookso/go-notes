
# 1.关于nil、值类型和引用类型

## 1.1. 值类型和引用类型:
值类型：
所有像int、float、bool和string这些类型都属于值类型，使用这些类型的变量直接指向存在内存中的值，值类型的变量的值存储在栈中。
当使用等号=将一个变量的值赋给另一个变量时，如 j = i ,实际上是在内存中将 i 的值进行了拷贝。可以通过 &i 获取变量 i 的内存地址

引用类型：
引用类型就是说保存的只是内存的地址，而不是具体的值，效率在大数据读取方面效率会高很多。
复杂的数据通常会需要使用多个字，这些数据一般使用引用类型保存。
一个引用类型的变量r1存储的是r1的值所在的内存地址（数字），或内存地址中第一个字所在的位置，这个内存地址被称之为指针，这个指针实际上也被存在另外的某一个字中。

值类型复制例子:
```
func main() {
    yourName:="jackma"
    myName:=yourName
    fmt.Println("yourName point add:",&yourName," myName point add:",&myName)
    fmt.Printf("yourName point add:%p\n",&yourName)

    var yourInfo = make([]string,50)
    //yourInfo=[]string{"name","age","sex"}
    yourInfo[0]="name"

    yourInfoA:=yourInfo[0:2]
    myInfo:=yourInfo // //其实就是浅拷贝,a和b同时指向同一个地址0xc00006c180 通过yourInfo和myInfo都可以改变值，myInfo值变了yourInfo的值也变了
    myInfo[0]="name_jack"
    fmt.Printf("yourInfo value:%v\n",yourInfo)
    fmt.Printf("yourInfo point add:%p,myInfo point add:%p,yourInfoA point add:%p\n",yourInfo,myInfo,yourInfoA)

    heInfo:=append(yourInfo,"hieght")
    fmt.Printf("heInfo point add:%p\n",heInfo)
    var uInfo=make([]string,30)
    copy(uInfo,yourInfo)
    fmt.Printf("uInfo point add:%p\n",uInfo)
}
```
输出可以看到值类型的两个变量内存地址并不一样
yourName point add: 0xc00000e200  myName point add: 0xc00000e210
yourName point add:0xc00000e200
以引用类型slice为例:引用类型的slice内存地址是一样的
yourInfo value:[name_jack age sex]
yourInfo point add:0xc000066180,myInfo point add:0xc000066180,yourInfoA point add:0xc000066180

heInfo point add:0xc000077500
uInfo point add:0xc00007e000

同时可以发现
apend和copy函数可以实现对slice的深拷贝,细节见slice章节


指针的零值,不是0，而是nil。任何未初始化的指针值都为nil。

slice和channel和map 以及interface使用make初始化,如果没有初始化,只是声明,就是nil,因为其本身就是引用类型
还有func函数和pointer指针未初始化前也是nil值

Go语言中的引用类型有：映射（map）,切片（slice）,通道（channel）,方法与函数。

func main() {
    // 指针声明
    var x = 100
    // 指针初始化,通过对x做&运算符来获取其地址，然后将该地址分配给指针p
    var p *int = &x

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
```
2.##


## 3.指针（地址）解引用
我们可以使用前置解引用操作符*来访问存储在一个指针所表示的地址处的值（即此指针所引用着的值）。
比如，对于基类型为T的指针类型的一个指针值p，我们可以用*p来表示地址p处的值。
此值的类型为T。*p称为指针p的解引用。解引用是取地址的逆过程
解引用一个nil指针将产生一个恐慌。

我们可以用内置函数new来为任何类型的值开辟一块内存并将此内存块的起始地址做为此值的地址返回。
假设T是任一类型，则函数调用new(T)返回一个类型为*T的指针值。 存储在返回指针值所表示的地址处的值（可被看作是一个匿名变量）为T的零值。
我们也可以使用前置取地址操作符&来获取一个可寻址的值的地址。
对于一个类型为T的可寻址的值t，我们可以用&t来取得它的地址。&t的类型为*T。

// 返回指针类型的数据

```
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

## 3. Go语言中指针限制
在Go中，指针是不能参与算术运算的。比如，对于一个指针p， 运算p++和p-2都是非法的

