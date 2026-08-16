package collection

import (
	"github.com/yuvv/candy/function"
)

type MapEntry[K comparable, V any] interface {
	GetKey() K
	GetValue() V
	SetValue(v V) V
	// Equals(o any) bool
}

type DefaultMapEntry[K comparable, V any] struct {
	key   K
	value V
}

func (e *DefaultMapEntry[K, V]) GetKey() K {
	return e.key
}

func (e *DefaultMapEntry[K, V]) GetValue() V {
	return e.value
}

func (e *DefaultMapEntry[K, V]) SetValue(v V) V {
	originV := e.value
	e.value = v
	return originV
}

type Map[K comparable, V any] interface {
	Size() int

	IsEmpty() bool

	ContainsKey(o K) bool
	Get(key K) V
	Put(key K, value V) V
	Remove(key K) V

	PutAll(mp Map[K, V])

	Clear()

	KeySet() Set[K]

	Values() Collection[V]

	EntrySet() Set[MapEntry[K, V]]

	GetOrDefault(key K, dv V) V
	PutIfAbsent(key K, val V) V

	ForEach(consumer function.BiConsumer[K, V])
	ReplaceAll(biFunction function.BiFunction[K, V, V])

	//Stream() stream.Stream[E]
	//
	//ParallelStream() stream.Stream[E]
}

type AbstractMap[K comparable, V any] struct {
	vEqualMethod func(x V, other any) bool
}

func (m *AbstractMap[K, V]) Size() int {
	panic("collection: AbstractMap.Size must be implemented by concrete map")
}

func (m *AbstractMap[K, V]) IsEmpty() bool {
	return m.Size() == 0
}
