package concurrent

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"
)

func TestTaskGroupWaitCollectsResultsInGoCallOrder(t *testing.T) {
	tg := NewTaskGroup[int](context.Background())

	tg.Go(func(ctx context.Context) (int, error) {
		<-time.After(20 * time.Millisecond)
		return 1, nil
	})
	tg.Go(func(ctx context.Context) (int, error) {
		return 2, nil
	})
	tg.Go(func(ctx context.Context) (int, error) {
		return 3, nil
	})

	got, err := tg.Wait()
	if err != nil {
		t.Fatalf("Wait error = %v", err)
	}
	want := []int{1, 2, 3}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Wait results = %#v, want %#v", got, want)
	}
}

func TestTaskGroupReturnsFirstErrorAndCancelsContext(t *testing.T) {
	boom := errors.New("boom")
	tg := NewTaskGroup[int](context.Background())
	started := make(chan struct{})
	releaseError := make(chan struct{})
	cancelled := make(chan struct{})

	tg.Go(func(ctx context.Context) (int, error) {
		close(started)
		<-releaseError
		return 0, boom
	})
	tg.Go(func(ctx context.Context) (int, error) {
		select {
		case <-ctx.Done():
			close(cancelled)
			return 0, ctx.Err()
		case <-time.After(time.Second):
			t.Fatal("context was not cancelled")
			return 0, nil
		}
	})

	<-started
	close(releaseError)
	_, err := tg.Wait()
	if !errors.Is(err, boom) {
		t.Fatalf("Wait error = %v, want boom", err)
	}
	select {
	case <-cancelled:
	case <-time.After(time.Second):
		t.Fatal("second task did not observe cancellation")
	}
}

func TestNewTaskGroupAcceptsNilContext(t *testing.T) {
	tg := NewTaskGroup[string](nil)
	tg.Go(func(ctx context.Context) (string, error) {
		if ctx == nil {
			t.Fatal("task context is nil")
		}
		return "ok", nil
	})

	got, err := tg.Wait()
	if err != nil {
		t.Fatalf("Wait error = %v", err)
	}
	if !reflect.DeepEqual(got, []string{"ok"}) {
		t.Fatalf("Wait results = %#v", got)
	}
}
