# Step 1. Initialize

## Install Go

https://golang.org/dl/

## Init a project

```
mkdir realworld
cd realworld
go mod init example.com/realworld
```

## Create the main.go

**cmd/main.go**

```go
package main

import "fmt"

func main() {
	fmt.Println("RealWorld back-end.")
}
```

## Install HTTP client

```shell script
go get -u github.com/astaxie/bat
```


## Run

```shell script
go run cmd/main.go
```

## Get a result

```shell script
go run github.com/astaxie/bat GET 127.0.0.1:3333
# or
# bat GET 127.0.0.1:3333
```