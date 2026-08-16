package types

import "testing"

func TestBitmapZeroValue(t *testing.T) {
	var b Bitmap

	if b.Get(0) {
		t.Fatal("zero value Bitmap reports bit 0 set")
	}
	if b.Get(70) {
		t.Fatal("zero value Bitmap reports bit 70 set")
	}
	if got := b.Len(); got != 0 {
		t.Fatalf("Len() = %d, want 0", got)
	}
	if got := b.Words(); got != 0 {
		t.Fatalf("Words() = %d, want 0", got)
	}
	if got := b.String(); got != "b:0" {
		t.Fatalf("String() = %q, want %q", got, "b:0")
	}
}

func TestBitmapSetGetDynamicGrowth(t *testing.T) {
	var b Bitmap

	b.Set(0)
	b.Set(70)

	if !b.Get(0) {
		t.Fatal("bit 0 is not set")
	}
	if !b.Get(70) {
		t.Fatal("bit 70 is not set")
	}
	if b.Get(69) {
		t.Fatal("bit 69 is unexpectedly set")
	}
	if got := b.Len(); got != 71 {
		t.Fatalf("Len() = %d, want 71", got)
	}
	if got := b.Words(); got != 2 {
		t.Fatalf("Words() = %d, want 2", got)
	}
}

func TestBitmapUnsetMissingAndPresentBits(t *testing.T) {
	var b Bitmap

	b.Unset(5)
	if got := b.Len(); got != 0 {
		t.Fatalf("Len() after unsetting missing bit = %d, want 0", got)
	}
	if got := b.Words(); got != 0 {
		t.Fatalf("Words() after unsetting missing bit = %d, want 0", got)
	}

	b.Set(1)
	b.Set(70)
	b.Unset(69)
	if !b.Get(1) || !b.Get(70) {
		t.Fatal("unsetting missing bit changed set bits")
	}

	b.Unset(1)
	if b.Get(1) {
		t.Fatal("bit 1 is still set after Unset")
	}
	if !b.Get(70) {
		t.Fatal("unsetting bit 1 changed bit 70")
	}
}

func TestBitmapLenShrinksAfterUnsettingHighestBit(t *testing.T) {
	var b Bitmap

	b.Set(2)
	b.Set(70)
	if got := b.Len(); got != 71 {
		t.Fatalf("Len() before unsetting highest bit = %d, want 71", got)
	}

	b.Unset(70)
	if got := b.Len(); got != 3 {
		t.Fatalf("Len() after unsetting highest bit = %d, want 3", got)
	}
	if got := b.Words(); got != 2 {
		t.Fatalf("Words() after unsetting highest bit = %d, want 2", got)
	}

	b.Unset(2)
	if got := b.Len(); got != 0 {
		t.Fatalf("Len() after unsetting all bits = %d, want 0", got)
	}
	if got := b.Words(); got != 2 {
		t.Fatalf("Words() after unsetting all bits = %d, want 2", got)
	}
}

func TestBitmapString(t *testing.T) {
	var b Bitmap

	b.Set(0)
	b.Set(3)

	if got := b.String(); got != "b:1001" {
		t.Fatalf("String() = %q, want %q", got, "b:1001")
	}
}
