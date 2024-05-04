package main

import (
	"learn/memprof/function"
	"os"
	"runtime/pprof"
)

func init() {
	cpuProf()
	memProf()
}

func main() {

	if err := function.WriteFileJson(1000000); err != nil {
		panic(err)
	}
}

func cpuProf() {
	f, err := os.Create("report/functions/cpu.pprof")
	if err != nil {
		panic(err)
	}
	pprof.StartCPUProfile(f)
	defer f.Close()
	defer pprof.StopCPUProfile()
}

func memProf() {
	f, err := os.Create("report/functions/mem.pprof")
	if err != nil {
		panic(err)
	}
	pprof.WriteHeapProfile(f)
	defer f.Close()
}
