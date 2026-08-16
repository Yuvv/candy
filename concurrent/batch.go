package concurrent

import "context"

// AwaitTask runs fn for each input and returns results in input order.
func AwaitTask[IN any, OUT any](ctx context.Context, fn func(context.Context, IN) (OUT, error), inputs []IN) ([]OUT, error) {
	tg := NewTaskGroup[OUT](ctx)
	for _, input := range inputs {
		input := input
		tg.Go(func(ctx context.Context) (OUT, error) {
			return fn(ctx, input)
		})
	}
	results, err := tg.Wait()
	if err != nil {
		return nil, err
	}
	if results == nil && inputs != nil {
		return []OUT{}, nil
	}
	return results, nil
}

// AwaitBatchTask runs fn for each input batch and flattens batch results in batch order.
func AwaitBatchTask[IN any, OUT any](ctx context.Context, fn func(context.Context, []IN) ([]OUT, error), inputs []IN, batchSize int) ([]OUT, error) {
	if batchSize <= 0 {
		batchSize = 1
	}

	tg := NewTaskGroup[[]OUT](ctx)
	for _, batch := range chunk(inputs, batchSize) {
		batch := batch
		tg.Go(func(ctx context.Context) ([]OUT, error) {
			return fn(ctx, batch)
		})
	}
	parts, err := tg.Wait()
	if err != nil {
		return nil, err
	}

	out := make([]OUT, 0, len(inputs))
	for _, part := range parts {
		out = append(out, part...)
	}
	return out, nil
}

// AwaitBatchReturnMapTask runs fn for each input batch and merges returned maps.
func AwaitBatchReturnMapTask[IN comparable, OUT any](ctx context.Context, fn func(context.Context, []IN) (map[IN]OUT, error), inputs []IN, batchSize int) (map[IN]OUT, error) {
	if batchSize <= 0 {
		batchSize = 1
	}

	tg := NewTaskGroup[map[IN]OUT](ctx)
	for _, batch := range chunk(inputs, batchSize) {
		batch := batch
		tg.Go(func(ctx context.Context) (map[IN]OUT, error) {
			return fn(ctx, batch)
		})
	}
	parts, err := tg.Wait()
	if err != nil {
		return nil, err
	}

	out := make(map[IN]OUT, len(inputs))
	for _, part := range parts {
		for key, value := range part {
			out[key] = value
		}
	}
	return out, nil
}

func chunk[T any](items []T, size int) [][]T {
	if len(items) == 0 {
		if items == nil {
			return nil
		}
		return [][]T{}
	}

	out := make([][]T, 0, (len(items)+size-1)/size)
	for start := 0; start < len(items); start += size {
		end := start + size
		if end > len(items) {
			end = len(items)
		}
		out = append(out, items[start:end])
	}
	return out
}
