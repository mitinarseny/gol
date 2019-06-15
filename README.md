# golog [![Build Status](https://img.shields.io/travis/mitinarseny/golog/master.svg?style=flat-square)](https://travis-ci.com/mitinarseny/golog) [![Coverage](https://img.shields.io/codecov/c/github/mitinarseny/golog/master.svg?style=flat-square)](https://codecov.io/gh/mitinarseny/golog) [![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/mitinarseny/golog)

A small wrapper for standard Go [log](https://golang.org/pkg/log/) package with enhanced prefix system.

## Install
```bash
go get github.com/mitinarseny/golog
```

## Documentation
See documentation for this package on [GoDoc](https://godoc.org/github.com/mitinarseny/golog).

## Usage
It is recommended to use it with `defer` keyword:
```go
l := golog.New(log.New(os.Stdout, "", 0))
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
	"github.com/mitinarseny/golog"
)

func main() {
	l := golog.New(log.New(os.Stdout, "", 0))
	
	r1 := l.SetPersistentPrefix("golog ")
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
golog is logger for Go with enhanced prefixes
golog is cool because you can grow prefixes easily
golog is cool because you do not have to repeat yourself
golog is cool because a lot of reasnons
golog is very cool
that's it :)
```
