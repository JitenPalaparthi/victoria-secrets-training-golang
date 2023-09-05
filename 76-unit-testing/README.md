# Unit Testing

- TDD: Test Driven Development : First write tests and make them pass
- TDD: Table Driven Development

- Testing file should append with _test.go

- Each Unit test function should start with Test

- To Test a functionality , use the tool go test


- To test a specific file 

```
go test -timeout 30s -run ^(TestReverseString|TestSrtingLength)$ demo/strings
```

- To test a specifc Unit Test

- verbost test

```
go test demo/strings -v
```

```
go test -timeout 30s -run ^TestReverseString$ demo/strings
```

- To test full package 

```
go test demo/strings
```
- Benchmarking 

```
go test -benchmem -run=^$ -bench ^BenchmarkSecondBiggest$ demo/slices
```

- time based bench mark

```
go test -bench=. -benchtime=10s
```

- count based  benchmark

```
go test -bench=. -count=10
```

- benchtest memory

```
go test -bench=. -benchtime=10s -benchmem
```
|Benchmark test name| Number of Loop Iterations| NanoSecond per operation| Bytes per Operations| Allocations for Operation|
|-------------------|--------------------------|------------------------|-----------------------|-----------------|
|BenchmarkSecondBiggest-12|        1851162|              6073 ns/op|            2040 B/op|          8 allocs/op|
|BenchmarkSecondBiggest2-12|      14312559|               738.9 ns/op|             0 B/op|          0 allocs/op|

# Task 

- Create a package called slice 

- Fill the slice with data.

- Write a method to find the reverse of the slice 

- Write a method to find the smallest number in the slice

- Write a method to fund the biggest number in the slice

- write a method to find the sum of all elements in the slice

- Write unit tests for all of them
