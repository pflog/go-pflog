package bytesf

import (
	"bytes"
	"sync"
)

type BufferPool interface {
	Allocate() *bytes.Buffer
	Release(*bytes.Buffer)
}

func NewBufferPool(size, maxSize int) *bufferPool {
	return &bufferPool{
		size:    size,
		maxSize: maxSize,
		mu:      new(sync.Mutex),
	}
}

type bufferPool struct {
	size    int
	maxSize int
	mu      *sync.Mutex
	slotA   *bytes.Buffer
	slotB   *bytes.Buffer
}

func (b *bufferPool) Allocate() *bytes.Buffer {
	b.mu.Lock()
	c := b.slotA
	if c != nil {
		b.slotA = b.slotB
		b.slotB = nil
	}
	b.mu.Unlock()
	if c == nil {
		c = b.newBuffer()
	} else {
		c.Reset()
	}

	return c
}

func (b *bufferPool) Release(c *bytes.Buffer) {
	if c.Len() >= b.maxSize {
		return
	}

	c.Reset()

	b.mu.Lock()
	b.slotB = b.slotA
	b.slotA = c
	b.mu.Unlock()
}

func (b *bufferPool) newBuffer() *bytes.Buffer {
	c := new(bytes.Buffer)
	c.Grow(b.size)
	c.Reset()

	return c
}
