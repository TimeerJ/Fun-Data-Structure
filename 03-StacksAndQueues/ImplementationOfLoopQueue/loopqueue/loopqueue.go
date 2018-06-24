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

func (l *LoopQueue) GetCapacity() int {
	return cap(l.Data) - 1
}



func (l *LoopQueue) IsEmpty() bool {
	return l.Front == l.Tail
}

func (l *LoopQueue) EnQueue(e interface{}) {
	if (l.Tail+1)%len(l.Data) == l.Front {
		l.resize(l.GetCapacity() * 2)
	}
	l.Data[l.Tail] = e
	l.Tail = (l.Tail + 1) % len(l.Data)
	l.Size ++
}

func (l *LoopQueue) GetFront() interface{} {
	if l.IsEmpty() {
		panic("Queue is empty.")
	}
	return l.Data[l.Front]
}

func (l *LoopQueue)resize(newCapacity int) {
	var arr array.Array
	for i := 0; i < l.Size; i++ {
		arr.Data[i] = l.Data[(i+l.Front) % len(l.Data)]
	}
	l.Data = arr.Data
	l.Front = 0
	l.Tail = l.Size
}


