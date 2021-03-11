package listNode

import (
	"fmt"
	"testing"
)

func TestLru(t *testing.T) {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	lru.Put(3, 3)
	lru.Put(4, 4)
	lru.Put(5, 5)
	v := lru.Get(4)
	fmt.Println("v", v)
}
