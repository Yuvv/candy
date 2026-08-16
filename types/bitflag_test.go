package types

import "testing"

func TestBitFlag(t *testing.T) {
	var uintptrFlags BitFlag[uintptr]
	uintptrFlags.Set(1)
	if !uintptrFlags.Has(1) {
		t.Fatalf("uintptr BitFlag missing set flag: %d", uintptrFlags.Value())
	}

	var flags BitFlag[uint8]
	flags.Set(1)
	flags.Set(4)
	if !flags.Has(1) || !flags.Has(4) || flags.Has(2) {
		t.Fatalf("unexpected flags after Set: %d", flags.Value())
	}
	flags.Clear(1)
	if flags.Has(1) || flags.Value() != 4 {
		t.Fatalf("unexpected flags after Clear: %d", flags.Value())
	}
	flags.Toggle(2)
	if !flags.Has(2) || flags.Value() != 6 {
		t.Fatalf("unexpected flags after Toggle on: %d", flags.Value())
	}
	flags.Toggle(2)
	if flags.Has(2) || flags.Value() != 4 {
		t.Fatalf("unexpected flags after Toggle off: %d", flags.Value())
	}
}
