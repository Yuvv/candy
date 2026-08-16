package strs

import (
	"fmt"
	"strings"
)

type Builder struct {
	builder strings.Builder
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) Append(value any) *Builder {
	_, _ = fmt.Fprint(&b.builder, value)
	return b
}

func (b *Builder) AppendString(value string) *Builder {
	b.builder.WriteString(value)
	return b
}

func (b *Builder) AppendByte(value byte) *Builder {
	_ = b.builder.WriteByte(value)
	return b
}

func (b *Builder) String() string {
	return b.builder.String()
}

func (b *Builder) Len() int {
	return b.builder.Len()
}
