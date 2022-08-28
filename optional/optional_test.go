package optional

import "testing"

func TestOptional_Get(t *testing.T) {
	t.Run("PanicForNillable-Init", func(t *testing.T) {
		defer func() {
			if err := recover(); err == nil {
				t.Errorf("should panic")
			}
		}()

		optional := Of[*int](nil)
		_ = optional.Get()
	})

	t.Run("GetNillableValue", func(t *testing.T) {
		optional := OfNillable[*int](nil)
		res := optional.Get()
		if res != nil {
			t.Fatalf("should return nil, got:%+v", res)
		}

		optional2 := OfNillable[*string](&[]string{"abc"}[0])
		res2 := optional2.Get()
		if res2 == nil || *res2 != "abc" {
			t.Fatalf("should return `abc`, got:%+v", res2)
		}
	})

	t.Run("GetNonNillableValue", func(t *testing.T) {
		optional := Of[int](1)
		res := optional.Get()
		if res != 1 {
			t.Fatalf("should return 1, got:%d", res)
		}

		optional2 := Of[float64](1.3e7)
		res2 := optional2.Get()
		if res2 != 1.3e7 {
			t.Fatalf("should return 1.3e7, got:%f", res2)
		}
	})
}

func TestOptional_IsEmpty(t *testing.T) {
	t.Run("NillableEmpty", func(t *testing.T) {
		optional := Empty[*int]()
		if !optional.IsEmpty() {
			t.Fatalf("Empty Optional should be empty")
		}

		optional2 := Empty[[]string]()
		if !optional2.IsEmpty() {
			t.Fatalf("Empty Optional should be empty")
		}
	})

	t.Run("NonNillableEmpty", func(t *testing.T) {
		optional := Empty[int64]()
		if !optional.IsEmpty() {
			t.Fatalf("Empty Optional should be empty")
		}

		optional2 := Empty[bool]()
		if !optional2.IsEmpty() {
			t.Fatalf("Empty Optional should be empty")
		}
	})
}

func TestOptional_OrElse(t *testing.T) {
	t.Run("NillableOrElse", func(t *testing.T) {
		optional := OfNillable[[]int](nil)
		res := optional.OrElse([]int{1})
		if len(res) != 1 || res[0] != 1 {
			t.Fatalf("should return [1], got: %+v", res)
		}

		optional2 := OfNillable[[]string]([]string{"1-1"})
		res2 := optional2.OrElse([]string{"test"})
		if len(res2) != 1 || res2[0] != "1-1" {
			t.Fatalf("should return \"1-1\", got: %s", res2)
		}
	})
}
