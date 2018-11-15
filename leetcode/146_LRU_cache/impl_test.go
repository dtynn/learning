package impl

import "testing"

func TestLRUCache(t *testing.T) {
	cache := Constructor(2)

	cache.Put(1, 1)
	cache.Put(2, 2)

	if cache.Get(1) != 1 {
		t.Fatal("expect 1 for 1")
	}

	cache.Put(3, 3)

	if cache.Get(2) != -1 {
		t.Fatal("expected evicted key 2")
	}

	cache.Put(4, 4)

	if cache.Get(1) != -1 {
		t.Fatal("expected evicted key 1")
	}

	if cache.Get(3) != 3 {
		t.Fatal("expect 3 for 3")
	}

	if cache.Get(4) != 4 {
		t.Fatal("expect 4 for 4")
	}
}

func TestLRUCache2(t *testing.T) {
	cache := Constructor(3)

	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	cache.Put(4, 4)

	cache.Get(4)
	cache.Get(3)
	cache.Get(2)
	cache.Get(1)

	cache.Put(5, 5)

	cache.Get(1)
	cache.Get(2)
	cache.Get(3)
	cache.Get(4)
	cache.Get(5)
}
