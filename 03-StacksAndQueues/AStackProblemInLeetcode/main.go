package main

import "fmt"

type data interface{}

type array struct {
	arr  []data
}

func (a *array) Push(e data) {
	a.arr = append(a.arr, e)
}

func (a *array) Pop() interface{} {
	top := a.arr[len(a.arr)-1]
	a.arr = a.arr[:len(a.arr)-1]
	return top
}

func (a *array) IsEmpty() bool {
	return len(a.arr) == 0
}

func isValid(s string) bool {
	stack := new(array)
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stack.Push(v)
		} else {
			if stack.IsEmpty() {
				return false
			}

			topdata := stack.Pop()
			if v== ')' && topdata != '(' {
				return false
			}
			if v == ']' && topdata != '[' {
				return false
			}
			if v == '}' && topdata != '{' {
				return false
			}
		}
	}
	return stack.IsEmpty()
}

func main() {
	a := isValid("()()()")
	fmt.Println(a)
}