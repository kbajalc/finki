package gol

type HeapMedian struct {
	Length int
	Delay  int

	Index int
	Ready bool
	First int
	Pivot int

	Value int

	Data  []int
	Small *Heap // lower half (extra element for odd)
	Large *Heap // upper half
}

func NewHeapMedian(length int) *HeapMedian {
	if length <= 0 {
		panic("window size must be positive")
	}
	var obj = &HeapMedian{
		Length: length,
		Delay:  length / 2,
		Value:  0,
		Data:   make([]int, length),
		Small:  NewHeap(length/2+2, false),
		Large:  NewHeap(length/2+2, true),
	}
	return obj
}

func (obj *HeapMedian) Reset() {
	obj.Index = 0
	obj.Ready = false
	obj.First = 0
	obj.Pivot = 0

	obj.Value = 0
	for i := 0; i < len(obj.Data); i++ {
		obj.Data[i] = 0
	}
	obj.Small.Reset()
	obj.Large.Reset()
}

func (obj *HeapMedian) Filter(x int) int {
	if obj.Ready {
		obj.Pop(obj.Index)
	}
	obj.Push(obj.Index, x)

	obj.Index++
	if obj.Index == obj.Length {
		obj.Ready = true
		obj.Index = 0
	}
	obj.First = obj.Data[obj.Index]

	var half = obj.Index + obj.Delay
	if half >= obj.Length {
		half -= obj.Length
	}
	obj.Pivot = obj.Data[half]

	var val int
	if obj.Small.Len > obj.Large.Len {
		val = obj.Small.Top()
	} else if obj.Small.Len == obj.Large.Len {
		val = (obj.Small.Top() + obj.Large.Top()) / 2
	} else {
		val = obj.Large.Top()
	}
	obj.Value = val
	return val
}

func (obj *HeapMedian) Pop(head int) {
	var val = obj.Data[head]
	obj.remove(val)
	obj.rebalance()
}

func (obj *HeapMedian) Push(tail int, x int) {
	obj.Data[tail] = x
	if obj.Small.Len == 0 || (obj.Small.Len > 0 && x <= obj.Small.Top()) {
		obj.Small.Push(x)
	} else {
		obj.Large.Push(x)
	}
	obj.rebalance()
}

func (obj *HeapMedian) targetSizes() (wantS, wantL int) {
	var count = obj.Small.Len + obj.Large.Len
	if count == 0 {
		return 0, 0
	}
	if count&1 == 1 {
		return (count + 1) / 2, count / 2
	} else {
		return count / 2, count / 2
	}
}

func (obj *HeapMedian) rebalance() {
	wantS, _ := obj.targetSizes()
	for obj.Small.Len > wantS {
		v := obj.Small.Pop()
		obj.Large.Push(v)
	}
	for obj.Large.Len > 0 && obj.Small.Len < wantS {
		v := obj.Large.Pop()
		obj.Small.Push(v)
	}
}

func (obj *HeapMedian) remove(val int) {
	if val <= obj.Small.Top() {
		if obj.Small.remove(val) {
			return
		}
	}
	if val >= obj.Large.Top() {
		if obj.Large.remove(val) {
			return
		}
	}
}
