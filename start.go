package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
)

func main() {
	// create file
	f, err := os.Create("/tmp/traceFile.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer trace.Stop()

	for i := 0; i < 10; i++ {
		s := make([]byte, 50000000)
		if s == nil {
			fmt.Println("Operation failed")
		}
	}

	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	// print mem stats with alloc, totalalloc, heapalloc, numgc
	ReadMemStats(&mem)
}

func ReadMemStats(mem *runtime.MemStats) {
	fmt.Println("mem.Alloc:", mem.Alloc)
	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc:", mem.HeapAlloc)
	fmt.Println("mem.NumGC:", mem.NumGC)
	fmt.Println("-----")
}
