If the main package is split in more than one file , then while compiling, building or running the application need to call all files.

```
go run main.go x.go y.go
go build main.go x.go y.go
go build -o funcs main.go x.go y.go
```