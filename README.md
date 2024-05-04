# golang pprof
## with unit test
```sh
# generate report with unit test
go test -v -run=TestFunctionName -cpuprofile report/test/cpu.prof -memprofile report/test/mem.prof
# benchmark
go test -bench=BenchName -benchmem -cpuprofile report/cpu.prof -memprofile report/mem.prof -benchtime=5s
```

##  without unit test 
- cpucprof
```go
	f, err := os.Create("report/functions/cpu.pprof")
	if err != nil {
		panic(err)
	}
	pprof.StartCPUProfile(f)
	defer f.Close()
	defer pprof.StopCPUProfile()
```

- memprof
```go
	f, err := os.Create("report/functions/mem.pprof")
	if err != nil {
		panic(err)
	}
	pprof.WriteHeapProfile(f)
	defer f.Close()
```

```sh
# see the report,
go tool pprof -http=:8080 report/test/mem.prof
```

- how to read the pprof report

[url](https://github.com/google/pprof/blob/main/doc/README.md#interpreting-the-callgraph)