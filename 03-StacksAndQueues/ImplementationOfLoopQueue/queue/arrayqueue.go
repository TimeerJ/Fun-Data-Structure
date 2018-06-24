package queue

import (
	"Test/数据结构与算法/玩转数据结构/03-StacksAndQueues/ArrayQueue/array"
	"strings"
	"fmt"
)

type ArrayQueue struct {
	array.Array
}

func New() *ArrayQueue{
	var arr ArrayQueue
	arr.Array = *array.New()
	return &arr
}

func (arr *ArrayQueue)Enqueue(e interface{}) {
	arr.AddLast(e)
}
func (arr *ArrayQueue) Dequeue() interface{} {
	return arr.RemoveFirst()
}

func (arr *ArrayQueue) GetFront() interface{} {
	return arr.GetFirst()
}

func (arr *ArrayQueue) String() string {
	res := strings.Replace(strings.Trim(fmt.Sprint(arr.Data), "[]"), " ", ", ", -1)
	return fmt.Sprintf("Queue: front [%s] tail", res)
}
