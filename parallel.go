package phash

import (
	"runtime"
	"sync"
	"sync/atomic"
)

// parallel starts parallel image processing based on the current GOMAXPROCS value.
// If GOMAXPROCS = 1 it uses no parallelization.
// If GOMAXPROCS > 1 it spawns N=GOMAXPROCS workers in separate goroutines.
// this is a copy pasta from github.com/disintegration/imaging
func parallel(dataSize int, fn func(partStart, partEnd int)) {
	numGoroutines := 1
	partSize := dataSize

	numProcs := runtime.GOMAXPROCS(0)
	if numProcs > 1 {
		numGoroutines = numProcs - 1
		partSize = dataSize / (numGoroutines)
		if partSize < 1 {
			partSize = 1
		}
	}

	if numGoroutines == 1 {
		fn(0, dataSize)
	} else {
		var wg sync.WaitGroup
		wg.Add(numGoroutines)
		idx := uint64(0)

		for p := 0; p < numGoroutines; p++ {
			go func(p int) {
				defer wg.Done()
				for {
					partStart := int(atomic.AddUint64(&idx, uint64(partSize))) - partSize
					if partStart >= dataSize {
						break
					}
					partEnd := partStart + partSize
					if partEnd > dataSize {
						partEnd = dataSize
					}
					fn(partStart, partEnd)
				}
			}(p)
		}

		wg.Wait()
	}
}
