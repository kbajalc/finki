package gol

// MEDIAN FILTER ////////////////////////////////

type SortMedian struct {
	Delay  int
	Length int
	Strict bool

	Sum    int
	Last   int
	Value  int
	First  int
	Pivot  int
	Index  int
	Ready  bool
	Buffer []int
	Sorted []int
}

func NewSortMedian(length int) *SortMedian {
	obj := &SortMedian{
		Delay:  length / 2,
		Length: length,
		Strict: true,
		Buffer: make([]int, length),
		Sorted: make([]int, length),
	}
	for i := 0; i < length; i++ {
		obj.Buffer[i] = 0
		obj.Sorted[i] = 0
	}
	return obj
}

func (obj *SortMedian) Reset() {
	obj.Sum = 0
	obj.Last = 0
	obj.Value = 0
	obj.First = 0
	obj.Pivot = 0
	obj.Index = 0
	obj.Ready = false
	for i := 0; i < obj.Length; i++ {
		obj.Buffer[i] = 0
		obj.Sorted[i] = 0
	}
}

func (obj *SortMedian) Filter(datum int) int {
	if !obj.Ready && obj.Index == 0 {
		for i := 0; i < obj.Length; i++ {
			obj.Buffer[i] = datum
			obj.Sorted[i] = datum
			obj.Sum += datum
		}
	}

	var rem = obj.Buffer[obj.Index]
	obj.Buffer[obj.Index] = datum
	obj.Index++
	if obj.Index == obj.Length {
		obj.Index = 0
		obj.Ready = true
	}
	obj.First = obj.Buffer[obj.Index]

	obj.Sum += datum - rem
	obj.Last = datum

	var half = obj.Index + obj.Delay
	if half >= obj.Length {
		half -= obj.Length
	}
	obj.Pivot = obj.Buffer[half]

	var ix = 0
	for i := 0; i < obj.Length; i++ {
		if obj.Sorted[i] == rem {
			ix = i
			break
		}
	}

	obj.Sorted[ix] = datum
	for i := ix; i < obj.Length-1; i++ {
		if obj.Sorted[i] > obj.Sorted[i+1] {
			obj.Sorted[i], obj.Sorted[i+1] = obj.Sorted[i+1], obj.Sorted[i]
		} else {
			break
		}
	}
	for i := ix; i > 0; i-- {
		if obj.Sorted[i] < obj.Sorted[i-1] {
			obj.Sorted[i], obj.Sorted[i-1] = obj.Sorted[i-1], obj.Sorted[i]
		} else {
			break
		}
	}

	var m = obj.Delay
	if obj.Strict {
		obj.Value = obj.Sorted[m]
	} else {
		obj.Value = obj.Sorted[m-1] + obj.Sorted[m] + obj.Sorted[m+1]/3
	}

	return obj.Value
}
