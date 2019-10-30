package main

import (
	"fmt"
	"sort"
	"strconv"
	"sync"
)

var s []int

func appendValue(i int) {
	s = append(s, i)
}

func main() {
	for i := 0; i < 10000; i++ { //10000个协程同时添加切片
		go appendValue(i)
	}

	for i, v := range s { //同时打印索引和值
		fmt.Println(i, ":", v)
	}
	var m = map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e", 6: "f"}
	sortMap(m)
	uu:=UserAges{ages: map[string]int{}}
	for i:=1;i<100;i++ {
		uu.Add("jack"+strconv.Itoa(i),i)
		num:=uu.Get("jack"+strconv.Itoa(i))
		fmt.Println(num)
	}

}

func sortMap(m map[int]string) {
	type kv struct {
		Key   int
		Value string
	}
	var ss []kv
	for k, v := range m {
		ss = append(ss, kv{Key: k, Value: v})
	}
	sort.Slice(ss, func(i, j int) bool {
		//return ss[i].Value > ss[j].Value // 降序
		// return ss[i].Value < ss[j].Value // 升序
		return ss[i].Key < ss[j].Key // 升序
	})
	for _, kv := range ss {
		fmt.Printf("%d, %s\n", kv.Key, kv.Value)
	}
}

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	//ua.Lock()
	//defer ua.Unlock()
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}
