package golog

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
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

		fmt.Print(buff.String())
		buff.Reset()
	}

	for n := range uns {
		prefix := fmt.Sprintf("prefix%d", len(uns)-1-n)
		s := fmt.Sprintf("with%d", len(uns)-1-n)

		l.Print(s)
		r.Equal(prefix+s+"\n", buff.String())

		uns[len(uns)-1-n]()

		fmt.Print(buff.String())
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

		fmt.Print(buff.String())
		buff.Reset()
	}
	for n := range uns {
		uns[len(uns)-1-n]()

		s := fmt.Sprintf("with%d", len(uns)-1-n)
		prefix := makeFormatFor(len(uns)-1-n, format)

		l.Print(s)
		r.Equal(prefix+s+"\n", buff.String())

		fmt.Print(buff.String())
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

		fmt.Print(buff.String())
		buff.Reset()
	}

	for n := range uns {
		prefix := fmt.Sprintf("prefix%d", len(uns)-1-n)
		s := fmt.Sprintf("with%d", len(uns)-1-n)

		l.Print(s)
		r.Equal(prefix+s+"\n", buff.String())

		uns[len(uns)-1-n]()

		fmt.Print(buff.String())
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

		fmt.Print(buff.String())
		buff.Reset()
	}
	for n := range uns {
		uns[len(uns)-1-n]()

		s := fmt.Sprintf("with%d", len(uns)-1-n)
		prefix := makeFormatFor(len(uns)-1-n, format)

		l.Print(s)
		r.Equal(prefix+s+"\n", buff.String())

		fmt.Print(buff.String())
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
	fmt.Print(buff.String())
	buff.Reset()

	restore := l.SetPrefix(prefix)
	withBoth := "with both"
	l.Print(withBoth)
	r.Equal(persistentPrefix+prefix+withBoth+"\n", buff.String())
	fmt.Print(buff.String())
	buff.Reset()

	restore()
	l.Print(withPersistent)
	r.Equal(persistentPrefix+withPersistent+"\n", buff.String())
	fmt.Print(buff.String())
	buff.Reset()

	restorePersistent()
	withoutAny := "without any"
	l.Print(withoutAny)
	r.Equal(withoutAny+"\n", buff.String())
	fmt.Print(buff.String())
	buff.Reset()
}
