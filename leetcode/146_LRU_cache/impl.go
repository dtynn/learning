package impl

type cacheItem struct {
	key  int
	val  int
	next *cacheItem
	prev *cacheItem
}

// LRUCache struct
type LRUCache struct {
	capacity int
	values   map[int]*cacheItem
	first    *cacheItem
	last     *cacheItem
}

// Constructor returns a LRUCache instance
func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		values:   map[int]*cacheItem{},
	}
}

// Get returns the value
func (l *LRUCache) Get(key int) int {
	item, ok := l.values[key]
	if !ok {
		return -1
	}

	l.setRecent(item)

	return item.val
}

// Put inserts a value
func (l *LRUCache) Put(key int, value int) {
	if item, ok := l.values[key]; ok {
		item.val = value
		l.setRecent(item)
		return
	}

	item := &cacheItem{
		key:  key,
		val:  value,
		next: l.first,
	}

	if l.first != nil {
		l.first.prev = item
	}

	l.first = item
	l.values[key] = item

	if l.last == nil {
		l.last = item
	}

	if len(l.values) > l.capacity {
		remove := l.last
		remove.prev.next = nil
		l.last = remove.prev
		delete(l.values, remove.key)
	}
}

func (l *LRUCache) setRecent(item *cacheItem) {
	if len(l.values) > 1 {
		// is last
		if l.last == item {
			l.last = item.prev
		}

		// not first
		if l.first != item {
			item.prev.next = item.next
			if item.next != nil {
				item.next.prev = item.prev
			}

			item.prev = nil
			item.next = l.first
			l.first.prev = item
			l.first = item
		}

	}
}
