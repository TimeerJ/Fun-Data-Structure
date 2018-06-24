package main

import (
	"Test/数据结构与算法/玩转数据结构/03-StacksAndQueues/ArrayStack/stack"
	"fmt"
)

func main() {
	stack := stack.New()
	for i := 0; i < 5; i++ {
		stack.Push(i)
		fmt.Println(stack)
	}

	stack.Pop()
	fmt.Println(stack)
}
