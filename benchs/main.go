package main

import (
	"crypto/rand"
	"fmt"
	"os"
	"runtime/pprof"
	"time"

	"github.com/gballet/go-verkle"
)

func main() {
	// benchmarkInsertInExisting()
	tempWay()
}

func benchmarkInsertInExisting() {
	// f, _ := os.Create("cpu.prof")
	// g, _ := os.Create("mem.prof")
	// pprof.StartCPUProfile(f)
	// defer pprof.StopCPUProfile()
	// defer pprof.WriteHeapProfile(g)
	// Number of existing leaves in tree
	n := 1000
	// Leaves to be inserted afterwards
	toInsert := 100
	total := n + toInsert

	keys := make([][]byte, n)
	toInsertKeys := make([][]byte, toInsert)
	value := []byte("value")

	for i := 0; i < 4; i++ {
		// Generate set of keys once
		for i := 0; i < total; i++ {
			key := make([]byte, 32)
			rand.Read(key)
			if i < n {
				keys[i] = key
			} else {
				toInsertKeys[i-n] = key
			}
		}
		fmt.Printf("Generated key set %d\n", i)

		// Create tree from same keys multiple times
		for i := 0; i < 5; i++ {
			root := verkle.New()
			for _, k := range keys {
				if err := root.Insert(k, value, nil); err != nil {
					panic(err)
				}
			}
			fmt.Println("Computing root commitment")
			root.ComputeCommitment()

			// Now insert the 10k leaves and measure time
			start := time.Now()
			for _, k := range toInsertKeys {
				if err := root.Insert(k, value, nil); err != nil {
					panic(err)
				}
			}
			root.ComputeCommitment()
			elapsed := time.Since(start)
			fmt.Printf("Took %v to insert and commit %d leaves\n", elapsed, toInsert)
		}
	}
}

func tempWay() {
	f, _ := os.Create("cpu.prof")
	g, _ := os.Create("mem.prof")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	defer pprof.WriteHeapProfile(g)
	fmt.Println("Temp way begin ----")

	startTime := time.Now()
	root := verkle.New()
	fmt.Println("create root took ", time.Now().Sub(startTime))

	startTime = time.Now()
	cmt := root.ComputeCommitment()
	fmt.Println("empty root commitment gen took: ", time.Now().Sub(startTime), " com: ", cmt)

	// for _, toInsert := range []int{1, 50, 100, 10000} {
	// 	toInsertKeys := make([][]byte, toInsert)
	// 	value := []byte("value")
	// 	for i := 0; i < toInsert; i++ {
	// 		key := make([]byte, 32)
	// 		rand.Read(key)
	// 		toInsertKeys[i] = key
	// 	}

	// 	// fmt.Println("toInsertKeys: ", toInsertKeys)

	// 	start := time.Now()
	// 	for _, k := range toInsertKeys {
	// 		if err := root.Insert(k, value, nil); err != nil {
	// 			fmt.Println("error ", err, " k ", k, " value ", value)
	// 			panic(err)
	// 		}
	// 	}
	// 	// fmt.Println("root ", root)
	// 	root.ComputeCommitment()
	// 	elapsed := time.Since(start)
	// 	fmt.Printf("Took %v to insert and commit %d leaves\n", elapsed, toInsert)
	// }

	fmt.Println("Temp way end ----")
}
