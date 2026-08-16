package cache

import "testing"

func TestLRUPutGetUpdateAndLen(t *testing.T) {
	c := NewLRU[string, int](2)
	c.Put("a", 1)
	c.Put("b", 2)
	c.Put("a", 3)

	if got, ok := c.Get("a"); !ok || got != 3 {
		t.Fatalf("Get(a) = %d,%v", got, ok)
	}
	if c.Len() != 2 {
		t.Fatalf("Len = %d", c.Len())
	}
}

func TestLRUEvictsLeastRecentlyUsed(t *testing.T) {
	c := NewLRU[string, int](2)
	c.Put("a", 1)
	c.Put("b", 2)
	if _, ok := c.Get("a"); !ok {
		t.Fatalf("expected a to be present")
	}
	c.Put("c", 3)

	if _, ok := c.Get("b"); ok {
		t.Fatalf("b should be evicted")
	}
	if got, ok := c.Get("a"); !ok || got != 1 {
		t.Fatalf("a should remain, got %d,%v", got, ok)
	}
	if got, ok := c.Get("c"); !ok || got != 3 {
		t.Fatalf("c should be present, got %d,%v", got, ok)
	}
}

func TestLRURemoveClearAndMissing(t *testing.T) {
	c := NewLRU[string, int](2)
	c.Put("a", 1)
	if !c.Remove("a") {
		t.Fatalf("Remove existing key should return true")
	}
	if c.Remove("missing") {
		t.Fatalf("Remove missing key should return false")
	}
	c.Put("b", 2)
	c.Clear()
	if c.Len() != 0 {
		t.Fatalf("Len after Clear = %d", c.Len())
	}
	if _, ok := c.Get("b"); ok {
		t.Fatalf("b should be absent after Clear")
	}
}

func TestNewLRUPanicsForInvalidCapacity(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatalf("NewLRU should panic for invalid capacity")
		}
	}()
	_ = NewLRU[string, int](0)
}
