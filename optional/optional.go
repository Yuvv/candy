package optional

import (
	"reflect"

	"github.com/yuvv/candy/function"
	"github.com/yuvv/candy/lang"
)

type Optional[T any] struct {
	// if non-null, the value;
	// if null, indicates no value is present.
	value T
	//
	nillable bool
	nilValue T
}

func (receiver *Optional[T]) Get() T {
	//if receiver.nillable && receiver.value == receiver.nilValue {
	//	panic("No value present")
	//}
	return receiver.value
}

func (receiver *Optional[T]) IsPresent() bool {
	return !receiver.nillable || !reflect.DeepEqual(receiver.value, receiver.nilValue)
}

func (receiver *Optional[T]) OrElse(orElse T) T {
	if receiver.IsPresent() {
		return receiver.value
	}
	return orElse
}

func (receiver *Optional[T]) IsEmpty() bool {
	return receiver.nillable && reflect.DeepEqual(receiver.value, receiver.nilValue)
}

func (receiver *Optional[T]) IfPresent(consumer function.Consumer[T]) {
	if receiver.IsPresent() {
		consumer(receiver.value)
	}
}

func (receiver *Optional[T]) IfPresentOrElse(consumer function.Consumer[T], emptyAction lang.Runnable) {
	if receiver.IsPresent() {
		consumer(receiver.value)
	} else {
		emptyAction.Run()
	}
}

func Empty[T any]() *Optional[T] {
	var val T
	return &Optional[T]{
		value:    val,
		nillable: true,
		nilValue: val,
	}
}

func Of[T any](t T) *Optional[T] {
	optional := OfNillable[T](t)
	if optional.nillable {
		panic("value is required non-nil")
	}
	return optional
}

func OfNillable[T any](t T) *Optional[T] {
	nullable := true
	tType := reflect.TypeOf(t)
	switch tType.Kind() {
	case reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Uintptr,
		reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128,
		reflect.String,
		reflect.Struct:
		nullable = false
		break
	}

	var nv T
	return &Optional[T]{
		value:    t,
		nillable: nullable,
		nilValue: nv,
	}
}
