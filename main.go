package main

import (
	"fmt"
	li "otus4hom/list"
)

func main() {
	list := &li.List{}
	list.PushBack("Hello")
	list.PushFront(123)
	list.PushBack(3.14)
	list.PushFront(true)

	fmt.Println("Длина списка:", list.Len())

	for item := list.Front(); item != nil; item = item.Next {
		fmt.Println(item.Value)
	}

	// Пример удаления элемента
	itemToRemove := list.Back()
	list.Remove(itemToRemove)
	fmt.Println("Длина списка после удаления:", list.Len())

	// Пример перемещения элемента в начало
	itemToMove := list.Back()
	list.MoveToFront(itemToMove)
	fmt.Println("Первый элемент после перемещения:", list.Front().Value)


	
	cache := li.NewCache(3)
	cache.Set("a", 1)
	cache.Set("b", 2)
	cache.Set('t', 3)
	fmt.Println(cache.Get("a")) // Выведет: 1 true
	cache.Set("d", 4)
	fmt.Println(cache.Get("t")) // Выведет: nil false, так как "b" был вытолкнут из-за размера кэша
}