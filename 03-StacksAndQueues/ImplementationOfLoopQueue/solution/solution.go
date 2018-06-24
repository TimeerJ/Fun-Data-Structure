package solution

import (
	"Test/数据结构与算法/玩转数据结构/03-StacksAndQueues/ArrayQueue/array"
	"Test/数据结构与算法/玩转数据结构/03-StacksAndQueues/ArrayQueue/queue"
	"fmt"
)

type TreeNode struct {
	Val    int
	Left   *TreeNode
	Right  *TreeNode
}

func NewTree(x int) *TreeNode {
	var tree TreeNode
	tree.Val = x
	return &tree
}

func LevelOrder(root *TreeNode) *array.Array {
	res := array.New()
	if root == nil {
		return res
	}

	queue := queue.New()
	queue.Enqueue(root)
	fmt.Println(queue.GetSize())

	for !queue.IsEmpty() {
		node := queue.Dequeue()
		level := queue.GetSize()
		fmt.Println(level)
		fmt.Println(res.GetSize())
		if level == res.GetSize() {
			res.AddLast(array.New())
		}


		resval := res.Get(level) // 返回是interface{} []
		resval.(*array.Array).AddLast(node.(*TreeNode).Val)


		if node.(*TreeNode).Left != nil {
			queue.Enqueue(node.(*TreeNode).Left)
		}
		if node.(*TreeNode).Right != nil {
			queue.Enqueue(node.(*TreeNode).Right)
		}

	}
	return res
}
