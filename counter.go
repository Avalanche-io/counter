package counter

import "sync/atomic"

type Counter int64

func New() *Counter {
	return new(Counter)
}

func (c *Counter) Add(val int64) int64 {
	return atomic.AddInt64((*int64)(c), val)
}

func (c *Counter) Up() int64 {
	return atomic.AddInt64((*int64)(c), 1)
}

func (c *Counter) Down() int64 {
	for v := c.Get(); v > 0; v = c.Get() {
		if atomic.CompareAndSwapInt64((*int64)(c), v, v-1) {
			return v - 1
		}
	}
	return 0
}

func (c *Counter) Subtract(val int64) int64 {
	for v := c.Get(); (v - val) >= 0; v = c.Get() {
		if atomic.CompareAndSwapInt64((*int64)(c), v, v-val) {
			return v - val
		}
	}
	return 0
}

func (c *Counter) Set(v int64) {
	atomic.StoreInt64((*int64)(c), v)
}

func (c *Counter) Get() int64 {
	return atomic.LoadInt64((*int64)(c))
}
