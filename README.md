# Learn golang
repo is made for training

## Golang
Go is an open source programming language by Google
* compile and static type langugae
* support garbage collector and concurrency

Official...
* Go mirror repo: https://github.com/golang/go
* Tutorial: https://go.dev/doc/tutorial/getting-started

## Install for MacOS
```
$ brew install go
```

previous versions example
```
$ brew install go@1.18
```

installed go lang versions
```
$ brew list | grep go
go
go@1.18
```

## Init
Init module and create directories
```
$ go mod init github.com/morenice/learn-language-go
$ mkdir -p lib
$ mkdir -p test
```

## Build
```
$ go fmt       # auto formating
$ go build     # source code build
```

## Module
Search go module in ```https://pkg.go.dev```

example
```
$ go get github.com/gogo/protobuf/proto

$ cat go.mod
module github.com/morenice/learn-language-go

go 1.19

require github.com/gogo/protobuf v1.3.2 // indirect
```

## Testing
```
$ go test
```
