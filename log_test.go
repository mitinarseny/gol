package golog_test

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
	"os"
)

func TestSetPrefix(t *testing.T) {
	r := require.New(t)

	var buff bytes.Buffer
	l := New(log.New(&buff, "", 0))

	uns := make([]RestoreFunc, 10)
	for n := range uns {
		prefix := fmt.Sprintf("prefix%d", n)
		s := fmt.Sprintf("with%d", n)

		uns[n] = l.SetPrefix(prefix)

		l.Print(s)
		r.Equal(prefix+s+"\n", buff.String())

		buff.Reset()
	}

	for n := range uns {
		prefix := fmt.Sprintf("prefix%d", len(uns)-1-n)
		s := fmt.Sprintf("with%d", len(uns)-1-n)

		l.Print(s)
		r.Equal(prefix+s+"\n", buff.String())

		uns[len(uns)-1-n]()

		buff.Reset()
	}
}

func TestSetPrefixf(t *testing.T) {
	r := require.New(t)

	var buff bytes.Buffer
	l := New(log.New(&buff, "", 0))

	format := "[%s]"

	makeFormatFor := func(count int, format string) string {
		var s string
		for i := 0; i < count; i++ {
			s = fmt.Sprintf(format, s)
		}
		return s
	}

	uns := make([]RestoreFunc, 10)
	for n := range uns {
		s := fmt.Sprintf("with%d", n)
		prefix := makeFormatFor(n, format)

		l.Print(s)
		r.Equal(prefix+s+"\n", buff.String())

		uns[n] = l.SetPrefixf(format)

		buff.Reset()
	}
	for n := range uns {
		uns[len(uns)-1-n]()

		s := fmt.Sprintf("with%d", len(uns)-1-n)
		prefix := makeFormatFor(len(uns)-1-n, format)

		l.Print(s)
		r.Equal(prefix+s+"\n", buff.String())

		buff.Reset()
	}
}

func TestInitialPersistentPrefix(t *testing.T) {
	r := require.New(t)

	var buff bytes.Buffer
	prefix := "initial prefix"
	l := New(log.New(&buff, prefix, 0))

	s := "test"
	l.Print(s)
	r.Equal(prefix+s+"\n", buff.String())
}

func TestSetPersistentPrefix(t *testing.T) {
	r := require.New(t)

	var buff bytes.Buffer
	l := New(log.New(&buff, "", 0))

	uns := make([]RestoreFunc, 10)
	for n := range uns {
		prefix := fmt.Sprintf("prefix%d", n)
		s := fmt.Sprintf("with%d", n)

		uns[n] = l.SetPersistentPrefix(prefix)

		l.Print(s)
		r.Equal(prefix+s+"\n", buff.String())

		buff.Reset()
	}

	for n := range uns {
		prefix := fmt.Sprintf("prefix%d", len(uns)-1-n)
		s := fmt.Sprintf("with%d", len(uns)-1-n)

		l.Print(s)
		r.Equal(prefix+s+"\n", buff.String())

		uns[len(uns)-1-n]()

		buff.Reset()
	}
}

func TestSetPersistentPrefixf(t *testing.T) {
	r := require.New(t)

	var buff bytes.Buffer
	l := New(log.New(&buff, "", 0))

	format := "[%s]"

	makeFormatFor := func(count int, format string) string {
		var s string
		for i := 0; i < count; i++ {
			s = fmt.Sprintf(format, s)
		}
		return s
	}

	uns := make([]RestoreFunc, 10)
	for n := range uns {
		s := fmt.Sprintf("with%d", n)
		prefix := makeFormatFor(n, format)

		l.Print(s)
		r.Equal(prefix+s+"\n", buff.String())

		uns[n] = l.SetPersistentPrefixf(format)

		buff.Reset()
	}
	for n := range uns {
		uns[len(uns)-1-n]()

		s := fmt.Sprintf("with%d", len(uns)-1-n)
		prefix := makeFormatFor(len(uns)-1-n, format)

		l.Print(s)
		r.Equal(prefix+s+"\n", buff.String())

		buff.Reset()
	}
}

func TestMixedPrefix(t *testing.T) {
	r := require.New(t)

	var buff bytes.Buffer
	l := New(log.New(&buff, "", 0))
	persistentPrefix := "persistentPrefix"
	prefix := "prefix"

	restorePersistent := l.SetPersistentPrefix(persistentPrefix)
	withPersistent := "with persistent"
	l.Print(withPersistent)
	r.Equal(persistentPrefix+withPersistent+"\n", buff.String())
	buff.Reset()

	restore := l.SetPrefix(prefix)
	withBoth := "with both"
	l.Print(withBoth)
	r.Equal(persistentPrefix+prefix+withBoth+"\n", buff.String())
	buff.Reset()

	restore()
	l.Print(withPersistent)
	r.Equal(persistentPrefix+withPersistent+"\n", buff.String())
	buff.Reset()

	restorePersistent()
	withoutAny := "without any"
	l.Print(withoutAny)
	r.Equal(withoutAny+"\n", buff.String())
	buff.Reset()
}

func ExampleLogger_SetPersistentPrefix() {
	l := New(log.New(os.Stdout, "", 0))

	reverse := l.SetPersistentPrefix("golog ")

	l.Println("has persistent prefixes")
	reverse()
	
	l.Println("that can be easily reversed")
}

func ExampleLogger_SetPrefix(){
	l := New(log.New(os.Stdout, "", 0))

	reverse := l.SetPrefix("golog ")

	l.Println("has prefixes")
	reverse()
	
	l.Println("that can be easily reversed")
}

func Example(){
	l := New(log.New(os.Stdout, "", 0))
	
	reversePersistent := l.SetPersistentPrefix("golog ")
	l.Println("is logger for Go with enhanced prefixes")
	
	reverse1 := l.SetPrefix("is cool because ")
	l.Println("you can grow prefixes easily")
	
	reverse2 := l.SetPrefixf("%syou ")
	l.Println("do not have to repeat yourself")
	
	reverse2()
	l.Println("it is easy to restore previous prefixes")
	
	reverse1()
	l.Println("is very cool")
	
	reversePersistent()
	l.Println("that's it :)")
}

func Example_deferred(){
	l := New(log.New(os.Stdout, "", 0))

	func(){
		defer l.SetPrefix("reverse functions ")()

		func(){
			defer l.SetPrefixf("%sare especially convenient ")()

			l.Println("when each function is associated with its own prefix, so")
		}()

		l.Println("have no need to be cared about")
	}()

	l.Println("they are executed automatically!")
}
