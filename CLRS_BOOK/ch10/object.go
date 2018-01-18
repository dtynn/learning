package ch10

const (
	NIL int = -1
)

func newObjectManager(size int) *objectManager {
	if size <= 0 {
		size = 10
	}

	om := &objectManager{
		free:  make([]int, size),
		space: make([]int, size*3),
	}

	for i := 0; i < size; i++ {
		om.free[i] = i * 3
	}

	return om
}

// 分配存储区域, 在必要的时候扩充存储区域
type objectManager struct {
	free  []int
	space []int
}

func (o *objectManager) allocate() int {
	if len(o.free) == 0 {
		o.expand()
	}

	idx := o.free[len(o.free)-1]
	o.free = o.free[:len(o.free)-1]
	return idx
}

func (o *objectManager) expand() {
	size := len(o.space) / 3
	freeSize := len(o.free)

	newfree := make([]int, size+freeSize)
	copy(newfree[size:], o.free)

	// 3 digit for each object
	newspace := make([]int, size*2*3)
	copy(newspace, o.space)

	for i := 0; i < size; i++ {
		newfree[i] = (size + i) * 3
	}

	o.free = newfree
	o.space = newspace
}

func (o *objectManager) release(i int) {
	o.free = append(o.free, i)
}

func (o *objectManager) key(i int) int {
	return o.space[i]
}

func (o *objectManager) next(i int) int {
	return o.space[i+1]
}

func (o *objectManager) prev(i int) int {
	return o.space[i+2]
}

func (o *objectManager) setKey(i, key int) {
	o.space[i] = key
}

func (o *objectManager) setNext(i, next int) {
	o.space[i+1] = next
}

func (o *objectManager) setPrev(i, prev int) {
	o.space[i+2] = prev
}
