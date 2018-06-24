package main

import "fmt"

func main() {
	arr := make([]int, 10)
	for i := range arr {
		arr[i] = i
	}
	fmt.Println(arr)

	scores := []int{100, 99, 66}
	for i := range scores {
		fmt.Println(scores[i])
	}

	for _, v := range scores {
		fmt.Println(v)
	}

	scores[0] = 96

	for i, v := range scores {
		fmt.Printf("%d : %d\n", i+1, v)
	}
}
