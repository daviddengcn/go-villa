package villa

import (
	"sync"
	"sync/atomic"
)

// Once is an object that will perform exactly one action. Different from build-in
// sync.Once, the function is defined in the Once struct.
type Once struct {
	m    sync.Mutex
	done uint32

	F func()
}

// Do calls the function f if and only if the method is being called for the
// first time with this receiver.  In other words, given
//      var once Once{F: f}
// if once.Do() is called multiple times, only the first call will invoke f.
//
// Do is intended for initialization that must be run exactly once.
//
// Because no call to Do returns until the one call to f returns, if f causes
// Do to be called, it will deadlock.
//
func (o *Once) Do() {
	if atomic.LoadUint32(&o.done) == 1 {
		return
	}
	// Slow-path.
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		o.F()
		atomic.StoreUint32(&o.done, 1)
	}
}

// AtomicBox is a place atomically storing an object(typed interface{})
type AtomicBox struct {
	sync.RWMutex
	object interface{}
}

// Get fetches the object from the box atomically
func (b *AtomicBox) Get() interface{} {
	b.RLock()
	defer b.RUnlock()

	return b.object
}

// Set stores the object into the box atomically
func (b *AtomicBox) Set(val interface{}) {
	b.Lock()
	defer b.Unlock()

	b.object = val
}
