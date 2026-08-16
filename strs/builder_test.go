package strs

import "testing"

func TestBuilderChaining(t *testing.T) {
	b := NewBuilder().AppendString("go").AppendByte('-').Append(18)
	if got := b.String(); got != "go-18" {
		t.Fatalf("String = %q", got)
	}
	if b.Len() != len("go-18") {
		t.Fatalf("Len = %d", b.Len())
	}
}
