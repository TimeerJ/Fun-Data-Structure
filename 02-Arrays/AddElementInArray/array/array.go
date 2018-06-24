package array

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

