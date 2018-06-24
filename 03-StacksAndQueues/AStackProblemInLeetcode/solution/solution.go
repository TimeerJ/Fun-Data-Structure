package solution

import "Test/数据结构与算法/玩转数据结构/03-StacksAndQueues/ArrayStack/stack"

func IsValid(s string) bool {
	stack := stack.New()
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stack.Push(v)
		} else {
			if stack.IsEmpty() {
				return false
			}

			topRune := stack.Pop()
			if v == ')' && topRune != '(' {
				return false
			}
			if v == ']' && topRune != '[' {
				return false
			}
			if v == '}' && topRune != '{' {
				return false
			}
		}
	}
	return stack.IsEmpty()
}
