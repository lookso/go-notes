/*
@Time : 2019/3/23 9:35 PM 
@Author : Tenlu
@File : basic
@Software: GoLand
*/
package main

type student struct {
	Name string
	Age  int
}
func main()  {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	// 错误写法,这样的写法初学者经常会遇到的，很危险!
	// 与Java的foreach一样,都是使用副本的方式。所以m[stu.Name]=&stu实际上一致指向同一个指针,
	// 最终该指针的值为遍历的最后一个struct的值拷贝。 就像想修改切片元素的属性：
	//for _, stu := range stus {
	//	m[stu.Name] = &stu
	//}
	//for k,v:=range m{
	//	println(k,"=>",v.Name)
	//}
	// 也是错误的用法
	//for _, stu := range stus {
	//	stu.Age = stu.Age+10
	//}

	// 结果输出:
	//wang => wang
	//zhou => wang
	//li => wang

	// 正确写法
	for i:=0;i<len(stus);i++  {
		m[stus[i].Name] = &stus[i]
	}

	for k,v:=range m{
		println(k,"=>",v.Name)
	}
	// 正确输出:
	//zhou => zhou
	//li => li
	//wang => wang


}

