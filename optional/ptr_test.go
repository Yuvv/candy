package optional

import "testing"

func TestOfPtrNilIsEmpty(t *testing.T) {
	optional := OfPtr[int](nil)
	if !optional.IsEmpty() {
		t.Fatalf("OfPtr(nil) should be empty")
	}
	if optional.IsPresent() {
		t.Fatalf("OfPtr(nil) should not be present")
	}
}

func TestOfPtrStoresDereferencedValue(t *testing.T) {
	value := 10
	optional := OfPtr(&value)
	if !optional.IsPresent() {
		t.Fatalf("OfPtr(&value) should be present")
	}
	if got := optional.Get(); got != 10 {
		t.Fatalf("OfPtr(&value).Get() = %d, want 10", got)
	}
}

func TestOfPtrCopiesValue(t *testing.T) {
	value := "before"
	optional := OfPtr(&value)

	value = "after"

	if got := optional.Get(); got != "before" {
		t.Fatalf("optional value after original mutation = %q, want before", got)
	}
}
