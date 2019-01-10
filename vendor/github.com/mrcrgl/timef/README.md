# High Performance Time Formatter

This package is built for highly optimized applications and packages
like loggers. The API allows you to preallocate the required memory,
this in turn leads to zero memory allocations within the formatter.

`timef` is up to two times faster compared to `time.Time.Format`.


## Usage

```go
import (
	"bytes"
	"fmt"
	"time"

	"github.com/mrcrgl/timef"
)

func toString() {
	f, _ := timef.Format(time.RFC3339, time.Now())

	fmt.Println(f)
	// Output: "2018-12-07T14:21:09+01:00"
}

func toBytes() {
	b := new(bytes.Buffer)
	b.WriteString("T: ")

	n, _ := b.Write(timef.FormatRFC3339(time.Now()))

	b.WriteString(fmt.Sprintf(" is %d bytes long", n))

	fmt.Println(b.String())
	// Output: "T: 2018-12-07T14:21:09+01:00 is 25 bytes long"
}

func writeAtBytes() {
	b := make([]byte, 64)

	var pos int
	off := copy(b, []byte("T: "))
	pos += off

	n, _ := timef.WriteRFC3339At(time.Now(), b, int64(off))
	pos += n

	n = copy(b[n+off:], []byte(fmt.Sprintf(" is %d bytes long", n)))
	pos += n

	fmt.Println(string(b[0:pos]))
	// Output: "T: 2018-12-07T14:31:10+01:00 is 25 bytes long"
}
```

## Benchmark

```
goos: darwin
goarch: amd64
pkg: github.com/mrcrgl/timef
Benchmark_Format-4          	10000000	       166 ns/op	      64 B/op	       2 allocs/op
Benchmark_FormatBytes-4     	10000000	       131 ns/op	      32 B/op	       1 allocs/op
Benchmark_FormatRFC3339-4   	10000000	       126 ns/op	      32 B/op	       1 allocs/op
Benchmark_WriteRFC3339-4    	20000000	       100 ns/op	       0 B/op	       0 allocs/op
Benchmark_TimeFormat-4      	10000000	       225 ns/op	      32 B/op	       1 allocs/op
```