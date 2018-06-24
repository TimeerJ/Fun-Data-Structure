package main

import (
	"Test/数据结构与算法/玩转数据结构/03-StacksAndQueues/ArrayQueue/queue"
	"fmt"
)

func main() {
	queue := queue.New()
	for i := 0; i < 10; i++ {
		queue.Enqueue(i)
		fmt.Println(queue)
		if i % 3 == 2 {
			queue.Dequeue()
			fmt.Println(queue)
		}
	}
}