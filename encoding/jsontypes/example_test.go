package jsontypes_test

import (
	"encoding/json"
	"fmt"

	"github.com/yuvv/candy/encoding/jsontypes"
)

func Example() {
	var payload struct {
		Enabled jsontypes.AdaptiveBool   `json:"enabled"`
		Count   jsontypes.AdaptiveNumber `json:"count"`
	}
	if err := json.Unmarshal([]byte(`{"enabled":"true","count":"42"}`), &payload); err != nil {
		panic(err)
	}

	fmt.Println(payload.Enabled.Bool(), payload.Count.Int64())
	// Output: true 42
}
