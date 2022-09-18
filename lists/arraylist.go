package lists

import (
	"sort"

	"github.com/yuvv/candy/collections"
	"github.com/yuvv/candy/function"
	"github.com/yuvv/candy/iters"
	"github.com/yuvv/candy/lang"
)

type _ArrayList[E any] struct {
	AbstractList[E]

	slice []E
}

func (lst *_ArrayList[E]) Iterator() iters.Iterator[E] {
	return NewItr[E](lst)
}

func (lst *_ArrayList[E]) ForEach(action function.Consumer[E]) {
	for _, item := range lst.slice {
		action(item)
	}
}

func (lst *_ArrayList[E]) Spliterator() iters.Spliterator[E] {
	//TODO implement me
	panic("implement me")
}

func (lst *_ArrayList[E]) Size() int {
	return len(lst.slice)
}

func (lst *_ArrayList[E]) IsEmpty() bool {
	return lst.Size() == 0
}

func (lst *_ArrayList[E]) Contains(o E) bool {
	for _, it := range lst.slice {
		if lst.equalMethod(it, o) {
			return true
		}
	}
	return false
}

func (lst *_ArrayList[E]) ToArray() []E {
	toReturn := make([]E, lst.Size())
	for i, item := range lst.slice {
		toReturn[i] = item
	}
	return toReturn
}

func (lst *_ArrayList[E]) Add(e E) bool {
	lst.modCount++
	lst.slice = append(lst.slice, e)
	return true
}

func (lst *_ArrayList[E]) Remove(o E) bool {
	for i, item := range lst.slice {
		if lst.equalMethod(item, o) {
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

func (lst *_ArrayList[E]) ContainsAll(collection collections.Collection[E]) bool {
	it := collection.Iterator()
	for it.HasNext() {
		if !lst.Contains(it.Next()) {
			return false
		}
	}
	return true
}

func (lst *_ArrayList[E]) AddAll(collection collections.Collection[E]) bool {
	lst.modCount++
	modified := false
	it := collection.Iterator()
	for it.HasNext() {
		modified = lst.Add(it.Next())
	}
	return modified
}

func (lst *_ArrayList[E]) RemoveAll(collection collections.Collection[E]) bool {
	return lst.RemoveIf(collection.Contains)
}

func (lst *_ArrayList[E]) RemoveIf(predicate function.Predicate[E]) bool {
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

func (lst *_ArrayList[E]) RetainAll(collection collections.Collection[E]) bool {
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

func (lst *_ArrayList[E]) Clear() {
	lst.modCount++
	lst.slice = lst.slice[:0]
}

func (lst *_ArrayList[E]) AddAllAt(idx int, collection collections.Collection[E]) {
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

func (lst *_ArrayList[E]) ReplaceAll(operator function.UnaryOperator[E]) {
	lst.modCount++
	for i := 0; i < len(lst.slice); i++ {
		lst.slice[i] = operator(lst.slice[i])
	}
}

func (lst *_ArrayList[E]) Sort(comparator lang.Comparator[E]) {
	sort.Slice(lst.slice, func(i, j int) bool {
		return comparator.Compare(lst.slice[i], lst.slice[j]) < 0
	})
	lst.modCount++
}

func (lst *_ArrayList[E]) Get(idx int) E {
	return lst.slice[idx]
}

func (lst *_ArrayList[E]) Set(idx int, ele E) E {
	origin := lst.slice[idx]
	lst.slice[idx] = ele
	return origin
}

func (lst *_ArrayList[E]) AddAt(idx int, ele E) {
	lst.modCount++
	lst.slice = append(lst.slice, lst.slice[len(lst.slice)-1])
	for i := len(lst.slice) - 2; i >= idx; i-- {
		lst.slice[i+1] = lst.slice[i]
	}
	lst.slice[idx] = ele
}

func (lst *_ArrayList[E]) RemoveAt(idx int) E {
	origin := lst.slice[idx]
	for i := idx + 1; i < len(lst.slice); i++ {
		lst.slice[i-1] = lst.slice[i]
	}
	lst.slice = lst.slice[:len(lst.slice)-1]
	lst.modCount++
	return origin
}

func (lst *_ArrayList[E]) IndexOf(o E) int {
	for i, item := range lst.slice {
		if lst.equalMethod(item, o) {
			return i
		}
	}
	return -1
}

func (lst *_ArrayList[E]) LastIndexOf(o E) int {
	for i := len(lst.slice) - 1; i >= 0; i-- {
		if lst.equalMethod(lst.slice[i], o) {
			return i
		}
	}
	return -1
}

func (lst *_ArrayList[E]) ListIterator() iters.ListIterator[E] {
	return NewListItr[E](lst)
}

func (lst *_ArrayList[E]) ListIteratorFrom(idx int) iters.ListIterator[E] {
	return NewListItrFrom[E](lst, idx)
}

// SubList return a new list with elements from fromIdx (inclusive) to toIdx (exclusive)
func (lst *_ArrayList[E]) SubList(fromIdx, toIdx int) List[E] {
	return &_ArrayList[E]{
		AbstractList: AbstractList[E]{
			modCount: lst.modCount,
		},
		slice: lst.slice[fromIdx:toIdx],
	}
}

func (lst *_ArrayList[E]) getModCount() int {
	return lst.modCount
}

/// ---------------------------------------------------------------------------

// NewArrayList return a pointer of _ArrayList[E] with default capacity of 4.
func NewArrayList[E comparable]() *_ArrayList[E] {
	return NewArrayListWithCap[E](4)
}

// NewArrayListWithCap return a pointer of _ArrayList[E] with specific capacity of cap.
func NewArrayListWithCap[E comparable](cap int) *_ArrayList[E] {
	if cap < 0 {
		panic("invalid capacity")
	}
	return &_ArrayList[E]{
		AbstractList: AbstractList[E]{
			modCount: 0,
			equalMethod: func(a E, o any) bool {
				return a == o
			},
		},
		slice: make([]E, 0, cap),
	}
}

// NewArrayListWithEle return a pointer of _ArrayList[E] contains all the given elements.
func NewArrayListWithEle[E comparable](elements ...E) *_ArrayList[E] {
	list := NewArrayListWithCap[E](len(elements))
	for _, item := range elements {
		list.Add(item)
	}
	return list
}

// NewSpecArrayList return a pointer of _ArrayList[E] with default capacity of 4.
// The function you passwd is used to check if other object equals to some element of the list.
func NewSpecArrayList[E any](em func(a E, o any) bool) *_ArrayList[E] {
	return NewSpecArrayListWithCap[E](em, 4)
}

// NewSpecArrayListWithCap return a pointer of _ArrayList[E] with specific capacity of cap.
// The function you passwd is used to check if other object equals to some element of the list.
func NewSpecArrayListWithCap[E any](em func(a E, o any) bool, cap int) *_ArrayList[E] {
	if cap < 0 {
		panic("invalid capacity")
	}
	return &_ArrayList[E]{
		AbstractList: AbstractList[E]{
			modCount:    0,
			equalMethod: em,
		},
		slice: make([]E, 0, cap),
	}
}

// NewSpecArrayListWithEle return a pointer of _ArrayList[E] with specific capacity of cap.
// The function you passwd is used to check if other object equals to some element of the list.
func NewSpecArrayListWithEle[E any](em func(a E, o any) bool, elements ...E) *_ArrayList[E] {
	list := NewSpecArrayListWithCap[E](em, len(elements))
	for _, item := range elements {
		list.Add(item)
	}
	return list
}
