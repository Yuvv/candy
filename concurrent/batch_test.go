package concurrent

import (
	"context"
	"errors"
	"reflect"
	"sort"
	"sync"
	"testing"
	"time"
)

func TestAwaitTaskPreservesInputOrder(t *testing.T) {
	got, err := AwaitTask(context.Background(), func(ctx context.Context, in int) (int, error) {
		if in == 3 {
			<-time.After(20 * time.Millisecond)
		}
		return in * 10, nil
	}, []int{3, 1, 2})
	if err != nil {
		t.Fatalf("AwaitTask error = %v", err)
	}
	want := []int{30, 10, 20}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("AwaitTask = %#v, want %#v", got, want)
	}
}

func TestAwaitBatchTaskDefaultsBatchSizeAndFlattensInBatchOrder(t *testing.T) {
	var mu sync.Mutex
	seenSizes := make([]int, 0)
	got, err := AwaitBatchTask(context.Background(), func(ctx context.Context, in []int) ([]string, error) {
		mu.Lock()
		seenSizes = append(seenSizes, len(in))
		mu.Unlock()

		out := make([]string, 0, len(in))
		for _, v := range in {
			out = append(out, string(rune('a'+v-1)))
		}
		return out, nil
	}, []int{1, 2, 3}, 0)
	if err != nil {
		t.Fatalf("AwaitBatchTask error = %v", err)
	}
	sort.Ints(seenSizes)
	if !reflect.DeepEqual(seenSizes, []int{1, 1, 1}) {
		t.Fatalf("batch sizes = %#v", seenSizes)
	}
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("AwaitBatchTask = %#v, want %#v", got, want)
	}
}

func TestAwaitBatchTaskFlattensBatchesInOrder(t *testing.T) {
	got, err := AwaitBatchTask(context.Background(), func(ctx context.Context, in []int) ([]int, error) {
		if in[0] == 1 {
			<-time.After(20 * time.Millisecond)
		}
		out := make([]int, len(in))
		copy(out, in)
		return out, nil
	}, []int{1, 2, 3, 4, 5}, 2)
	if err != nil {
		t.Fatalf("AwaitBatchTask error = %v", err)
	}
	want := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("AwaitBatchTask = %#v, want %#v", got, want)
	}
}

func TestAwaitBatchReturnMapTaskMergesMaps(t *testing.T) {
	got, err := AwaitBatchReturnMapTask(context.Background(), func(ctx context.Context, in []int) (map[int]string, error) {
		out := make(map[int]string, len(in))
		for _, v := range in {
			out[v] = string(rune('a' + v - 1))
		}
		return out, nil
	}, []int{1, 2, 3}, 2)
	if err != nil {
		t.Fatalf("AwaitBatchReturnMapTask error = %v", err)
	}
	want := map[int]string{1: "a", 2: "b", 3: "c"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("AwaitBatchReturnMapTask = %#v, want %#v", got, want)
	}
}

func TestAwaitTaskReturnsError(t *testing.T) {
	boom := errors.New("boom")
	_, err := AwaitTask(context.Background(), func(ctx context.Context, in int) (int, error) {
		if in == 2 {
			return 0, boom
		}
		return in, nil
	}, []int{1, 2, 3})
	if !errors.Is(err, boom) {
		t.Fatalf("AwaitTask error = %v, want boom", err)
	}
}
