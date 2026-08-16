package concurrent

import (
	"context"
	"sync"
)

// TaskGroup runs context-aware tasks and collects their results in Go call order.
type TaskGroup[OUT any] struct {
	ctx    context.Context
	cancel context.CancelFunc

	wg sync.WaitGroup

	mu       sync.Mutex
	results  []OUT
	firstErr error
}

// NewTaskGroup creates a TaskGroup with a cancellable child context.
func NewTaskGroup[OUT any](ctx context.Context) *TaskGroup[OUT] {
	if ctx == nil {
		ctx = context.Background()
	}
	child, cancel := context.WithCancel(ctx)
	return &TaskGroup[OUT]{ctx: child, cancel: cancel}
}

// Go starts fn in its own goroutine.
func (g *TaskGroup[OUT]) Go(fn func(context.Context) (OUT, error)) {
	g.mu.Lock()
	index := len(g.results)
	var zero OUT
	g.results = append(g.results, zero)
	g.wg.Add(1)
	g.mu.Unlock()

	go func() {
		defer g.wg.Done()

		result, err := fn(g.ctx)

		g.mu.Lock()
		defer g.mu.Unlock()
		if err != nil {
			if g.firstErr == nil {
				g.firstErr = err
				g.cancel()
			}
			return
		}
		g.results[index] = result
	}()
}

// Wait waits for all tasks to finish and returns their results in Go call order.
func (g *TaskGroup[OUT]) Wait() ([]OUT, error) {
	g.wg.Wait()
	g.cancel()

	g.mu.Lock()
	defer g.mu.Unlock()
	if g.firstErr != nil {
		return nil, g.firstErr
	}
	results := append([]OUT(nil), g.results...)
	return results, nil
}
