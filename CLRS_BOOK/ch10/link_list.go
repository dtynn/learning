package ch10

type LinkObject interface {
	Key() int
	Next() LinkObject
	Prev() LinkObject
}

type linkObject struct {
	index int
	sl    *SingleLinkList
}

func (l linkObject) Key() int {
	return l.sl.mgr.key(l.index)
}

func (l linkObject) Next() LinkObject {
	next := l.sl.mgr.next(l.index)
	if next == l.sl.sentinel {
		return nil
	}

	return linkObject{
		index: next,
		sl:    l.sl,
	}
}

func (l linkObject) Prev() LinkObject {
	prev := l.sl.mgr.prev(l.index)
	if prev == l.sl.sentinel {
		return nil
	}

	return linkObject{
		index: prev,
		sl:    l.sl,
	}
}

func NewSingleLinkList() *SingleLinkList {
	mgr := newObjectManager(0)
	sentinel := mgr.allocate()
	mgr.setNext(sentinel, sentinel)
	mgr.setPrev(sentinel, sentinel)

	return &SingleLinkList{
		sentinel: sentinel,
		mgr:      mgr,
	}
}

type SingleLinkList struct {
	sentinel int
	mgr      *objectManager
}

func (s *SingleLinkList) Insert(k int) {
	newo := s.mgr.allocate()

	next := s.mgr.next(s.sentinel)
	s.mgr.setPrev(newo, s.sentinel)
	s.mgr.setNext(newo, next)
	s.mgr.setKey(newo, k)

	s.mgr.setNext(s.sentinel, newo)
	s.mgr.setPrev(next, newo)
}

func (s *SingleLinkList) Search(k int) LinkObject {
	next := s.mgr.next(s.sentinel)
	for next != s.sentinel {
		if s.mgr.key(next) == k {
			return linkObject{
				index: next,
				sl:    s,
			}
		}

		next = s.mgr.next(next)
	}

	return nil
}

func (s *SingleLinkList) Delete(obj LinkObject) {
	if obj == nil {
		return
	}
	lo, ok := obj.(linkObject)
	if !ok {
		return
	}

	prev := s.mgr.prev(lo.index)
	next := s.mgr.next(lo.index)
	s.mgr.setNext(prev, next)
	s.mgr.setPrev(next, prev)

	s.mgr.release(lo.index)
}
