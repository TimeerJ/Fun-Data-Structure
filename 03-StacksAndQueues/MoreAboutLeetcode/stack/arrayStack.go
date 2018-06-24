package stack

import (
	"Test/数据结构与算法/玩转数据结构/03-StacksAndQueues/ArrayStack/array"
	"fmt"
	"strings"
)

type ArrayStack struct {
	array.Array
}

func NewArrayStack(cap int) *ArrayStack {
	var arr ArrayStack
	arr.Array = *array.NewArray(cap)
	return &arr
}

func New() *ArrayStack {
	var arr ArrayStack
	arr.Array = *array.New()
	return &arr
}

func (arr *ArrayStack) Push(e interface{}) {
	arr.AddLast(e)
}

func (arr *ArrayStack) Pop() interface{} {
	return arr.RemoveLast()
}

func (arr *ArrayStack) Peek() interface{} {
	return arr.GetLast()
}


func (arr *ArrayStack) String() string {
	res := strings.Replace(strings.Trim(fmt.Sprint(arr.Data), "[]"), " ", ", ", -1)
	return fmt.Sprintf("Stack: [%s] top", res)
}