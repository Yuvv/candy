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
