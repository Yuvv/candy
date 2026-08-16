package optional

import "testing"

func TestEmptyAndOfPresenceAndGet(t *testing.T) {
	var empty Optional[int] = Empty[int]()
	if !empty.IsEmpty() {
		t.Fatalf("Empty should be empty")
	}
	if empty.IsPresent() {
		t.Fatalf("Empty should not be present")
	}

	var present Optional[int] = Of(0)
	if !present.IsPresent() {
		t.Fatalf("Of should be present")
	}
	if present.IsEmpty() {
		t.Fatalf("Of should not be empty")
	}
	if got := present.Get(); got != 0 {
		t.Fatalf("Get() = %d, want 0", got)
	}
}

func TestGetPanicsOnEmpty(t *testing.T) {
	defer func() {
		got := recover()
		if got == nil {
			t.Fatalf("Get should panic on empty")
		}
		if got != "optional: no value present" {
			t.Fatalf("panic = %v, want %q", got, "optional: no value present")
		}
	}()

	_ = Empty[string]().Get()
}

func TestOrElseAndOrElseGet(t *testing.T) {
	empty := Empty[string]()
	if got := empty.OrElse("fallback"); got != "fallback" {
		t.Fatalf("empty.OrElse() = %q, want fallback", got)
	}

	present := Of("value")
	if got := present.OrElse("fallback"); got != "value" {
		t.Fatalf("present.OrElse() = %q, want value", got)
	}

	calls := 0
	if got := empty.OrElseGet(func() string {
		calls++
		return "generated"
	}); got != "generated" {
		t.Fatalf("empty.OrElseGet() = %q, want generated", got)
	}
	if calls != 1 {
		t.Fatalf("empty supplier calls = %d, want 1", calls)
	}

	if got := present.OrElseGet(func() string {
		calls++
		return "unused"
	}); got != "value" {
		t.Fatalf("present.OrElseGet() = %q, want value", got)
	}
	if calls != 1 {
		t.Fatalf("present supplier calls = %d, want still 1", calls)
	}
}

func TestIfPresentNotCalledOnEmpty(t *testing.T) {
	called := false
	Empty[int]().IfPresent(func(value int) {
		called = true
	})
	if called {
		t.Fatalf("IfPresent should not call consumer for empty optional")
	}

	var got int
	Of(42).IfPresent(func(value int) {
		got = value
	})
	if got != 42 {
		t.Fatalf("IfPresent got %d, want 42", got)
	}
}

func TestMap(t *testing.T) {
	empty := Map(Empty[int](), func(value int) string {
		t.Fatalf("mapper should not be called for empty optional")
		return ""
	})
	if !empty.IsEmpty() {
		t.Fatalf("Map should keep empty optional empty")
	}

	present := Map(Of(21), func(value int) string {
		return "value"
	})
	if !present.IsPresent() {
		t.Fatalf("Map should return present optional for present input")
	}
	if got := present.Get(); got != "value" {
		t.Fatalf("mapped value = %q, want value", got)
	}
}

func TestFlatMap(t *testing.T) {
	empty := FlatMap(Empty[int](), func(value int) Optional[string] {
		t.Fatalf("mapper should not be called for empty optional")
		return Of("")
	})
	if !empty.IsEmpty() {
		t.Fatalf("FlatMap should keep empty optional empty")
	}

	present := FlatMap(Of(21), func(value int) Optional[string] {
		return Of("mapped")
	})
	if !present.IsPresent() {
		t.Fatalf("FlatMap should return present optional for present input")
	}
	if got := present.Get(); got != "mapped" {
		t.Fatalf("flat mapped value = %q, want mapped", got)
	}
}
