package main

import (
	"Test/数据结构与算法/玩转数据结构/02-Arrays/GenericDataStructures/array"
	"fmt"
)

type Student struct {
	Name   string
	Score  int
}

type Stringer interface {
	String() string
}

func New(studentName string, studentScore int) *Student {
	return &Student{
		Name: studentName,
		Score: studentScore,
	}
}

func (s *Student) String() string {
	return fmt.Sprintf("Student(name: %s, score: %d)", s.Name, s.Score)
}

func main() {
	arr := new(array.Array)
	arr.AddLast(New("Alice", 100))
	arr.AddLast(New("Bob", 66))
	arr.AddLast(New("Charlie", 88))
	fmt.Println(arr)
}




