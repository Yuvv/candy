package maps

import (
	"github.com/yuvv/candy/collections"
	"github.com/yuvv/candy/sets"
	"github.com/yuvv/candy/util/function"
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

func (e DefaultMapEntry[K, V]) GetKey() K {
	return e.key
}

func (e DefaultMapEntry[K, V]) GetValue() V {
	return e.value
}

func (e DefaultMapEntry[K, V]) SetValue(v V) V {
	originV := e.value
	e.value = v
	return originV
}

type Map[K comparable, V any] interface {
	Size() int

	IsEmpty() bool

	ContainsKey(o comparable) bool
	Get(key K) V
	Put(key K, value V) V
	Remove(key K) V

	PutAll(mp Map[K, V])

	Clear()

	KeySet() sets.Set[K]

	Values() collections.Collection[V]

	EntrySet() sets.Set[MapEntry[K, V]]

	GetOrDefault(key K, dv V) V
	PutIfAbsent(key K, val V) V

	ForEach(consumer function.BiConsumer[K, V])
	ReplaceAll(biFunction function.BiFunction[K, V, V])

	//Stream() stream.Stream[E]
	//
	//ParallelStream() stream.Stream[E]
}

type AbstractMap[K comparable, V any] struct {
}

func (m *AbstractMap[K, V]) Size() int {
	panic("implement me")
}

func (m *AbstractMap[K, V]) IsEmpty() bool {
	return m.Size() == 0
}
