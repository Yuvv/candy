package strs

import "testing"

func TestJoinSliceAndJoiner(t *testing.T) {
	got := JoinSlice([]int{1, 2, 3}, func(v int) string { return NewBuilder().Append(v).String() }, "[", ",", "]")
	if got != "[1,2,3]" {
		t.Fatalf("JoinSlice = %q", got)
	}

	joiner := NewStringJoiner("|").WithPrefix("<").WithSuffix(">")
	if got := joiner.Join([]string{"a", "b"}); got != "<a|b>" {
		t.Fatalf("Joiner.Join = %q", got)
	}
}

func TestNumericJoiners(t *testing.T) {
	if got := NewIntJoiner(",").Join([]int{1, 2}); got != "1,2" {
		t.Fatalf("NewIntJoiner = %q", got)
	}
	if got := NewUint64Joiner(";").Join([]uint64{3, 4}); got != "3;4" {
		t.Fatalf("NewUint64Joiner = %q", got)
	}
}
