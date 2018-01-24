package counter

import "sync/atomic"

type UnsignedCounter uint64

func NewUnsigned() *UnsignedCounter {
	return new(UnsignedCounter)
}

func (c *UnsignedCounter) Add(val uint64) uint64 {
	return atomic.AddUint64((*uint64)(c), val)
}

func (c *UnsignedCounter) Up() uint64 {
	return atomic.AddUint64((*uint64)(c), 1)
}

func (c *UnsignedCounter) Down() uint64 {
	for v := c.Get(); v > 0; v = c.Get() {
		if atomic.CompareAndSwapUint64((*uint64)(c), v, v-1) {
			return v - 1
		}
	}
	return 0
}

func (c *UnsignedCounter) Subtract(val uint64) uint64 {
	for v := c.Get(); (v - val) >= 0; v = c.Get() {
		if atomic.CompareAndSwapUint64((*uint64)(c), v, v-val) {
			return v - val
		}
	}
	return 0
}

func (c *UnsignedCounter) Set(v uint64) {
	atomic.StoreUint64((*uint64)(c), v)
}

func (c *UnsignedCounter) Get() uint64 {
	return atomic.LoadUint64((*uint64)(c))
}
