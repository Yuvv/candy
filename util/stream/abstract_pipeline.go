package stream

import (
	"github.com/yuvv/jog/lang"
	"github.com/yuvv/jog/util"
	"github.com/yuvv/jog/util/function"
)

type PipelineHelper[POUT any] interface {
	//GetSourceShape() StreamShape

	GetStreamAndOpFlags() int

	ExactOutputSizeIfKnown[PIN](spliterator util.Spliterator[PIN])

	// todo: others
}

type _AbstractPipeline[EIN, EOUT any, S BaseStream[EOUT, S]] struct {
	PipelineHelper[EOUT]

	// sourceStage is backlink to the head of the pipeline chain (self if this is the source stage).
	sourceStage *_AbstractPipeline[EIN, EOUT, S]

	// previousStage is the "upstream" pipeline, or null if this is the source stage.
	previousStage *_AbstractPipeline[EIN, EOUT, S]

	// nextStage is the next stage in the pipeline, or null if this is the last stage. Effectively final at the point of linking to the next pipeline.
	nextStage *_AbstractPipeline[EIN, EOUT, S]

	sourceOrOpFlags int

	depth int

	combinedFlags int

	sourceSpliterator util.Spliterator[any]

	sourceSupplier function.Supplier[util.Spliterator[any]]

	sourceCloseAction lang.Runnable

	linkedOrConsumed bool

	sourceAnyStateful bool

	parallel bool
}

func (receiver _AbstractPipeline[EIN, EOUT, S]) IsParallel() bool {
	return receiver.sourceStage.parallel
}

func (receiver _AbstractPipeline[EIN, EOUT, S]) GetStreamFlags() int {
	return toStreamFlags(receiver.combinedFlags)
}
