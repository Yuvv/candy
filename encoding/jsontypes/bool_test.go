package jsontypes

import (
	"encoding/json"
	"testing"
)

func TestAdaptiveBoolUnmarshalAcceptedValues(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantNil bool
		want    bool
	}{
		{name: "true literal", input: `true`, want: true},
		{name: "false literal", input: `false`, want: false},
		{name: "quoted true", input: `"true"`, want: true},
		{name: "quoted false", input: `"false"`, want: false},
		{name: "one", input: `1`, want: true},
		{name: "zero", input: `0`, want: false},
		{name: "negative integer", input: `-3`, want: true},
		{name: "quoted one", input: `"1"`, want: true},
		{name: "quoted zero", input: `"0"`, want: false},
		{name: "null", input: `null`, wantNil: true},
		{name: "empty string", input: `""`, wantNil: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got AdaptiveBool
			if err := json.Unmarshal([]byte(tt.input), &got); err != nil {
				t.Fatalf("json.Unmarshal(%s) error = %v", tt.input, err)
			}

			gotPtr := got.BoolPtr()
			if tt.wantNil {
				if gotPtr != nil {
					t.Fatalf("BoolPtr() = %v, want nil", *gotPtr)
				}
				return
			}
			if gotPtr == nil {
				t.Fatalf("BoolPtr() = nil, want %v", tt.want)
			}
			if *gotPtr != tt.want {
				t.Fatalf("BoolPtr() = %v, want %v", *gotPtr, tt.want)
			}
		})
	}
}

func TestAdaptiveBoolUnmarshalInvalidValue(t *testing.T) {
	var got AdaptiveBool
	if err := json.Unmarshal([]byte(`"yes"`), &got); err == nil {
		t.Fatalf("json.Unmarshal invalid AdaptiveBool should return error")
	}
}

func TestAdaptiveBoolAccessors(t *testing.T) {
	var trueValue AdaptiveBool
	if err := json.Unmarshal([]byte(`true`), &trueValue); err != nil {
		t.Fatalf("json.Unmarshal true error = %v", err)
	}
	if !trueValue.Bool() {
		t.Fatalf("Bool() = false, want true")
	}
	if got := trueValue.Int64(); got != 1 {
		t.Fatalf("Int64() = %d, want 1", got)
	}
	if got := trueValue.Int64Ptr(); got == nil || *got != 1 {
		t.Fatalf("Int64Ptr() = %v, want 1", got)
	}

	var falseValue AdaptiveBool
	if err := json.Unmarshal([]byte(`0`), &falseValue); err != nil {
		t.Fatalf("json.Unmarshal false error = %v", err)
	}
	if falseValue.Bool() {
		t.Fatalf("Bool() = true, want false")
	}
	if got := falseValue.Int64(); got != 0 {
		t.Fatalf("Int64() = %d, want 0", got)
	}
	if got := falseValue.Int64Ptr(); got == nil || *got != 0 {
		t.Fatalf("Int64Ptr() = %v, want 0", got)
	}

	var empty AdaptiveBool
	if err := json.Unmarshal([]byte(`null`), &empty); err != nil {
		t.Fatalf("json.Unmarshal null error = %v", err)
	}
	if empty.Bool() {
		t.Fatalf("empty Bool() = true, want false")
	}
	if got := empty.BoolPtr(); got != nil {
		t.Fatalf("empty BoolPtr() = %v, want nil", *got)
	}
	if got := empty.Int64(); got != 0 {
		t.Fatalf("empty Int64() = %d, want 0", got)
	}
	if got := empty.Int64Ptr(); got != nil {
		t.Fatalf("empty Int64Ptr() = %v, want nil", *got)
	}
}

func TestAdaptiveBoolMarshalJSON(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "true", input: `"1"`, want: `true`},
		{name: "false", input: `false`, want: `false`},
		{name: "null", input: `""`, want: `null`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var value AdaptiveBool
			if err := json.Unmarshal([]byte(tt.input), &value); err != nil {
				t.Fatalf("json.Unmarshal(%s) error = %v", tt.input, err)
			}
			got, err := json.Marshal(value)
			if err != nil {
				t.Fatalf("json.Marshal error = %v", err)
			}
			if string(got) != tt.want {
				t.Fatalf("json.Marshal = %s, want %s", got, tt.want)
			}
		})
	}
}
