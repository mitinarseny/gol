# golog
Logger for Go with enhanced prefixes

## Install
```bash
go get github.com/mitinarseny/golog
```

## Usage
Consider an example below:
```go
// main.go
package main

import (
	"os"
	"log"
	"github.com/mitinarseny/golog"
)

func main() {
	l := golog.New(log.New(os.Stdout, "", 0))
	
	restorePersistent := l.SetPersistentPrefix("golog ")
	l.Println("is logger for Go with enhanced prefixes")
	
	restore := l.SetPrefix("is cool because ")
	l.Println("reason 1")
	l.Println("reason 2")
	l.Println("reason 3")
	
	restore()
	l.Println("is very cool")
	restorePersistent()
	l.Println("that's it :)")
}
```
This code will produce following output:
```
golog is logger for Go with enhanced prefixes
golog is cool because reason1
golog is cool because reason2
golog is cool because reason3
golog is very cool
that's it :)
```
