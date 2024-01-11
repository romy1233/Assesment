var sharedBuffer = make([]byte, 1024)
var readMutex sync.RWMutex
var writeMutex sync.RWMutex

func reader() {
	for {
		readLock(&readMutex)
		// read from sharedBuffer
	}
}

func writer() {
	for {
		writeLock(&writeMutex)
		// write to sharedBuffer
	}
}