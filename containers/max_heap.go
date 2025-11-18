package containers

import (
	"errors"
)

// In a binary heap, the keys are stored in an array such that each key is guaranteed to be larger than (or equal to)
// the keys at two other specific positions. In turn, each of those keys must be larger than (or equal to) two
// additional keys, and so forth.

// In a heap, the parent of the node in a position k is in position floor(k/2) and, conversely, the two children
// of the node in position k are in positions 2k and (2k + 1). We can travel up and down by doing simple arithmetic on
// array indices: to move up the tree from a[k] we set k to k/2. To move *down* we set k to 2k or (2k + 1)
//
// We can to develop logarithmic-time insert and remove max implementations

// The heap operations that we consider work by first making a simple modification that could violate the heap
// condition, then traveling through the heap, modifying the heap as required to ensure the heap condition is satisfied
// everywhere. We refer to this process as *reheapifying*, or restoring heap order.
type Heap struct {
	data []int
	n    int // max capacity
	siz  int // current size
}

func NewHeap(n int, data ...int) *Heap {
	if len(data) > n {
		return nil // TODO error
	}

	if len(data) == 0 {
		return &Heap{data: make([]int, n), n: n}
	}
	//We could optimize for n=1 and n=2 and skip the 'down' process, but it's an early optimization

	l := len(data)
	for i := l/2 - 1; i >= 0; i-- {
		down(data, i, l)
	}

	return &Heap{data: data, n: n, siz: l}
}

// Add insert a new element into the heap
func (h *Heap) Add(e int) error {
	if h.siz == h.n {
		return errHeapFull
	}

	iat := h.siz
	h.data[iat] = e
	h.siz++

	up(h.data, iat)
	return nil
}

var errHeapEmpty = errors.New("empty heap")
var errHeapFull = errors.New("full heap")

func (h *Heap) Peek() (int, error) {
	if h.siz == 0 {
		return 0, errHeapEmpty
	}

	return h.data[0], nil
}

func (h *Heap) Get() (int, error) {
	if h.siz == 0 {
		return 0, errHeapEmpty
	}

	root := h.data[0] //get the root
	last := h.siz - 1
	h.siz--
	if h.siz > 0 {
		h.data[0] = h.data[last]
		down(h.data, 0, h.siz)
	}

	return root, nil
}

// Size returns the number of elements
func (h *Heap) Size() int {
	return h.siz
}

// a node has been removed from position i
// We want to assure that the greater-than relationship still holds
func down(data []int, i int, siz int) {
	for {
		l := 2*i + 1
		if l >= siz {
			break
		}

		j := l
		// r is the (i+1)-th child of the ith node
		if r := l + 1; r < siz && data[r] > data[l] {
			j = r
		}

		if data[i] >= data[j] {
			break
		}

		data[i], data[j] = data[j], data[i]

		// Continue working from the last swapped node
		i = j
	}
}

// A new element is inserted, we need to check wether it is greater than it's parent

func up(data []int, i int) {
	for i > 0 {
		p := (i - 1) / 2
		if data[p] >= data[i] {
			return
		}
		data[p], data[i] = data[i], data[p]
		i = p
	}
}
