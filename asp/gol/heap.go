package gol

// Fixed-capacity heap.
type Heap struct {
	Cap int   // maximum number of elements
	Min bool  // max-heap if false, min-heap if true
	Len int   // current number of elements
	Buf []int // backing array, length = capacity
}

func NewHeap(capy int, mini bool) *Heap {
	if capy <= 0 {
		panic("Heap capacity must be > 0")
	}
	return &Heap{
		Cap: capy,
		Min: mini,
		Len: 0,
		Buf: make([]int, capy),
	}
}

func (h *Heap) Reset() {
	h.Len = 0
	for i := range h.Buf {
		h.Buf[i] = 0
	}
}

func (h *Heap) Top() int {
	if h.Len == 0 {
		panic("Heap.Top on empty heap")
	}
	return h.Buf[0]
}

func (h *Heap) Push(x int) {
	if h.Len == len(h.Buf) {
		panic("Heap overflow")
	}
	h.Buf[h.Len] = x
	h.up(h.Len)
	h.Len++
}

func (h *Heap) Pop() int {
	if h.Len == 0 {
		panic("Heap.Pop on empty heap")
	}
	top := h.Buf[0]
	h.Len--
	h.swap(0, h.Len)
	h.down(0)
	return top
}

func (h *Heap) cmp(i, j int) bool {
	ai, aj := h.Buf[i], h.Buf[j]
	if h.Min {
		return ai < aj
	} else {
		return ai > aj
	}
}

func (h *Heap) swap(i, j int) {
	h.Buf[i], h.Buf[j] = h.Buf[j], h.Buf[i]
}

func (h *Heap) up(i int) {
	for {
		if i == 0 {
			return
		}
		p := (i - 1) >> 1
		if !h.cmp(i, p) {
			return
		}
		h.swap(i, p)
		i = p
	}
}

func (h *Heap) down(i int) {
	for {
		l := i*2 + 1
		if l >= h.Len {
			return
		}
		j := l
		r := l + 1
		if r < h.Len && h.cmp(r, l) {
			j = r
		}
		if !h.cmp(j, i) {
			return
		}
		h.swap(i, j)
		i = j
	}
}

// Remove the deepest occurrence of val from the heap.
func (h *Heap) remove(val int) bool {
	for i := h.Len - 1; i >= 0; i-- {
		if h.Buf[i] == val {
			_ = h.removeAt(i)
			return true
		}
	}
	return false
}

// Remove the element at index i from the heap.
func (h *Heap) removeAt(i int) int {
	if i == 0 {
		return h.Pop()
	}
	if i >= h.Len {
		panic("invalid index")
	}
	r := h.Buf[i]
	h.Len--
	h.swap(i, h.Len)
	// Try both up and down since the swapped element
	// may violate heap property in either direction.
	h.up(i)
	h.down(i)
	return r
}
