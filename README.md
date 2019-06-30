# gol [![Build Status](https://img.shields.io/travis/mitinarseny/gol/master.svg?style=flat-square&logo=travis-ci)](https://travis-ci.org/mitinarseny/gol) [![Coverage](https://img.shields.io/codecov/c/github/mitinarseny/gol/master.svg?style=flat-square&logo=codecov&logoColor=success)](https://codecov.io/gh/mitinarseny/gol) [![ColangCI](https://golangci.com/badges/github.com/mitinarseny/gol.svg)](https://golangci.com/r/github.com/mitinarseny/gol) [![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/mitinarseny/gol)

A small wrapper for standard [log](https://golang.org/pkg/log/) package in Go with enhanced prefixes.

## Install
```bash
go get github.com/mitinarseny/gol
```

## Documentation
See documentation for this package on [GoDoc](https://godoc.org/github.com/mitinarseny/gol).

## Usage
Create `gol.Logger` from `log.Logger`:
```go
l := gol.New(log.New(os.Stdout, "", 0))
```
Type `gol.Logger` has `*log.Logger` embedded in it, so you can use all functions from [log](https://golang.org/pkg/log/) package.

You can use following functions to change prefix of logger:
* `l.SetPersistentPrefix("prefix")`
* `l.SetPersistentPrefixf("%s pattern")`
* `l.SetPrefix("prefix")`
* `l.SetPrefixf("%s pattern")`

All these functions return `RestoreFunc` which can be used to restore previous prefix.
It is recommended to use it with `defer` keyword:
```go
func someFunc() {
	defer l.SetPrefixf("  %s")()
	...
}
```

## Example
Consider an example below:
```go
// main.go
import (
	"os"
	"log"
	"github.com/mitinarseny/gol"
)

func main() {
	l := gol.New(log.New(os.Stdout, "", 0))
	
	r1 := l.SetPersistentPrefix("gol ")
	l.Println("is logger for Go with enhanced prefixes")
	
	r2 := l.SetPrefix("is cool because ")
	l.Println("you can grow prefixes easily")
	
	r3 := l.SetPrefixf("%syou ")
	l.Println("do not have to repeat yourself")
	
	r3()
	l.Println("it is easy to restore previous prefixes")
	
	r2()
	l.Println("is very cool")
	
	r1()
	l.Println("that's it :)")
}
```
This code will produce following output:
```
gol is logger for Go with enhanced prefixes
gol is cool because you can grow prefixes easily
gol is cool because you do not have to repeat yourself
gol is cool because a lot of reasnons
gol is very cool
that's it :)
```
