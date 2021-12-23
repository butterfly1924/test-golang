package main

import (
	"fmt"
	"runtime"
)

func main() {
	var mem runtime.MemStats

	runtime.ReadMemStats(&mem)

	// print mem stats with alloc, totalalloc, heapalloc, numgc
	fmt.Println("mem.Alloc:", mem.Alloc)
	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc:", mem.HeapAlloc)
	fmt.Println("mem.NumGC:", mem.NumGC)
	fmt.Println("-----")
}
