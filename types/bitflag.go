package types

type unsignedFlag interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type BitFlag[T unsignedFlag] struct {
	value T
}

func (f *BitFlag[T]) Set(flag T) {
	f.value |= flag
}

func (f *BitFlag[T]) Clear(flag T) {
	f.value &^= flag
}

func (f *BitFlag[T]) Toggle(flag T) {
	f.value ^= flag
}

func (f BitFlag[T]) Has(flag T) bool {
	return f.value&flag == flag
}

func (f BitFlag[T]) Value() T {
	return f.value
}
