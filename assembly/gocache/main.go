package main

import (
	"github.com/lifei6671/gocache"
	"log"
	"time"
)

func main() {
	// 使用内置缓存
	gocache.AddWithCacheItem(gocache.CacheItem{Key: "cache", Value: "cache-1", RemovedCallback: func(key string, value interface{}, reason gocache.CacheEntryRemovedReason) {
		log.Println(key, value, reason)
	}})
	gocache.Add("cache", "cache-2")

	log.Println(gocache.Count())
	log.Println(gocache.Get("cache"))

	// 自定义缓存
	m := gocache.NewMemoryCache(time.Second * 2)

	m.AddWithCacheItem(gocache.CacheItem{Key: "cache", Value: "cache-1", RemovedCallback: func(key string, value interface{}, reason gocache.CacheEntryRemovedReason) {
		log.Println(key, value, reason)
	}})
	m.Add("cache", "cache-2")

	log.Println(m.Count())
	log.Println(m.Get("cache"))

}
