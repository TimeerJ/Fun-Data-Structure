package main

type ListNode struct {
	Val     int
	Next	*ListNode
}

func removeElements(head *ListNode, val int) *ListNode {

	for head != nil && head.Val == val {
		delNode := head
		head = head.Next
		delNode.Next = nil
	}

	if head == nil {
		return head
	}

	prev := head
	for prev.Next != nil {
		if prev.Next.Val == val {
			delNode := prev.Next
			prev.Next = delNode.Next
			delNode.Next = nil
		}

		prev = prev.Next
	}

	return head
}
