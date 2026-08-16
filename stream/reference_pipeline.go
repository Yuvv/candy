package stream

import (
	"github.com/yuvv/candy/function"
	"github.com/yuvv/candy/iter"
)

type _ReferencePipeline[PIN, POUT any] struct {
	_AbstractPipeline[PIN, POUT, Stream[POUT]]
}

func (receiver _ReferencePipeline[EIN, EOUT]) sourceStageSpliterator() iter.Spliterator[EOUT] {
	// todo
	return nil
}

func (receiver _ReferencePipeline[EIN, EOUT]) forEachOrdered(consumer function.Consumer[EOUT]) {
	// todo
}

func (receiver _ReferencePipeline[EIN, EOUT]) forEach(consumer function.Consumer[EOUT]) {
	// todo
}

// -------------------------------- this is a divider ----------------------

// _Head is a head pipeline
type _Head[EIN, EOUT any] struct {
	_ReferencePipeline[EIN, EOUT]
}

func (receiver _Head[EIN, EOUT]) opIsStateful() bool {
	panic("UnsupportedOperationException")
}

func (receiver _Head[EIN, EOUT]) opWrapSink() bool {
	panic("UnsupportedOperationException")
}

func (receiver _Head[EIN, EOUT]) forEach(consumer function.Consumer[EOUT]) {
	if receiver.IsParallel() {
		receiver._ReferencePipeline.forEach(consumer)
	} else {
		receiver.sourceStageSpliterator().ForEachRemaining(consumer)
	}
}

func (receiver _Head[EIN, EOUT]) forEachOrdered(consumer function.Consumer[EOUT]) {
	if receiver.IsParallel() {
		receiver._ReferencePipeline.forEachOrdered(consumer)
	} else {
		receiver.sourceStageSpliterator().ForEachRemaining(consumer)
	}
}
