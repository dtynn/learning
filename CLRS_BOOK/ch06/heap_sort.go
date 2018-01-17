package ch06

type MaxHeap struct {
}

type heapable interface {
	Len() int
	Swap(i, j int)
	Val(i int) int
}

type heapableInts []int

func (h heapableInts) Len() int {
	return len(h)
}

func (h heapableInts) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h heapableInts) Val(i int) int {
	return h[i]
}

func (m MaxHeap) Sort(in []int) []int {
	m.buildMaxHeap(heapableInts(in))
	heap := heapableInts(in[:])

	for len(heap) > 1 {
		heap[0], heap[len(heap)-1] = heap[len(heap)-1], heap[0]
		heap = heap[:len(heap)-1]
		m.maxHeapify(heap, 0)
	}

	return in
}

func (m MaxHeap) left(i int) int {
	return i<<1 + 1
}

func (m MaxHeap) right(i int) int {
	return i<<1 + 2
}

func (m MaxHeap) parent(i int) int {
	return (i - 1) >> 1
}

func (m MaxHeap) maxHeapify(heap heapable, i int) {
	largest := i

	size := heap.Len()

	if left := m.left(i); left < size && heap.Val(left) > heap.Val(largest) {
		largest = left
	}

	if right := m.right(i); right < size && heap.Val(right) > heap.Val(largest) {
		largest = right
	}

	if largest != i {
		heap.Swap(i, largest)
		m.maxHeapify(heap, largest)
	}
}

func (m MaxHeap) buildMaxHeap(heap heapable) {
	for i := heap.Len()/2 - 1; i >= 0; i-- {
		m.maxHeapify(heap, i)
	}
}
