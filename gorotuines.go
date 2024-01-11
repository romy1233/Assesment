package main

import (
	"fmt"
	"sync"
	"time"
)

var sharedBuffer = make([]byte, 1024)
var readMutex sync.RWMutex
var writeMutex sync.RWMutex

func reader() {
	for {
		readLock(&readMutex)
		fmt.Println("Reader read from buffer")
		readMutex.Unlock()
		time.Sleep(time.Millisecond * 100)
	}
}

func writer() {
	for {
		writeLock(&writeMutex)
		fmt.Println("Writer wrote to buffer")
		writeMutex.Unlock()
		time.Sleep(time.Millisecond * 200)
	}
}

func readLock(m *sync.RWMutex) {
	m.RLock()
}

func writeLock(m *sync.RWMutex) {
	m.Lock()
}

func main() {
	for i := 0; i < 8; i++ {
		go reader()
	}
	for i := 0; i < 2; i++ {
		go writer()
	}
	time.Sleep(time.Second * 5)
}