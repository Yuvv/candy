package iter

import (
	"fmt"
	"testing"
)

func assertPanicMessage(t *testing.T, want string, action func()) {
	t.Helper()
	defer func() {
		got := recover()
		if got == nil {
			t.Fatalf("expected panic %q", want)
		}
		if message := fmt.Sprint(got); message != want {
			t.Fatalf("panic = %q, want %q", message, want)
		}
	}()
	action()
}

func TestAbstractIteratorPanicsIdentifyRequiredMethod(t *testing.T) {
	iterator := &AbstractIterator[int]{}
	tests := []struct {
		name   string
		want   string
		action func()
	}{
		{name: "HasNext", want: "iter: AbstractIterator.HasNext must be implemented by concrete iterator", action: func() { iterator.HasNext() }},
		{name: "Next", want: "iter: AbstractIterator.Next must be implemented by concrete iterator", action: func() { iterator.Next() }},
		{name: "Remove", want: "iter: AbstractIterator.Remove must be implemented by concrete iterator", action: iterator.Remove},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assertPanicMessage(t, test.want, test.action)
		})
	}
}
