package go_pool

import (
	"errors"
	"math"
	"runtime"
)

const (
	// DefaultAntsPoolSize is the default capacity for a default goroutine pool.
	DefaultAntsPoolSize = math.MaxInt32

	// DefaultCleanIntervalTime is the interval time to clean up goroutines.
	DefaultCleanIntervalTime = 1
)

var (
	// Error types for the Ants API.
	//---------------------------------------------------------------------------
	// ErrInvalidPoolSize will be returned when setting a negative number as pool capacity.
	ErrInvalidPoolSize = errors.New("invalid size for pool")

	// ErrInvalidPoolExpiry will be returned when setting a negative number as the periodic duration to purge goroutines.
	ErrInvalidPoolExpiry = errors.New("invalid expiry for pool")

	// ErrPoolClosed will be returned when submitting task to a closed pool.
	ErrPoolClosed = errors.New("this pool has been closed")
	//---------------------------------------------------------------------------

	// workerChanCap determines whether the channel of a worker should be a buffered channel
	// to get the best performance. Inspired by fasthttp at https://github.com/valyala/fasthttp/blob/master/workerpool.go#L139
	workerChanCap = func() int {
		// Use blocking workerChan if GOMAXPROCS=1.
		// This immediately switches Serve to WorkerFunc, which results
		// in higher performance (under go1.5 at least).
		if runtime.GOMAXPROCS(0) == 1 {
			return 0
		}

		// Use non-blocking workerChan if GOMAXPROCS>1,
		// since otherwise the Serve caller (Acceptor) may lag accepting
		// new connections if WorkerFunc is CPU-bound.
		return 1
	}()

	defaultAntsPool, _ = NewPool(DefaultAntsPoolSize)
)

// Init a instance pool when importing ants.

// Submit submits a task to pool.
func Submit(task func()) error {
	return defaultAntsPool.Submit(task)
}

// Running returns the number of the currently running goroutines.
func Running() int {
	return defaultAntsPool.Running()
}

// Cap returns the capacity of this default pool.
func Cap() int {
	return defaultAntsPool.Cap()
}

// Free returns the available goroutines to work.
func Free() int {
	return defaultAntsPool.Free()
}

// Release Closes the default pool.
func Release() {
	_ = defaultAntsPool.Release()
}
