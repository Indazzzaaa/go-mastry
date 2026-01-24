package main

import (
	"fmt"
	"runtime"
	"time"
)

func allocate() {
	var data [][]byte
	for i := 0; i < 1000; i++ {
		b := make([]byte, 1024*1024) // 1MB
		data = append(data, b)
	}
}

func printMem() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("HeapAlloc = %d MB\n", m.HeapAlloc/1024/1024)
}

func main() {

	printMem()

	allocate()

	printMem()

	runtime.GC() // force GC

	time.Sleep(500 * time.Millisecond)

	printMem()
}
