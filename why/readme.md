# golang细节三百问
## 1.关于nil
```
1.1.make和new()有什么区别？
1.2.slice和channel和map 以及interface 使用make初始化,如果没有初始化,只是声明,就是nil,
因为其本身就是引用类型

```

## 2.slice和array

```
2.1切片(slice) 原理
2.2 切片的追加,删除和插入,append实现原理,append追加元素到切片时，如果切片定义的容量不够怎么办？
2.3 copy()

```

## 2.chan
```
1.chan原理是什么？
2.无缓冲channel和有缓冲channel
3.select-case 、nil
4.closed
5.make创建chan,设置容量值,如果容量超过设置的值,怎么办？
6.chan底层结构是什么样的？
7.chan同步goruine

```

## 4.map

```
1.map 声明和初始化,make和new(),make()长度、
2.遍历map、
3.判断指定元素是否在map中、
4.golang中map不是并发安全、
5.map 底层结构 hashtable和searchtree、
6.map元素只读,不能更改其属性,且不能获取其地址
```

## 5.error,defer和recover
```
```

## 6.结构体
```
1.结构体转json、
2.结构体字段中json 打tag,omitempty 表示为空则不输出
3.

```

## 7.interface接口
```
1.接口类型推断
2.接口反射reflect.TypeOf()
4.

```

## 7.go协程
```
7.1.实现并发同步有哪些方式？
7.2.协程顺序输出的实现

```

```
## 8. Golang 调度器
```
8.1.如何理解上下文context、内核态和用户态
8.2.go runtime scheduler -gmp


## 9.重点介绍Golang包Sync
```
9.1. golang 如何实现单例模式
9.2. 互斥锁和读写锁

```

## 10.golang 设计模式
```
1.组合模式
2.
```


```
值类型和引用类型

https://www.jianshu.com/p/18d3bde7d835
golang传值和传引用
https://blog.csdn.net/cyk2396/article/details/78893828
```

参考资料:

https://gfw.go101.org/article/type-system-overview.html


golang技术博客
1.CtoLIb码库
https://www.ctolib.com/go
2. no code no life
http://tigerb.cn/

3.IO多路复用与Go网络库的实现
https://ninokop.github.io/2018/02/18/go-net/