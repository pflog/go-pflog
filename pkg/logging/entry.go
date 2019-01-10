package logging

import (
	"time"
)

type Encoder interface {
	Encode(in *Entry) (out []byte, err error)
}

type Decoder interface {
	Decode(in []byte, out *Entry) (err error)
}

type Entry struct {
	Severity   Severity
	Timestamp  time.Time
	Containers []Container
}
