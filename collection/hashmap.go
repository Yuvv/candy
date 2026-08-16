package collection

import (
	"github.com/yuvv/candy/function"
)

type HashMap[K comparable, V any] struct {
	AbstractMap[K, V]

	hMap map[K]V
}

func (h *HashMap[K, V]) Size() int {
	return len(h.hMap)
}

func (h *HashMap[K, V]) ContainsKey(o K) bool {
	_, ok := h.hMap[o]
	return ok
}

func (h *HashMap[K, V]) Get(key K) V {
	return h.hMap[key]
}

func (h *HashMap[K, V]) Put(key K, value V) V {
	v := h.hMap[key]
	h.hMap[key] = value
	return v
}

func (h *HashMap[K, V]) Remove(key K) V {
	v := h.hMap[key]
	delete(h.hMap, key)
	return v
}

func (h *HashMap[K, V]) PutAll(mp Map[K, V]) {
	it := mp.EntrySet().Iterator()
	for it.HasNext() {
		entry := it.Next()
		h.hMap[entry.GetKey()] = entry.GetValue()
	}
}

func (h *HashMap[K, V]) Clear() {
	for k := range h.hMap {
		delete(h.hMap, k)
	}
}

func (h *HashMap[K, V]) KeySet() Set[K] {
	set := NewHashSetWithCap[K](len(h.hMap))
	for k := range h.hMap {
		set.Add(k)
	}
	return set
}

func (h *HashMap[K, V]) Values() Collection[V] {
	list := NewSpecArrayListWithCap[V](h.vEqualMethod, h.Size())
	for _, v := range h.hMap {
		list.Add(v)
	}
	return list
}

func (h *HashMap[K, V]) EntrySet() Set[*DefaultMapEntry[K, V]] {
	set := NewHashSetWithCap[*DefaultMapEntry[K, V]](h.Size())
	for k, v := range h.hMap {
		set.Add(&DefaultMapEntry[K, V]{k, v})
	}
	return set
}

func (h *HashMap[K, V]) GetOrDefault(key K, dv V) V {
	v, ok := h.hMap[key]
	if ok {
		return v
	}
	return dv
}

func (h *HashMap[K, V]) PutIfAbsent(key K, val V) V {
	v, ok := h.hMap[key]
	if ok {
		return v
	}
	h.hMap[key] = val
	return val
}

func (h *HashMap[K, V]) ForEach(consumer function.BiConsumer[K, V]) {
	for k, v := range h.hMap {
		consumer(k, v)
	}
}

func (h *HashMap[K, V]) ReplaceAll(biFunction function.BiFunction[K, V, V]) {
	for k, v := range h.hMap {
		h.hMap[k] = biFunction(k, v)
	}
}

// ---------------------------------------------------------------------------

// NewHashMap return a hashmap
func NewHashMap[K comparable, V comparable]() *HashMap[K, V] {
	return &HashMap[K, V]{
		AbstractMap: AbstractMap[K, V]{
			vEqualMethod: func(x V, other any) bool {
				return x == other
			},
		},
		hMap: map[K]V{},
	}
}

// NewHashMapWithMap initialize a hashmap with a specific aMap
func NewHashMapWithMap[K comparable, V comparable](aMap Map[K, V]) *HashMap[K, V] {
	hm := NewHashMap[K, V]()
	hm.PutAll(aMap)
	return hm
}

// NewSpecHashMap return a hashmap
func NewSpecHashMap[K comparable, V any](vem func(x V, other any) bool) *HashMap[K, V] {
	if vem == nil {
		panic("param vem cannot be nil")
	}
	return &HashMap[K, V]{
		AbstractMap: AbstractMap[K, V]{
			vEqualMethod: vem,
		},
		hMap: map[K]V{},
	}
}

// NewSpecHashMapWithMap initialize a hashmap with a specific aMap
func NewSpecHashMapWithMap[K comparable, V any](vem func(x V, other any) bool, aMap Map[K, V]) *HashMap[K, V] {
	hm := NewSpecHashMap[K, V](vem)
	hm.PutAll(aMap)
	return hm
}
