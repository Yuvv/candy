package jsontypes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

// AdaptiveBool unmarshals JSON booleans from bool, integer, and string forms.
type AdaptiveBool struct {
	value *bool
}

// UnmarshalJSON accepts true, false, quoted boolean strings, integers, quoted
// integers, null, and an empty string.
func (b *AdaptiveBool) UnmarshalJSON(data []byte) error {
	data = bytes.TrimSpace(data)
	if bytes.Equal(data, []byte("null")) {
		b.value = nil
		return nil
	}

	var text string
	if len(data) > 0 && data[0] == '"' {
		if err := json.Unmarshal(data, &text); err != nil {
			return err
		}
		if text == "" {
			b.value = nil
			return nil
		}
	} else {
		text = string(data)
	}

	switch text {
	case "true":
		return b.set(true)
	case "false":
		return b.set(false)
	}

	intValue, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		return fmt.Errorf("jsontypes: invalid AdaptiveBool value %q", text)
	}
	return b.set(intValue != 0)
}

// MarshalJSON returns true, false, or null.
func (b AdaptiveBool) MarshalJSON() ([]byte, error) {
	if b.value == nil {
		return []byte("null"), nil
	}
	if *b.value {
		return []byte("true"), nil
	}
	return []byte("false"), nil
}

// Bool returns the stored value, or false when empty.
func (b AdaptiveBool) Bool() bool {
	if b.value == nil {
		return false
	}
	return *b.value
}

// BoolPtr returns a pointer to the stored value, or nil when empty.
func (b AdaptiveBool) BoolPtr() *bool {
	if b.value == nil {
		return nil
	}
	value := *b.value
	return &value
}

// Int64 returns 1 for true, and 0 for false or empty.
func (b AdaptiveBool) Int64() int64 {
	if b.Bool() {
		return 1
	}
	return 0
}

// Int64Ptr returns a pointer to 1 for true or 0 for false, or nil when empty.
func (b AdaptiveBool) Int64Ptr() *int64 {
	if b.value == nil {
		return nil
	}
	value := b.Int64()
	return &value
}

func (b *AdaptiveBool) set(value bool) error {
	b.value = &value
	return nil
}
