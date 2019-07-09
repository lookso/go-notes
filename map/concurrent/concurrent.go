package concurrent

import (
	"testing"
	"sync"
	"strconv"
)

// golang 并发不安全
// M
type M struct {
	Map    map[string]string
}

// Set ...
func (m *M) Set(key, value string) {
	m.Map[key] = value
}

// Get ...
func (m *M) Get(key string) string {
	return m.Map[key]
}
// TestMap  ...
func TestMap(t *testing.T) {
	c := M{Map: make(map[string]string)}
	wg := sync.WaitGroup{}
	for i := 0; i < 21; i++ {
		wg.Add(1)
		go func(n int) {
			k, v := strconv.Itoa(n), strconv.Itoa(n)
			c.Set(k, v)
			t.Logf("k=:%v,v:=%v\n", k, c.Get(k))
			wg.Done()
		}(i)
	}
	wg.Wait()
	t.Log("ok finished.")
}