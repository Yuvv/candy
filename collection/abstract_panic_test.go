package collection

import (
	"fmt"
	"testing"
)

func assertExactPanic(t *testing.T, want string, action func()) {
	t.Helper()
	defer func() {
		got := recover()
		if got == nil {
			t.Fatalf("expected panic %q", want)
		}
		if message := fmt.Sprint(got); message != want {
			t.Fatalf("panic = %q, want %q", message, want)
		}
	}()
	action()
}

func TestAbstractCollectionPanicsExplainRequiredOrUnsupportedMethods(t *testing.T) {
	collection := &AbstractCollection[int]{}
	tests := []struct {
		name   string
		want   string
		action func()
	}{
		{name: "Iterator", want: "collection: AbstractCollection.Iterator must be implemented by concrete collection", action: func() { collection.Iterator() }},
		{name: "Size", want: "collection: AbstractCollection.Size must be implemented by concrete collection", action: func() { collection.Size() }},
		{name: "Add", want: "collection: AbstractCollection.Add must be implemented by concrete collection", action: func() { collection.Add(1) }},
		{name: "Spliterator", want: "collection: AbstractCollection.Spliterator is unsupported", action: func() { collection.Spliterator() }},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assertExactPanic(t, test.want, test.action)
		})
	}
}

func TestAbstractListPanicIdentifiesRequiredMethod(t *testing.T) {
	list := &AbstractList[int]{}
	assertExactPanic(t, "collection: AbstractList.Get must be implemented by concrete list", func() { list.Get(0) })
}

func TestAbstractMapPanicIdentifiesRequiredMethod(t *testing.T) {
	abstractMap := &AbstractMap[string, int]{}
	assertExactPanic(t, "collection: AbstractMap.Size must be implemented by concrete map", func() { abstractMap.Size() })
}

func TestConcreteSpliteratorPanicsExplainUnsupportedOperation(t *testing.T) {
	t.Run("ArrayList", func(t *testing.T) {
		assertExactPanic(t, "collection: ArrayList.Spliterator is unsupported", func() { NewArrayList[int]().Spliterator() })
	})
	t.Run("HashSet", func(t *testing.T) {
		assertExactPanic(t, "collection: HashSet.Spliterator is unsupported", func() { NewHashSet[int]().Spliterator() })
	})
}
