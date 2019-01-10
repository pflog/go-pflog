package bytesf

import (
	"sync"
)

func NewListPool(size, maxSize int) *listPool {
	return &listPool{
		size:    size,
		maxSize: maxSize,
		mu:      new(sync.Mutex),
	}
}

type listPool struct {
	size    int
	maxSize int
	mu      *sync.Mutex
	slotA   []byte
	slotB   []byte
}

func (b *listPool) Allocate() []byte {
	b.mu.Lock()
	c := b.slotA
	if c != nil {
		b.slotA = b.slotB
		b.slotB = nil
	}
	b.mu.Unlock()
	if c == nil {
		c = b.newList()
	}

	return c
}

func (b *listPool) Release(c []byte) {
	if len(c) >= b.maxSize {
		return
	}

	b.mu.Lock()
	b.slotB = b.slotA
	b.slotA = c
	b.mu.Unlock()
}

func (b *listPool) newList() []byte {
	c := make([]byte, b.size, b.maxSize)

	return c
}
