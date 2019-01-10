package container

import (
	"io"

	"github.com/satori/go.uuid"
)

var ctxContainerKind = []byte("ctx")

func NewCtx(uid string, scope string) *Ctx {
	u, _ := uuid.FromString(uid)

	return &Ctx{
		UUID:  u,
		Scope: []byte(scope),
	}
}

type Ctx struct {
	UUID  uuid.UUID
	Scope []byte
}

func (m *Ctx) WriteTextTo(writer io.Writer) (n int, err error) {
	if m.UUID == uuid.Nil {
		return
	}

	var i int
	n, err = writer.Write([]byte(m.UUID.String()))
	if err != nil {
		return
	}

	if len(m.Scope) == 0 {
		return
	}

	i, err = writer.Write([]byte(" "))
	n += i
	if err != nil {
		return
	}

	i, err = writer.Write(m.Scope)
	n += i
	if err != nil {
		return
	}

	return
}

func (m *Ctx) ReadTextFrom(reader io.Reader) (n int, err error) {
	b := make([]byte, 36)
	n, err = reader.Read(b)
	if err != nil {
		return
	}

	if n == 0 {
		return
	}

	m.UUID = uuid.FromStringOrNil(string(b))

	return
}

func (_ Ctx) Enclosed() bool {
	return true
}

func (_ Ctx) Kind() []byte {
	return ctxContainerKind
}
