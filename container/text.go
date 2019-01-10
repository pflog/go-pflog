package container

import (
	"io"
)

var textContainerKind = []byte("txt")

func NewText(b []byte) *Text {
	return &Text{
		b: b,
	}
}

type Text struct {
	b []byte
}

func (m *Text) WriteTextTo(writer io.Writer) (int, error) {
	return writer.Write(m.b)
}

func (m *Text) ReadTextFrom(reader io.Reader) (int, error) {
	return 0, nil
}

func (_ Text) Enclosed() bool {
	return true
}

func (_ Text) Kind() []byte {
	return textContainerKind
}
