package main

import (
	"fmt"
	"Test/数据结构与算法/玩转数据结构/02-Arrays/AddElementInArray/array"
)


func main() {
	arr := array.New()
	arr.Add(0, 1)
	fmt.Printf("%v, %d\n", arr.Data, arr.Size)
	arr.Add(1, 2)
	fmt.Printf("%v, %d\n", arr.Data, arr.Size)
	arr.AddFirst("asdfdsf")
	fmt.Printf("%v, %d\n", arr.Data, arr.Size)
}