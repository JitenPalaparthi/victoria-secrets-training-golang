- Create a directory . Example

 ```
  mkdir 1-helloworld
  cd 1-helloworld
  go mod init demo
```
- explanation: The name demo is , the name of the go module.
- Ultimately this name demo becomes the virtual root path of the project

- To run go application

```
go run main.go
```

- To see workdir

```
go run --work main.go
```

- To build a binary

```
go build main.go
```
- To create a binary with custom name
```  
go build -o hello main.go
```

- To cross compile 

    - get the list of available os/arch

    ```
    go tool dist list
    ```
    darwin/amd64
    darwin/arm64


    - Git commands

    ```
    git init // do it only once
    ```

    ```
    git add <files>
    ```

    ```
    git commit -m <message, make sure the message is in present simple>
    ```

    ```
    git push origin main
    ```
    ```
    git push -f origin main
    ```