package lists

import (
	"github.com/yuvv/candy/function"
	"github.com/yuvv/candy/iters"
)

type Itr[E comparable] struct {
	iters.AbstractIterator[E]

	cursor           int // 0
	lastRet          int // -1
	expectedModCount int // 0

	targetList List[E]
}

func (itr *Itr[E]) HasNext() bool {
	return itr.cursor != itr.targetList.Size()
}

func (itr *Itr[E]) Next() E {
	itr.checkForCoModification()

	i := itr.cursor
	next := itr.targetList.Get(i)
	itr.lastRet = i
	itr.cursor = i + 1
	return next
}

func (itr *Itr[E]) Remove() {
	if itr.lastRet < 0 {
		panic("IllegalStateException")
	}
	itr.checkForCoModification()
	itr.targetList.RemoveAt(itr.lastRet)
	if itr.lastRet < itr.cursor {
		itr.cursor--
	}
	itr.lastRet = -1
	itr.expectedModCount = itr.targetList.getModCount()
}

func (itr *Itr[E]) ForEachRemaining(action function.Consumer[E]) {
	for itr.HasNext() {
		action(itr.Next())
	}
}

func (itr *Itr[E]) checkForCoModification() {
	if itr.expectedModCount != itr.targetList.getModCount() {
		panic("ConcurrentModificationException")
	}
}

func NewItr[E comparable](lst List[E]) *Itr[E] {
	return &Itr[E]{
		cursor:           0,
		lastRet:          -1,
		expectedModCount: lst.getModCount(),
		targetList:       lst,
	}
}

// ---------------------------------------------------------------------

// ListItr implements ListIterator
type ListItr[E comparable] struct {
	Itr[E]
}

func (itr *ListItr[E]) HasPrevious() bool {
	return itr.cursor != 0
}

func (itr *ListItr[E]) Previous() E {
	itr.checkForCoModification()
	i := itr.cursor - 1
	previous := itr.targetList.Get(i)
	itr.lastRet = i
	itr.cursor = i
	return previous
}

func (itr *ListItr[E]) NextIndex() int {
	return itr.cursor
}

func (itr *ListItr[E]) PreviousIndex() int {
	return itr.cursor - 1
}

func (itr *ListItr[E]) Set(e E) {
	if itr.lastRet < 0 {
		panic("IllegalStateException")
	}
	itr.checkForCoModification()
	itr.targetList.Set(itr.lastRet, e)
	itr.expectedModCount = itr.targetList.getModCount()
}

func (itr *ListItr[E]) Add(e E) {
	itr.checkForCoModification()
	i := itr.cursor
	itr.targetList.AddAt(i, e)
	itr.lastRet = -1
	itr.cursor = i + 1
	itr.expectedModCount = itr.targetList.getModCount()
}
