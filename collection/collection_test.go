package collection

import (
	"testing"

	"github.com/yuvv/candy/function"
	"github.com/yuvv/candy/iter"
)

type defaultMethodCollection struct {
	AbstractCollection[int]
	values []int
}

func newDefaultMethodCollection(values ...int) *defaultMethodCollection {
	collection := &defaultMethodCollection{values: append([]int(nil), values...)}
	collection.GetEleEqualMethod = func() func(int, any) bool {
		return func(value int, other any) bool {
			otherValue, ok := other.(int)
			return ok && value == otherValue
		}
	}
	collection.iteratorMethod = collection.Iterator
	collection.sizeMethod = collection.Size
	collection.addMethod = collection.Add
	return collection
}

func (c *defaultMethodCollection) Iterator() iter.Iterator[int] {
	return &defaultMethodIterator{collection: c, last: -1}
}

func (c *defaultMethodCollection) Size() int {
	return len(c.values)
}

func (c *defaultMethodCollection) Add(value int) bool {
	for _, existing := range c.values {
		if existing == value {
			return false
		}
	}
	c.values = append(c.values, value)
	return true
}

type defaultMethodIterator struct {
	collection *defaultMethodCollection
	next       int
	last       int
}

func (i *defaultMethodIterator) HasNext() bool {
	return i.next < len(i.collection.values)
}

func (i *defaultMethodIterator) Next() int {
	value := i.collection.values[i.next]
	i.last = i.next
	i.next++
	return value
}

func (i *defaultMethodIterator) Remove() {
	if i.last < 0 {
		panic("Remove called before Next")
	}
	copy(i.collection.values[i.last:], i.collection.values[i.last+1:])
	i.collection.values = i.collection.values[:len(i.collection.values)-1]
	i.next = i.last
	i.last = -1
}

func (i *defaultMethodIterator) ForEachRemaining(action function.Consumer[int]) {
	for i.HasNext() {
		action(i.Next())
	}
}

func TestAbstractCollectionRemoveReturnsWhetherValueWasFound(t *testing.T) {
	collection := newDefaultMethodCollection(1, 2, 3)

	if collection.AbstractCollection.Remove(4) {
		t.Fatal("Remove returned true for missing value")
	}
	if !collection.AbstractCollection.Remove(2) {
		t.Fatal("Remove returned false for existing value")
	}
	if len(collection.values) != 2 || collection.values[0] != 1 || collection.values[1] != 3 {
		t.Fatalf("values after Remove = %v, want [1 3]", collection.values)
	}
}

func TestAbstractCollectionRemoveIfRemovesAllMatches(t *testing.T) {
	collection := newDefaultMethodCollection(1, 2, 3, 4)

	if !collection.AbstractCollection.RemoveIf(func(value int) bool { return value%2 == 0 }) {
		t.Fatal("RemoveIf returned false after removing matches")
	}
	if len(collection.values) != 2 || collection.values[0] != 1 || collection.values[1] != 3 {
		t.Fatalf("values after RemoveIf = %v, want [1 3]", collection.values)
	}
	if collection.AbstractCollection.RemoveIf(func(value int) bool { return value > 10 }) {
		t.Fatal("RemoveIf returned true when no value matched")
	}
}

func TestAbstractCollectionAddAllAccumulatesModifications(t *testing.T) {
	collection := newDefaultMethodCollection(1)
	source := newDefaultMethodCollection(2, 1)

	if !collection.AbstractCollection.AddAll(source) {
		t.Fatal("AddAll returned false when an earlier add modified the collection")
	}
	if len(collection.values) != 2 || collection.values[0] != 1 || collection.values[1] != 2 {
		t.Fatalf("values after AddAll = %v, want [1 2]", collection.values)
	}
}
