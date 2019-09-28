package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	// 第一种
	rand.Seed(time.Now().UnixNano()) // 种下初始值,随机种子
	num:=rand.Int()
	fmt.Println(num)
	fmt.Println(rand.Intn(100))
	// 第二种
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println(r.Int())
	// 范围
	fmt.Println(r.Intn(100)) // 返回[0,100]之间的一个随机整数
	fmt.Println(r.Perm(10)) // 返回一个有n个元素的，[0,n)范围内整数的伪随机排列的切片。

	//返回一个取值范围在[0.0, 1.0)的伪随机float32值。
	fmt.Println(r.Float32())

	fmt.Println(getPerm(10,r))


}
func getPerm(n int,r *rand.Rand) []int {
	m := make([]int, n)
	// In the following loop, the iteration when i=0 always swaps m[0] with m[0].
	// A change to remove this useless iteration is to assign 1 to i in the init
	// statement. But Perm also effects r. Making this change will affect
	// the final state of r. So this change can't be made for compatibility
	// reasons for Go 1.
	for i := 0; i < n; i++ {
		j := r.Intn(i + 1)
		m[i] = m[j]
		m[j] = i
	}
	return m
}