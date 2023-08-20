package main

import (
	"fmt"

	lru "github.com/hashicorp/golang-lru/v2"
)

func main() {
	cache, err := lru.New[string, any](10)
	if err != nil {
		panic(err)
	}
	cache.Add("test_1", nil)
	cache.Add("test_2", nil)
	cache.Add("test_3", nil)
	cache.Add("test_1", nil)
	cache.Add("test_2", nil)

	if key, _, ok := cache.RemoveOldest(); ok {
		fmt.Println(key)
	} else {
		fmt.Println(cache.Len())
	}
	if key, _, ok := cache.RemoveOldest(); ok {
		fmt.Println(key)
	} else {
		fmt.Println(cache.Len())
	}
	if key, _, ok := cache.RemoveOldest(); ok {
		fmt.Println(key)
	} else {
		fmt.Println(cache.Len())
	}

	cache, err = lru.NewWithEvict(2, func(key string, value any) {
		fmt.Println("overflow, remove " + key)
	})
	if err != nil {
		panic(err)
	}
	cache.Add("test_1", nil)
	cache.Add("test_2", nil)
	cache.Add("test_3", nil)
	cache.Add("test_1", nil)
	cache.Add("test_2", nil)

	if key, _, ok := cache.RemoveOldest(); ok {
		fmt.Println(key)
	} else {
		fmt.Println(cache.Len())
	}
	if key, _, ok := cache.RemoveOldest(); ok {
		fmt.Println(key)
	} else {
		fmt.Println(cache.Len())
	}
	if key, _, ok := cache.RemoveOldest(); ok {
		fmt.Println(key)
	} else {
		fmt.Println(cache.Len())
	}
}
