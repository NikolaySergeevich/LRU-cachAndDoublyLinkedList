package main

import (
	"fmt"
	li "otus4hom/list"
)

func main() {
	cache := li.NewCache(3)
	cache.Set("a", 1)
	cache.Set("b", 2)
	cache.Set("t", 3)
	fmt.Println(cache.Get("a")) // Выведет: 1 true
	cache.Set("d", 4)
	for i := cache.Queue.Front(); i != nil; i = i.Next {
		fmt.Println(i.Value)
	}
	fmt.Println(cache.Get("b")) // Выведет: nil false, так как "b" был вытолкнут из-за размера кэша
	for i := cache.Queue.Front(); i != nil; i = i.Next {
		fmt.Println(i.Value)
	}
}