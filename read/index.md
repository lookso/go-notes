# golang细节三百问
## 1.关于nil
```
1.make和new()有什么区别？
2.slice和channel和map 以及interface 使用make初始化,如果没有初始化,只是声明,就是nil,
因为其本身就是引用类型

```

## 2.channel
```
1.channel原理是什么？
2.无缓冲channel和有缓冲channel
3.select-case 、nil
4.closed
5.make创建chan,设置容量值,如果容量超过设置的值,怎么办？

```

## 3.slice和array

```
len和cap、append和copy(),len超过设置的长度,怎么办？
```

## 4.map

```
map 声明和初始化,make和new(),make()长度、
遍历map、
判断指定元素是否在map中、
golang中map不是并发安全、
map 底层结构 hashtable和searchtree、
map元素只读,不能更改其属性,且不能获取其地址
```

## 5.defer和recover
```
```


## 7.go协程
```
并发同步
```


```
值类型和引用类型

https://www.jianshu.com/p/18d3bde7d835
golang传值和传引用
https://blog.csdn.net/cyk2396/article/details/78893828
```

参考资料:

https://gfw.go101.org/article/type-system-overview.html