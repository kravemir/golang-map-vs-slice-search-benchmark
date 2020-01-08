# golang-map-vs-slice-search-benchmark

Benchmark tests two implementations of indexes with string keys, and int64:

* sorted `[]string` with keys, and companion `[]int64` with values,
* native `map[string]int64`.

## Results

On my machine:

```
$ go version
go version go1.13.5 linux/amd64

$ go test -bench=. -benchtime 100000x
goos: linux
goarch: amd64
pkg: golang-map-vs-slice-search-benchmark
BenchmarkFind/IntIndex_Capacity_nil-4         	  100000	      3508 ns/op
BenchmarkFind/IntIndex_Capacity_100-4         	  100000	      3566 ns/op
BenchmarkFind/MapIndex-4                      	  100000	      1131 ns/op
PASS
ok  	golang-map-vs-slice-search-benchmark	0.827s
```

## Conclusion

Native golang's map wins (using given test data).
