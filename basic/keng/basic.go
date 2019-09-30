package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}
	b := a[:]
	a = append(a, 5)
	b[0] = 0
	fmt.Println(cap(a), cap(b), a, b)
	// 8 4 [1 2 3 4 5] [0 2 3 4]

	fmt.Println(split(5))
	// map
	strMap := map[int]string{1: "a", 2: "b", 3:
	"c"}
	strMap[4] = "d"
	delete(strMap, 0) // map 不能指定索引删除,需要指定key 来删除
	// delete(strMap,1)
	fmt.Printf("%v\n", strMap)

	aa := "012"
	pa := &aa // 0xc000092030
	bb := []byte(*pa) // []byte("012")
	pb := &bb

	aa += "3"
	*pa += "4"
	bb[1] = '5'

	println(*pa) // 01234
	println(string(*pb)) // 052
	fmt.Println("--------keng-------")
	keng()
}

func split(sum int) (x, y, z int) {
	x = sum * 4 / 9 // 因为x 是整形 int 结果是2
	y = sum - x     // 3
	z = x + y       // 5
	return
}

func keng()  {
	i:=1
	s:=[]string{"A","B","C"}
	i,s[i-1]=2,"Z"
	fmt.Printf("s:%s\n",s)

	a := 1
	b := 2
	a, b, a = b, a, b
	c,d:=1,2
	d++
	fmt.Println(a,b,c,d) // 2 1 1 3
}
