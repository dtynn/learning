package ch06

import (
	"fmt"
)

type Element struct {
	key int
	ele interface{}
}

type queueHeap []Element

func (q queueHeap) Len() int {
	return len(q)
}

func (q queueHeap) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q queueHeap) Val(i int) int {
	return q[i].key
}

type MaxPriorityQueue struct {
	queueHeap
	max MaxHeap
}

func (m *MaxPriorityQueue) Insert(ele interface{}) {
	for i := range m.queueHeap {
		if m.queueHeap[i].ele == ele {
			return
		}
	}

	m.queueHeap = append(m.queueHeap, Element{
		ele: ele,
	})

	return
}

func (m *MaxPriorityQueue) Maximum() interface{} {
	if m.queueHeap.Len() == 0 {
		return nil
	}

	return m.queueHeap[0].ele
}

func (m *MaxPriorityQueue) ExtractMax() interface{} {
	size := m.queueHeap.Len()
	if size == 0 {
		return nil
	}

	if size > 1 {
		m.queueHeap[0], m.queueHeap[size-1] = m.queueHeap[size-1], m.queueHeap[0]
	}

	ele := m.queueHeap[size-1].ele
	m.queueHeap = m.queueHeap[:size-1]
	if m.queueHeap.Len() > 1 {
		m.maxHeapify(0)
	}

	return ele
}

func (m *MaxPriorityQueue) IncreaseKey(ele interface{}, k int) error {
	shouldRebuild := false

	for i := range m.queueHeap {
		if m.queueHeap[i].ele != ele {
			continue
		}

		if m.queueHeap[i].key > k {
			return fmt.Errorf("current key %v is greater than k %v", m.queueHeap[i].key, k)
		}

		shouldRebuild = m.queueHeap[i].key != k
		m.queueHeap[i].key = k
	}

	if shouldRebuild {
		m.max.buildMaxHeap(m.queueHeap)
	}

	return nil
}

func (m *MaxPriorityQueue) maxHeapify(i int) {
	m.max.maxHeapify(m.queueHeap, i)
}
