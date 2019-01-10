package text

import (
	"testing"
	"time"

	"github.com/pflog/go-pflog/container"
	"github.com/pflog/go-pflog/pkg/logging"
)

func BenchmarkEncode(b *testing.B) {

	e := new(logging.Entry)
	e.Severity = logging.SeverityInfo
	e.Timestamp = time.Now()
	e.Containers = []logging.Container{
		//container.NewCtx("f6d4bba8-e8b2-4cc1-b7e5-d9cf8fa2ebc1", ""),
		container.NewMessage([]byte("Hello friends!!!")),
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		Encode(e)
	}
}

func BenchmarkEncoder_Encode(b *testing.B) {
	e := new(logging.Entry)
	e.Severity = logging.SeverityInfo
	e.Timestamp = time.Now()
	e.Containers = []logging.Container{
		//container.NewCtx("f6d4bba8-e8b2-4cc1-b7e5-d9cf8fa2ebc1", ""),
		container.NewMessage([]byte("Hello friends!!!")),
	}

	b.ResetTimer()
	encoder := NewEncoder()
	for n := 0; n < b.N; n++ {
		encoder.Encode(e)
	}
}
