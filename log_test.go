package golog

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestAppend(t *testing.T) {
	r := require.New(t)

	var buff bytes.Buffer
	logger := log.New(&buff, "", 0)

	l := New(logger)

	makeAppendFor := func(count int, f func(int) string) string {
		var s string
		for i := 0; i < count; i++ {
			s += f(i)
		}
		return s
	}

	uns := make([]UnLevelFunc, 10)
	for n := range uns {
		s := fmt.Sprintf("with%d", n)
		l.Print(s)
		r.Equal(
			makeAppendFor(n, func(i int) string {
				return fmt.Sprintf("prefix%d", i)
			})+s+"\n",
			buff.String(),
		)

		uns[n] = l.Append(fmt.Sprintf("prefix%d", n))

		buff.Reset()
	}

	for n := range uns {
		uns[len(uns)-1-n]()

		s := fmt.Sprintf("with%d", len(uns)-1-n)
		l.Print(s)
		r.Equal(
			makeAppendFor(len(uns)-1-n, func(i int) string {
				return fmt.Sprintf("prefix%d", i)
			})+s+"\n",
			buff.String(),
		)

		buff.Reset()
	}
}

func TestReplace(t *testing.T) {
	r := require.New(t)

	var buff bytes.Buffer
	logger := log.New(&buff, "", 0)

	l := New(logger)
	uns := make([]UnLevelFunc, 10)
	for n := range uns {
		prefix := fmt.Sprintf("prefix%d", n)
		s := fmt.Sprintf("with%d", n)

		uns[n] = l.Replace(prefix)

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

func TestFormat(t *testing.T) {
	r := require.New(t)

	var buff bytes.Buffer
	logger := log.New(&buff, "", 0)

	l := New(logger)

	format := "[%s]"

	makeFormatFor := func(count int, format string) string {
		var s string
		for i := 0; i < count; i++ {
			s = fmt.Sprintf(format, s)
		}
		return s
	}

	uns := make([]UnLevelFunc, 10)
	for n := range uns {
		s := fmt.Sprintf("with%d", n)
		prefix := makeFormatFor(n, format)

		l.Print(s)
		r.Equal(prefix+s+"\n", buff.String())

		uns[n] = l.Format(format)

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
