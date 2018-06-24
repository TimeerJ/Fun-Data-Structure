package array

import (
	"fmt"
)

type data interface{}

type Array struct {
	Data    []data
	Size	int
}

// 无参工厂函数,默认Array的容量cap=10
func New() *Array {
	return &Array{
		Data: make([]data, 0, 10),
		Size: 0,
	}
}

// 有参工厂函数，传入Array的容量cap构造Array
func NewArray(cap int) *Array {
	if cap < 0 {
		return New()
	}
	arr := &Array{ make([]data, 0, cap), 0 }
	return arr
}

// 获取Array的容量
func (a *Array) getCap() int {
	return len(a.Data)
}

// 获取Array的元素个数
func (a *Array) getSize() int {
	return a.Size
}

// 返回Array是否为空
func (a *Array) isEmpyt() bool {
	return a.Size == 0
}

// 向后所有元素添加一个新元素
func (a *Array) AddLast(e data) {
	a.Data = append(a.Data, e)
	a.Size ++
}

// 在所有元素前添加一个元素
func (a *Array) AddFirst(e data) {
	a.Add(0, e)
}

//在index索引的位置插入一个新元素e
func (a *Array) Add(index int, e data) {
	if index < 0 || index > a.Size {
		panic("Add failed. Require index >= 0 and index <= size.")
	}
	a.Data = append(a.Data[:index], append([]data{e}, a.Data[index:]...)...)
	a.Data[index] = e
	a.Size++
}

// 获取 index 索引位置的元素
func (a *Array) Get(index int) interface{} {
	if index < 0 || index >= a.Size {
		panic("Get failed. Index is illegal.")
	}

	return a.Data[index]
}

// 修改index 索引位置的元素为e
func (a *Array) Set(index int, e data) {
	if index < 0 || index >= a.Size {
		panic("Set failed. Index is illegal.")
	}

	a.Data[index] = e
}

// 查找数组中是否有元素e
func (a *Array) Contains(e data) bool {
	for _, v := range a.Data {
		if v == e {
			return true
		}
	}
	return  false
}

// 查找Array中元素e所在的索引,如果不存在元素e,则返回-1
func (a *Array) Find(e data) int {
	for i, v := range a.Data {
		if v == e {
			return i
		}
	}
	return -1
}

// 从Array中删除index位置的元素,返回删除的元素
func (a *Array) Remove(index int) interface{} {
	if index < 0 || index >= a.Size {
		panic("Remove failed. Index is illegal.")
	}
	ret := a.Data[index]
	a.Data = append(a.Data[:index],a.Data[index+1:]...)
	a.Size --

	if a.Size == cap(a.Data)/2 {
		a.resize(cap(a.Data)/2)
	}
	return ret
}

// 从Array中删除第一个元素,返回删除的元素
func (a *Array) RemoveFirst() interface{} { return a.Remove(0) }

// 从Array中删除最后一个元素,返回删除的元素
func (a *Array) RemoveLast() interface{} { return a.Remove(a.Size - 1) }

// 从Array中删除元素e
func (a *Array) RemoveElement(e data) {
	index := a.Find(e)
	if index != -1 {
		a.Remove(index)
	}
}

type Stringer interface {
	String() string
}

func (a *Array) String() string {
	e := fmt.Sprint(a.Data)

	return fmt.Sprintf("Array: size = %d, capacity = %d\n%s", a.Size, cap(a.Data), e)
}

func (a *Array) resize(newCap int) {
	newArr := make([]data, 0, newCap)
	a.Data = append(newArr, a.Data...)
}