package concurrentMap

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"
)

// 10万次的赋值,10万次的读取
var times int = 100000

// 测试ConcurrentMap
func BenchmarkTestConcurrentMap(b *testing.B) {
	for k := 0; k < b.N; k++ {
		b.StopTimer()
		// 产生10000个不重复的键值对(string -> int)
		testKV := map[string]int{}
		for i := 0; i < 10000; i++ {
			testKV[strconv.Itoa(i)] = i
		}

		// 新建一个ConcurrentMap
		pMap := NewConcurrentMap()

		// set到map中
		for k, v := range testKV {
			pMap.Set(k, v)
		}

		// 开始计时
		b.StartTimer()

		wg := sync.WaitGroup{}
		wg.Add(2)

		// 赋值
		go func() {
			// 对随机key,赋值times次
			for i := 0; i < times; i++ {
				index := rand.Intn(times)
				pMap.Set(strconv.Itoa(index), index+1)
			}
			wg.Done()
		}()

		// 读取
		go func() {
			// 对随机key,读取times次
			for i := 0; i < times; i++ {
				index := rand.Intn(times)
				pMap.Get(strconv.Itoa(index))
			}
			wg.Done()
		}()

		// 等待两个协程处理完毕
		wg.Wait()
	}
}

// 测试sync.map
func BenchmarkTestSyncMap(b *testing.B) {
	for k := 0; k < b.N; k++ {
		b.StopTimer()
		// 产生10000个不重复的键值对(string -> int)
		testKV := map[string]int{}
		for i := 0; i < 10000; i++ {
			testKV[strconv.Itoa(i)] = i
		}

		// 新建一个sync.Map
		pMap := &sync.Map{}

		// set到map中
		for k, v := range testKV {
			pMap.Store(k, v)
		}

		// 开始计时
		b.StartTimer()

		wg := sync.WaitGroup{}
		wg.Add(2)

		// 赋值
		go func() {
			// 对随机key,赋值10万次
			for i := 0; i < times; i++ {
				index := rand.Intn(times)
				pMap.Store(strconv.Itoa(index), index+1)
			}
			wg.Done()
		}()

		// 读取
		go func() {
			// 对随机key,读取10万次
			for i := 0; i < times; i++ {
				index := rand.Intn(times)
				pMap.Load(strconv.Itoa(index))
			}
			wg.Done()
		}()

		// 等待两个协程处理完毕
		wg.Wait()
	}
}

//————————————————
//版权声明：本文为CSDN博主「YZF_Kevin」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
//原文链接：https://blog.csdn.net/yzf279533105/article/details/98636679
