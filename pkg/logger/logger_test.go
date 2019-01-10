package logger

import (
	"testing"

	"github.com/pflog/go/pkg/encoding/text"

	"github.com/pflog/go/container"
	uuid "github.com/satori/go.uuid"

	"github.com/golang/glog"
)

type noopWriter struct{}

func (_ noopWriter) Write([]byte) (int, error) {
	return 0, nil
}

func BenchmarkLogger_GlogInfo(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		glog.Info("Foobar")
	}
}

func BenchmarkLogger_GlogWithInfo(b *testing.B) {
	id := uuid.FromStringOrNil("deba2283-423c-4461-bccd-cc05e6319d8a")
	scope := "request"

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		glog.Infof("Foobar: %s %s", id, scope)
	}
}

func BenchmarkLogger_Info(b *testing.B) {
	l := logger{
		encoder: text.NewEncoder(),
		output:  noopWriter{},
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		l.Info("Foobar")
	}
}

func BenchmarkLogger_WithInfo(b *testing.B) {
	l := logger{
		encoder: text.NewEncoder(),
		output:  noopWriter{},
	}

	id := uuid.FromStringOrNil("deba2283-423c-4461-bccd-cc05e6319d8a")
	scope := "request"

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		l.With(container.NewCtx(id.String(), scope)).Info("Foobar")
	}
}

func BenchmarkLogger_WithInfo_PureStruct(b *testing.B) {
	l := logger{
		encoder: text.NewEncoder(),
		output:  noopWriter{},
	}

	id := uuid.FromStringOrNil("deba2283-423c-4461-bccd-cc05e6319d8a")
	scope := "request"

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		l.With(&container.Ctx{UUID: id, Scope: []byte(scope)}).Info("Foobar")
	}
}

func BenchmarkLogger_WithInfo_Prealloc(b *testing.B) {
	l := logger{
		encoder: text.NewEncoder(),
		output:  noopWriter{},
	}

	id := uuid.FromStringOrNil("deba2283-423c-4461-bccd-cc05e6319d8a")
	scope := "request"
	c := container.NewCtx(id.String(), scope)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		l.With(c).Info("Foobar")
	}
}

func BenchmarkLogger_WithInfo_PreparedLogger(b *testing.B) {
	l := logger{
		encoder: text.NewEncoder(),
		output:  noopWriter{},
	}

	id := uuid.FromStringOrNil("deba2283-423c-4461-bccd-cc05e6319d8a")
	scope := "request"

	l2 := l.With(container.NewCtx(id.String(), scope))

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		l2.Info("Foobar")
	}
}
