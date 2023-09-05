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

# Task 

- Create a package called slice 

- Fill the slice with data 

- Write a method to find the reverse of the slice 

 - Write a method to find the smallest number in the slice

 - Write a method to fund the biggest number in the slice

 - write a method to find the sum of all elements in the slice


 - Write unit tests for all of them

 