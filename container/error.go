package container

import (
	"bytes"
)

var errorContainerKind = []byte("error")

func NewError(err error, stack []byte) *Error {
	return &Error{
		Text:  NewText(bytes.NewBufferString(err.Error()).Bytes()),
		Stack: stack,
	}
}

type Error struct {
	*Text
	Stack []byte
}

func (_ Error) Enclosed() bool {
	return false
}

func (_ Error) Kind() []byte {
	return errorContainerKind
}
