package main

import (
	"Test/数据结构与算法/玩转数据结构/02-Arrays/ContainFindAndRemove/array"
	"fmt"
)

func main() {
	arr := array.NewArray(20)
	for i := 0; i < 10; i++ {
		arr.AddFirst(i)
	}
	fmt.Println(arr)

	arr.Add(1, 100)
	fmt.Println(arr)

	arr.AddFirst(-1)
	fmt.Println(arr)

	arr.Remove(2)
	fmt.Println(arr)

	arr.RemoveElement(4)
	fmt.Println(arr)

	arr.RemoveFirst()
	fmt.Println(arr)
}