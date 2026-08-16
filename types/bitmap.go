package types

import (
	"math/bits"
	"strings"
)

// Bitmap is a dynamically growing bit set backed by uint64 words.
type Bitmap struct {
	words []uint64
}

// Set turns on bit, growing the bitmap as needed.
func (b *Bitmap) Set(bit uint) {
	word := int(bit / 64)
	if word >= len(b.words) {
		grown := make([]uint64, word+1)
		copy(grown, b.words)
		b.words = grown
	}

	b.words[word] |= 1 << (bit % 64)
}

// Unset turns off bit. Unsetting a bit outside the allocated words is a no-op.
func (b *Bitmap) Unset(bit uint) {
	word := int(bit / 64)
	if word >= len(b.words) {
		return
	}

	b.words[word] &^= 1 << (bit % 64)
}

// Get reports whether bit is set.
func (b *Bitmap) Get(bit uint) bool {
	word := int(bit / 64)
	if word >= len(b.words) {
		return false
	}

	return b.words[word]&(1<<(bit%64)) != 0
}

// Len returns the highest set bit plus one, or 0 when no bits are set.
func (b *Bitmap) Len() int {
	for i := len(b.words) - 1; i >= 0; i-- {
		if b.words[i] != 0 {
			return i*64 + bits.Len64(b.words[i])
		}
	}
	return 0
}

// Words returns the allocated word count.
func (b *Bitmap) Words() int {
	return len(b.words)
}

// String returns the bitmap bits from index 0 to Len()-1, low bit first.
func (b *Bitmap) String() string {
	length := b.Len()
	if length == 0 {
		return "b:0"
	}

	var sb strings.Builder
	sb.Grow(2 + length)
	sb.WriteString("b:")
	for bit := 0; bit < length; bit++ {
		if b.Get(uint(bit)) {
			sb.WriteByte('1')
		} else {
			sb.WriteByte('0')
		}
	}
	return sb.String()
}
