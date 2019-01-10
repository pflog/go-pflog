# High Performance Zero Alloc bytes.Buffer

This package provides a pooled `bytes.Buffer` with zero memory
allocation for high performance applications.

`mrcrgl/bytesf`'s Pooled Buffer is ~3.5x faster compared to reallocation.
It's built for cases where the buffer is high frequent and just
temporary used.

Example cases are binary operations (API's, Binfiles, File operations, ...)

## Usage

```go
package main

import (
	"fmt"

	"github.com/mrcrgl/bytesf"
)

func main() {
	// Setup default buffer size of 64 bytes
	rb := bytesf.NewBufferPool(64, 256)

	// Allocate buffer
	b := rb.Allocate()

	b.WriteString("Hello World!")
	fmt.Printf("% x", b.Bytes())

	// Release ownership of buffer
	rb.Release(b)
}
```

## Benchmark

```
goos: darwin
goarch: amd64
pkg: github.com/mrcrgl/bytesf
Benchmark_NotPooled/Text_0_(76_bytes)-4         	10000000	       143 ns/op	     240 B/op	       2 allocs/op
Benchmark_NotPooled/Text_1_(127_bytes)-4        	10000000	       142 ns/op	     240 B/op	       2 allocs/op
Benchmark_NotPooled/Text_2_(162_bytes)-4        	 5000000	       282 ns/op	     688 B/op	       3 allocs/op
Benchmark_NotPooled/Text_3_(128_bytes)-4        	10000000	       142 ns/op	     240 B/op	       2 allocs/op
Benchmark_NotPooled/Text_4_(220_bytes)-4        	 5000000	       289 ns/op	     720 B/op	       3 allocs/op
Benchmark_NotPooled/Text_5_(59_bytes)-4         	10000000	       141 ns/op	     240 B/op	       2 allocs/op
Benchmark_Pooled/Text_0_(76_bytes)-4            	50000000	        38.5 ns/op	       0 B/op	       0 allocs/op
Benchmark_Pooled/Text_1_(127_bytes)-4           	50000000	        37.4 ns/op	       0 B/op	       0 allocs/op
Benchmark_Pooled/Text_2_(162_bytes)-4           	30000000	        40.6 ns/op	       0 B/op	       0 allocs/op
Benchmark_Pooled/Text_3_(128_bytes)-4           	50000000	        37.0 ns/op	       0 B/op	       0 allocs/op
Benchmark_Pooled/Text_4_(220_bytes)-4           	30000000	        40.3 ns/op	       0 B/op	       0 allocs/op
Benchmark_Pooled/Text_5_(59_bytes)-4            	50000000	        35.7 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/mrcrgl/bytesf	19.897s
```