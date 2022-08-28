package lists

import (
	"sort"

	"github.com/yuvv/candy/collections"
	"github.com/yuvv/candy/function"
	"github.com/yuvv/candy/iters"
	"github.com/yuvv/candy/lang"
)

type ArrayList[E comparable] struct {
	AbstractList[E]

	slice []E
}

func (lst *ArrayList[E]) Iterator() iters.Iterator[E] {
	return NewItr[E](lst)
}

func (lst *ArrayList[E]) ForEach(action function.Consumer[E]) {
	for _, item := range lst.slice {
		action(item)
	}
}

func (lst *ArrayList[E]) Spliterator() iters.Spliterator[E] {
	//TODO implement me
	panic("implement me")
}

func (lst *ArrayList[E]) Size() int {
	return len(lst.slice)
}

func (lst *ArrayList[E]) IsEmpty() bool {
	return lst.Size() == 0
}

func (lst *ArrayList[E]) Contains(o E) bool {
	for _, it := range lst.slice {
		if it == o {
			return true
		}
	}
	return false
}

func (lst *ArrayList[E]) ToArray() []E {
	toReturn := make([]E, lst.Size())
	for i, item := range lst.slice {
		toReturn[i] = item
	}
	return toReturn
}

func (lst *ArrayList[E]) Add(e E) bool {
	lst.modCount++
	lst.slice = append(lst.slice, e)
	return true
}

func (lst *ArrayList[E]) Remove(o E) bool {
	for i, item := range lst.slice {
		if item == o {
			for j := i + 1; j < len(lst.slice); j++ {
				lst.slice[j-1] = lst.slice[j]
			}
			lst.slice = lst.slice[:len(lst.slice)-1]
			lst.modCount++
			return true
		}
	}
	return false
}

func (lst *ArrayList[E]) ContainsAll(collection collections.Collection[E]) bool {
	it := collection.Iterator()
	for it.HasNext() {
		if !lst.Contains(it.Next()) {
			return false
		}
	}
	return true
}

func (lst *ArrayList[E]) AddAll(collection collections.Collection[E]) bool {
	lst.modCount++
	modified := false
	it := collection.Iterator()
	for it.HasNext() {
		modified = lst.Add(it.Next())
	}
	return modified
}

func (lst *ArrayList[E]) RemoveAll(collection collections.Collection[E]) bool {
	return lst.RemoveIf(collection.Contains)
}

func (lst *ArrayList[E]) RemoveIf(predicate function.Predicate[E]) bool {
	i := 0
	for j := 0; j < len(lst.slice); j++ {
		if !predicate(lst.slice[j]) {
			lst.slice[i] = lst.slice[j]
			i++
		}
	}
	if i >= len(lst.slice) {
		return false
	}
	lst.slice = lst.slice[:i]
	lst.modCount++
	return true
}

func (lst *ArrayList[E]) RetainAll(collection collections.Collection[E]) bool {
	modified := false
	it := lst.Iterator()
	for it.HasNext() {
		if !collection.Contains(it.Next()) {
			it.Remove()
			modified = true
		}
	}
	return modified
}

func (lst *ArrayList[E]) Clear() {
	lst.modCount++
	lst.slice = lst.slice[:0]
}

func (lst *ArrayList[E]) AddAllAt(idx int, collection collections.Collection[E]) {
	lst.modCount++
	nSlice := make([]E, len(lst.slice)+collection.Size())
	i := 0
	for i < idx {
		nSlice[i] = lst.slice[i]
		i++
	}
	it := collection.Iterator()
	for it.HasNext() {
		nSlice[i] = it.Next()
		i++
	}
	for j := idx; j < len(lst.slice); j++ {
		nSlice[i] = lst.slice[j]
		i++
	}
	lst.slice = nSlice
}

func (lst *ArrayList[E]) ReplaceAll(operator function.UnaryOperator[E]) {
	lst.modCount++
	for i := 0; i < len(lst.slice); i++ {
		lst.slice[i] = operator(lst.slice[i])
	}
}

func (lst *ArrayList[E]) Sort(comparator lang.Comparator[E]) {
	sort.Slice(lst.slice, func(i, j int) bool {
		return comparator.Compare(lst.slice[i], lst.slice[j]) < 0
	})
	lst.modCount++
}

func (lst *ArrayList[E]) Get(idx int) E {
	return lst.slice[idx]
}

func (lst *ArrayList[E]) Set(idx int, ele E) E {
	origin := lst.slice[idx]
	lst.slice[idx] = ele
	return origin
}

func (lst *ArrayList[E]) AddAt(idx int, ele E) {
	lst.modCount++
	lst.slice = append(lst.slice, lst.slice[len(lst.slice)-1])
	for i := len(lst.slice) - 2; i >= idx; i-- {
		lst.slice[i+1] = lst.slice[i]
	}
	lst.slice[idx] = ele
}

func (lst *ArrayList[E]) RemoveAt(idx int) E {
	origin := lst.slice[idx]
	for i := idx + 1; i < len(lst.slice); i++ {
		lst.slice[i-1] = lst.slice[i]
	}
	lst.slice = lst.slice[:len(lst.slice)-1]
	lst.modCount++
	return origin
}

func (lst *ArrayList[E]) IndexOf(o E) int {
	for i, item := range lst.slice {
		if item == o {
			return i
		}
	}
	return -1
}

func (lst *ArrayList[E]) LastIndexOf(o E) int {
	for i := len(lst.slice) - 1; i >= 0; i-- {
		if lst.slice[i] == o {
			return i
		}
	}
	return -1
}

func (lst *ArrayList[E]) ListIterator() iters.ListIterator[E] {
	return NewListItr[E](lst)
}

func (lst *ArrayList[E]) ListIteratorFrom(idx int) iters.ListIterator[E] {
	return NewListItrFrom[E](lst, idx)
}

// SubList return a new list with elements from fromIdx (inclusive) to toIdx (exclusive)
func (lst *ArrayList[E]) SubList(fromIdx, toIdx int) List[E] {
	return &ArrayList[E]{
		AbstractList: AbstractList[E]{
			modCount: lst.modCount,
		},
		slice: lst.slice[fromIdx:toIdx],
	}
}

func (lst *ArrayList[E]) getModCount() int {
	return lst.modCount
}

/// ---------------------------------------------------------------------------

func NewArrayList[E comparable]() *ArrayList[E] {
	return &ArrayList[E]{
		slice: []E{},
	}
}

func NewArrayListWithCap[E comparable](cap int) *ArrayList[E] {
	if cap < 0 {
		panic("invalid capacity")
	}
	return &ArrayList[E]{
		slice: make([]E, 0, cap),
	}
}

func NewArrayListWithEle[E comparable](elements ...E) *ArrayList[E] {
	list := NewArrayListWithCap[E](len(elements))
	for _, item := range elements {
		list.Add(item)
	}
	return list
}
