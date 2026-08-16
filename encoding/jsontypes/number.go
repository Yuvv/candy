package jsontypes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// AdaptiveNumber unmarshals JSON numbers from numeric and string forms.
type AdaptiveNumber struct {
	intValue   *int64
	floatValue *float64
}

// UnmarshalJSON accepts numbers, quoted numbers, null, and an empty string.
func (n *AdaptiveNumber) UnmarshalJSON(data []byte) error {
	data = bytes.TrimSpace(data)
	if bytes.Equal(data, []byte("null")) {
		n.clear()
		return nil
	}

	var text string
	if len(data) > 0 && data[0] == '"' {
		if err := json.Unmarshal(data, &text); err != nil {
			return err
		}
	} else {
		text = string(data)
	}

	text = strings.TrimSpace(text)
	if text == "" {
		n.clear()
		return nil
	}

	if isIntegerToken(text) {
		intValue, err := strconv.ParseInt(text, 10, 64)
		if err != nil {
			return fmt.Errorf("jsontypes: invalid AdaptiveNumber integer value %q", text)
		}
		n.setInt(intValue)
		return nil
	}

	if !json.Valid([]byte(text)) {
		return fmt.Errorf("jsontypes: invalid AdaptiveNumber value %q", text)
	}
	floatValue, err := strconv.ParseFloat(text, 64)
	if err != nil {
		return fmt.Errorf("jsontypes: invalid AdaptiveNumber value %q", text)
	}
	n.setFloat(floatValue)
	return nil
}

// MarshalJSON returns an integer, float, or null.
func (n AdaptiveNumber) MarshalJSON() ([]byte, error) {
	if n.intValue != nil {
		return []byte(strconv.FormatInt(*n.intValue, 10)), nil
	}
	if n.floatValue != nil {
		return []byte(strconv.FormatFloat(*n.floatValue, 'f', -1, 64)), nil
	}
	return []byte("null"), nil
}

// String returns the stored number as a string, or an empty string when empty.
func (n AdaptiveNumber) String() string {
	if n.intValue != nil {
		return strconv.FormatInt(*n.intValue, 10)
	}
	if n.floatValue != nil {
		return strconv.FormatFloat(*n.floatValue, 'f', -1, 64)
	}
	return ""
}

// StringPtr returns a pointer to String, or nil when empty.
func (n AdaptiveNumber) StringPtr() *string {
	if n.intValue == nil && n.floatValue == nil {
		return nil
	}
	value := n.String()
	return &value
}

// Int64 returns the stored integer value, or 0 for float-only and empty values.
func (n AdaptiveNumber) Int64() int64 {
	if n.intValue == nil {
		return 0
	}
	return *n.intValue
}

// Int64Ptr returns a pointer to the stored integer value, or nil for float-only and empty values.
func (n AdaptiveNumber) Int64Ptr() *int64 {
	if n.intValue == nil {
		return nil
	}
	value := *n.intValue
	return &value
}

// Float64 returns the stored float value, float64(integer), or 0 when empty.
func (n AdaptiveNumber) Float64() float64 {
	if n.floatValue == nil {
		return 0
	}
	return *n.floatValue
}

// Float64Ptr returns a pointer to the float value, or nil when empty.
func (n AdaptiveNumber) Float64Ptr() *float64 {
	if n.floatValue == nil {
		return nil
	}
	value := *n.floatValue
	return &value
}

func isIntegerToken(text string) bool {
	if text == "" {
		return false
	}
	if text[0] == '-' {
		text = text[1:]
		if text == "" {
			return false
		}
	}
	for _, r := range text {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

func (n *AdaptiveNumber) clear() {
	n.intValue = nil
	n.floatValue = nil
}

func (n *AdaptiveNumber) setInt(value int64) {
	n.intValue = &value
	floatValue := float64(value)
	n.floatValue = &floatValue
}

func (n *AdaptiveNumber) setFloat(value float64) {
	n.intValue = nil
	n.floatValue = &value
}
