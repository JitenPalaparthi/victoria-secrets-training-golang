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


- To code coverage

```
go test ./... -cover
```
- To cover package

```
cd slices
go test -cover -coverpkg .
```

- Test cover profile

```
go test -cover -coverpkg ./... -coverprofile coverage.out
```

- Test cover profile in HTML

```
go tool cover -html=coverage.out -o coverage.html
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



- cover profile information

|covererage file|from line:column to line:colum|number of statements|count|
|---------------|------------------------------|--------------------|-----|
|demo/slices/slices.go|13.47,19.20|2|1|
|demo/slices/slices.go|19.20,21.3|1|1|


# Gomock

- Install a tool called mockgen

```
go install go.uber.org/mock/mockgen@latest
```

Step-1 : Install mockgen tool

Step-2: Which ever the package that is using thirdparty api or integrations, run mockgen tool

Step-3: run 

```
go generate ./...
```