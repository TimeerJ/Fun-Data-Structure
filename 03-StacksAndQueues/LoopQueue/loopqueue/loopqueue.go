package loopqueue

import (
	"Test/数据结构与算法/玩转数据结构/03-StacksAndQueues/LoopQueue/queue"
	"Test/数据结构与算法/玩转数据结构/03-StacksAndQueues/LoopQueue/array"
)

type LoopQueue struct {
	array.Array
	queue.Queue
	Front, Tail  int
}

func New() *LoopQueue {
	var que *LoopQueue
	que.Array =	*array.New()
	que.Queue = queue.New()
	return que
}

func (l *LoopQueue) GetCap() int {
	return cap(l.Data) - 1
}

func (l *LoopQueue) IsEmpty() bool {
	return l.Front == l.Tail
}

