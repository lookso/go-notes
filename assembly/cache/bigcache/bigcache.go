package main

import (
	"fmt"
	"time"

	"github.com/allegro/bigcache"
)

func main() {

	cache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))

	cache.Set("my-key", []byte("value"))


	entry, _ := cache.Get("my-key")
	fmt.Println(string(entry))
	hits := cache.Stats().Hits
	fmt.Println(hits)
}
