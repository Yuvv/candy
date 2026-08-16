package stream

import (
	"testing"

	"github.com/yuvv/candy/function"
	"github.com/yuvv/candy/iter"
)

func assertPanicMessage(t *testing.T, want string, fn func()) {
	t.Helper()

	defer func() {
		if got := recover(); got != want {
			t.Fatalf("panic = %v, want %q", got, want)
		}
	}()

	fn()
}

func TestOfPanicsWhenUnimplemented(t *testing.T) {
	assertPanicMessage(t, "stream: Of is not implemented", func() {
		Of(1, 2, 3)
	})
}

func TestEmptyPanicsWhenUnimplemented(t *testing.T) {
	assertPanicMessage(t, "stream: Empty is not implemented", func() {
		Empty[int]()
	})
}

func TestConcatPanicsWhenUnimplemented(t *testing.T) {
	assertPanicMessage(t, "stream: Concat is not implemented", func() {
		Concat[int](nil, nil)
	})
}

func TestSupportBySpliteratorPanicsWhenUnimplemented(t *testing.T) {
	assertPanicMessage(t, "stream: SupportBySpliterator is not implemented", func() {
		SupportBySpliterator[int](nil, false)
	})
}

func TestSupportBySupplierPanicsWhenUnimplemented(t *testing.T) {
	assertPanicMessage(t, "stream: SupportBySupplier is not implemented", func() {
		var supplier function.Supplier[iter.Spliterator[int]]
		SupportBySupplier(supplier, C_CONCURRENT, false)
	})
}
