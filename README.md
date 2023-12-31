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

    - First time pushing to github? Then configure your username and email

    ```
       git config --global user.name Jiten
       git config --global user.email JitenP@Outlook.Com
    ```

- To clone git repo
    ```
        git clone https://github.com/JitenPalaparthi/victoria-secrets-training-golang
    ```


- To create a library 
    - Create a repo in github.com
    - clone the repo in your local system


# Presentation links

- Threads in Go : https://docs.google.com/presentation/d/1t8ERWC5v75OKko_5Riw8wlwZ-5Dht3K_0uMXBc80zfc/edit?usp=sharing

-Golang Presentation:  https://docs.google.com/presentation/d/1WVvsbvgHKBrNrKtnT4XWRfrsfkNlbw5u6L9O1DeVBn0/edit?usp=sharing



- Topics Covered : 

https://1drv.ms/x/s!Aur57CAI2lNYwmIxh7_fnIk5jEma?e=eElzdV