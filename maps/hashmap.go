package maps

import (
	"github.com/yuvv/candy/collections"
	"github.com/yuvv/candy/function"
	"github.com/yuvv/candy/lists"
	"github.com/yuvv/candy/sets"
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

func (h *HashMap[K, V]) KeySet() sets.Set[K] {
	set := sets.NewHashSetWithCap[K](len(h.hMap))
	for k := range h.hMap {
		set.Add(k)
	}
	return set
}

func (h *HashMap[K, V]) Values() collections.Collection[V] {
	list := lists.NewArrayListWithCap[V](h.Size())
	for _, v := range h.hMap {
		list.Add(v)
	}
	return list
}

func (h *HashMap[K, V]) EntrySet() sets.Set[MapEntry[K, V]] {
	set := sets.NewHashSetWithCap[MapEntry[K, V]](h.Size())
	for k, v := range h.hMap {
		set.Add(DefaultMapEntry[K, V]{k, v})
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
func NewHashMap[K comparable, V any]() *HashMap[K, V] {
	return &HashMap[K, V]{
		hMap: map[K]V{},
	}
}

// NewHashMapWithMap initialize a hashmap with a specific aMap
func NewHashMapWithMap[K comparable, V any](aMap Map[K, V]) *HashMap[K, V] {
	hm := &HashMap[K, V]{
		hMap: map[K]V{},
	}
	hm.PutAll(aMap)
	return hm
}
