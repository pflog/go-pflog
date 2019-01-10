package text

import (
	"bytes"

	"github.com/mrcrgl/pflog/pkg/logging"

	"github.com/mrcrgl/bytesf"
	"github.com/mrcrgl/timef"
)

type severityStringRep []byte

var (
	infoStringRep    severityStringRep = []byte("I")
	warningStringRep severityStringRep = []byte("W")
	errorStringRep   severityStringRep = []byte("E")
	fatalStringRep   severityStringRep = []byte("F")

	space                 []byte = []byte(" ")
	containerContentBegin []byte = []byte("{")
	containerContentEnd   []byte = []byte("}")
	lineEnd               []byte = []byte(";\n")
)

func NewEncoder() *encoder {
	return &encoder{
		bp: bytesf.NewBufferPool(128, 512),
	}
}

type encoder struct {
	bp bytesf.BufferPool
}

func (e *encoder) Encode(in *logging.Entry) ([]byte, error) {
	b := e.bp.Allocate()
	defer e.bp.Release(b)

	err := encode(in, b)

	return b.Bytes(), err
}

func Encode(in *logging.Entry) ([]byte, error) {
	bs := make([]byte, 26, 256)
	b := bytes.NewBuffer(bs)
	b.Reset()

	err := encode(in, b)

	return b.Bytes(), err
}

func encode(in *logging.Entry, b *bytes.Buffer) (err error) {
	switch in.Severity {
	case logging.SeverityInfo:
		b.Write(infoStringRep)
		break
	case logging.SeverityWarning:
		b.Write(warningStringRep)
		break
	case logging.SeverityError:
		b.Write(errorStringRep)
		break
	case logging.SeverityFatal:
		b.Write(fatalStringRep)
		break
	default:
		b.WriteByte('?')
	}

	//fmt.Printf("len=%d\n", b.Len())

	//b.WriteString(in.Timestamp.Format(time.RFC3339))
	b.Write(timef.FormatRFC3339(in.Timestamp))

	//fmt.Printf("len=%d\n", b.Len())

	for _, c := range in.Containers {
		b.Write(space)

		if c.Enclosed() {
			b.Write(c.Kind())
			b.Write(containerContentBegin)
		}

		_, err := c.WriteTextTo(b)
		if err != nil {
			return err
		}

		if c.Enclosed() {
			b.Write(containerContentEnd)
		}

		//fmt.Printf("container=%d len=%d\n", i, b.Len())
	}

	b.Write(lineEnd)

	return nil
}
