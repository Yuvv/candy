package collection

import (
	"reflect"

	"github.com/yuvv/candy/function"
	"github.com/yuvv/candy/iter"
)

type HashSet[E comparable] struct {
	AbstractCollection[E]

	set map[E]bool
}

type hashsetIter struct {
	// todo:
}

func (h *HashSet[E]) Iterator() iter.Iterator[E] {
	//TODO implement me
	panic("implement me")
}

func (h *HashSet[E]) ForEach(action function.Consumer[E]) {
	for k := range h.set {
		action(k)
	}
}

func (h *HashSet[E]) Spliterator() iter.Spliterator[E] {
	//TODO implement me
	panic("implement me")
}

func (h *HashSet[E]) Size() int {
	return len(h.set)
}

func (h *HashSet[E]) Contains(o E) bool {
	oType := reflect.TypeOf(o)
	if !oType.Comparable() {
		return false
	}
	_, ok := h.set[o]
	return ok
}

func (h *HashSet[E]) ToArray() []E {
	res := make([]E, 0, len(h.set))
	for k := range h.set {
		res = append(res, k)
	}
	return res
}

// Add Returns true if this set did not already contain the specified element
func (h *HashSet[E]) Add(e E) bool {
	_, ok := h.set[e]
	h.set[e] = true
	return !ok
}

// Remove Returns true if the set contained the specified element
func (h *HashSet[E]) Remove(o E) bool {
	_, ok := h.set[o]
	if ok {
		delete(h.set, o)
		return true
	}
	return false
}

func (h *HashSet[E]) ContainsAll(collection Collection[E]) bool {
	it := collection.Iterator()
	for it.HasNext() {
		_, ok := h.set[it.Next()]
		if !ok {
			return false
		}
	}
	return true
}

func (h *HashSet[E]) AddAll(collection Collection[E]) bool {
	modified := false
	it := collection.Iterator()
	for it.HasNext() {
		h.set[it.Next()] = true
		modified = true
	}
	return modified
}

func (h *HashSet[E]) RemoveAll(collection Collection[E]) bool {
	modified := false
	it := collection.Iterator()
	for it.HasNext() {
		ele := it.Next()
		_, ok := h.set[ele]
		if ok {
			delete(h.set, ele)
			modified = true
		}
	}
	return modified
}

func (h *HashSet[E]) RemoveIf(predicate function.Predicate[E]) bool {
	modified := false
	for k := range h.set {
		if predicate(k) {
			delete(h.set, k)
			modified = true
		}
	}
	return modified
}

func (h *HashSet[E]) RetainAll(collection Collection[E]) bool {
	modified := false
	for k := range h.set {
		if !collection.Contains(k) {
			delete(h.set, k)
			modified = true
		}
	}
	return modified
}

func (h *HashSet[E]) Clear() {
	for k := range h.set {
		delete(h.set, k)
	}
}

/// ---------------------------------------------------------------------------

// NewHashSet return a new hashset
func NewHashSet[E comparable]() *HashSet[E] {
	return NewHashSetWithCap[E](4)
}

// NewHashSetWithCap return a new hashset with expected capacity
func NewHashSetWithCap[E comparable](cap int) *HashSet[E] {
	return &HashSet[E]{
		set: make(map[E]bool, cap),
	}
}
