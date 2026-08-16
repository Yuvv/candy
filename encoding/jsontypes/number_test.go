package jsontypes

import (
	"encoding/json"
	"testing"
)

func TestAdaptiveNumberUnmarshalAcceptedValues(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantNil   bool
		wantInt   *int64
		wantFloat *float64
	}{
		{name: "integer", input: `123`, wantInt: int64Ptr(123), wantFloat: float64Ptr(123)},
		{name: "negative integer", input: `-45`, wantInt: int64Ptr(-45), wantFloat: float64Ptr(-45)},
		{name: "float", input: `123.5`, wantFloat: float64Ptr(123.5)},
		{name: "quoted integer", input: `"123"`, wantInt: int64Ptr(123), wantFloat: float64Ptr(123)},
		{name: "quoted float", input: `"123.5"`, wantFloat: float64Ptr(123.5)},
		{name: "null", input: `null`, wantNil: true},
		{name: "empty string", input: `""`, wantNil: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got AdaptiveNumber
			if err := json.Unmarshal([]byte(tt.input), &got); err != nil {
				t.Fatalf("json.Unmarshal(%s) error = %v", tt.input, err)
			}

			if tt.wantNil {
				if got.Int64Ptr() != nil {
					t.Fatalf("Int64Ptr() = %v, want nil", *got.Int64Ptr())
				}
				if got.Float64Ptr() != nil {
					t.Fatalf("Float64Ptr() = %v, want nil", *got.Float64Ptr())
				}
				return
			}
			if tt.wantInt == nil {
				if got.Int64Ptr() != nil {
					t.Fatalf("Int64Ptr() = %v, want nil", *got.Int64Ptr())
				}
			} else if got.Int64Ptr() == nil || *got.Int64Ptr() != *tt.wantInt {
				t.Fatalf("Int64Ptr() = %v, want %v", got.Int64Ptr(), *tt.wantInt)
			}
			if tt.wantFloat == nil {
				if got.Float64Ptr() != nil {
					t.Fatalf("Float64Ptr() = %v, want nil", *got.Float64Ptr())
				}
			} else if got.Float64Ptr() == nil || *got.Float64Ptr() != *tt.wantFloat {
				t.Fatalf("Float64Ptr() = %v, want %v", got.Float64Ptr(), *tt.wantFloat)
			}
		})
	}
}

func TestAdaptiveNumberUnmarshalInvalidValue(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{name: "text", input: `"abc"`},
		{name: "quoted nan", input: `"NaN"`},
		{name: "quoted infinity", input: `"Infinity"`},
		{name: "quoted hex float", input: `"0x1p2"`},
		{name: "overflow integer", input: `9223372036854775808`},
		{name: "quoted overflow integer", input: `"9223372036854775808"`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got AdaptiveNumber
			if err := json.Unmarshal([]byte(tt.input), &got); err == nil {
				t.Fatalf("json.Unmarshal(%s) should return error", tt.input)
			}
		})
	}
}

func TestAdaptiveNumberAccessors(t *testing.T) {
	var intValue AdaptiveNumber
	if err := json.Unmarshal([]byte(`123`), &intValue); err != nil {
		t.Fatalf("json.Unmarshal int error = %v", err)
	}
	if got := intValue.String(); got != "123" {
		t.Fatalf("String() = %q, want 123", got)
	}
	if got := intValue.StringPtr(); got == nil || *got != "123" {
		t.Fatalf("StringPtr() = %v, want 123", got)
	}
	if got := intValue.Int64(); got != 123 {
		t.Fatalf("Int64() = %d, want 123", got)
	}
	if got := intValue.Int64Ptr(); got == nil || *got != 123 {
		t.Fatalf("Int64Ptr() = %v, want 123", got)
	}
	if got := intValue.Float64(); got != 123 {
		t.Fatalf("Float64() = %v, want 123", got)
	}
	if got := intValue.Float64Ptr(); got == nil || *got != 123 {
		t.Fatalf("Float64Ptr() = %v, want 123", got)
	}

	var floatValue AdaptiveNumber
	if err := json.Unmarshal([]byte(`123.5`), &floatValue); err != nil {
		t.Fatalf("json.Unmarshal float error = %v", err)
	}
	if got := floatValue.String(); got != "123.5" {
		t.Fatalf("String() = %q, want 123.5", got)
	}
	if got := floatValue.Int64(); got != 0 {
		t.Fatalf("float-only Int64() = %d, want 0", got)
	}
	if got := floatValue.Int64Ptr(); got != nil {
		t.Fatalf("float-only Int64Ptr() = %v, want nil", *got)
	}
	if got := floatValue.Float64(); got != 123.5 {
		t.Fatalf("Float64() = %v, want 123.5", got)
	}
	if got := floatValue.Float64Ptr(); got == nil || *got != 123.5 {
		t.Fatalf("Float64Ptr() = %v, want 123.5", got)
	}

	var empty AdaptiveNumber
	if err := json.Unmarshal([]byte(`null`), &empty); err != nil {
		t.Fatalf("json.Unmarshal null error = %v", err)
	}
	if got := empty.String(); got != "" {
		t.Fatalf("empty String() = %q, want empty", got)
	}
	if got := empty.StringPtr(); got != nil {
		t.Fatalf("empty StringPtr() = %v, want nil", *got)
	}
	if got := empty.Int64(); got != 0 {
		t.Fatalf("empty Int64() = %d, want 0", got)
	}
	if got := empty.Int64Ptr(); got != nil {
		t.Fatalf("empty Int64Ptr() = %v, want nil", *got)
	}
	if got := empty.Float64(); got != 0 {
		t.Fatalf("empty Float64() = %v, want 0", got)
	}
	if got := empty.Float64Ptr(); got != nil {
		t.Fatalf("empty Float64Ptr() = %v, want nil", *got)
	}
}

func TestAdaptiveNumberMarshalJSON(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "integer", input: `"123"`, want: `123`},
		{name: "float", input: `"123.5"`, want: `123.5`},
		{name: "null", input: `""`, want: `null`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var value AdaptiveNumber
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

func int64Ptr(v int64) *int64 {
	return &v
}

func float64Ptr(v float64) *float64 {
	return &v
}
