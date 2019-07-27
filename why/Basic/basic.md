## 记录
编译器会将未使⽤用的局部变量当做错误。全局变量即便未使用不会当做错误,可以使用_ 忽略避免报错

```$xslt
str:="abc"
fmt.Println(unsafe.Sizeof(str)) // 获取变量类型大小

```

## 字符串 String

1.string中[]rune和[]byte区别？

```$xslt
package main
func main(){    
    a := "abc你好"
	fmt.Println("a arr",a)   //a arr abc你好
	for index, data := range a {
		fmt.Printf("%d:%c\n", index, data)
	}
	/*
	0:a
    1:b
    2:c
    3:你
    6:好
    */
	s1 := []byte(a)
	fmt.Println("s1 arr",s1)  // s1 arr [97 98 99 228 189 160 229 165 189]
	for index, data := range s1 {
		fmt.Printf("byte:%d:%c\n", index, data)
	}
	/*
	byte:0:a
    byte:1:b
    byte:2:c
    byte:3:ä
    byte:4:½
    byte:5: 
    byte:6:å
    byte:7:¥
    byte:8:½
    */
	b := "abc你好"
	s2 := []rune(b)
	fmt.Println("s2 arr",s2)   //s2 arr [97 98 99 20320 22909]
	for index, data := range s2 {
		fmt.Printf("rune:%d:%c\n", index, data)
	}
	/*
	rune:0:a
    rune:1:b
    rune:2:c
    rune:3:你
    rune:4:好
    */
	fmt.Println("a len", len([]rune(a))) // 5
	fmt.Println("a len", len(a))         // 9
	
	s := "abc"
    fmt.Println("s value",s[0:1]) // 切记获取字符串子串是通过:截取而不是s[0]下标形式,不能获取子串的指针地址
}
```
字符串是不可变值类型，内部⽤用指针指向 UTF-8 字节数组。
可以看到，一个英文字符对应一个byte，一个中文字符对应三个byte。一个rune对应一个UTF-8字符，所以一个中文字符对应一个rune。
Go语言中byte和rune实质上就是uint8和int32类型。
byte用来强调数据是raw data，而不是数字；而rune用来表示Unicode的code point。

字符串是不可变的，在 Go 中字符串是不可变的。字符串一旦被创建就无法改变。
要修改字符串，可先将其转换成 []rune 或 []byte，完成后再转换为 string。⽆论哪种转换，都会重新分配内存，并复制字节数组。


2.for循环
break 可⽤用于 for、switch、select，⽽ continue 仅能⽤用于 for 循环。
switch age:
    case 12:
        fmt.Println(age)  // break可以省略,如果要使用类似continue 跳过当前循环可以使fallthrough
    case 13:
         fmt.Println(age)
         


