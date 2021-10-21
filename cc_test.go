package cc_test

import (
	"bytes"
	"errors"
	"io"
	"testing"

	"github.com/gonutz/cc"
	"github.com/gonutz/check"
)

func TestXORReader(t *testing.T) {
	a := bytes.NewReader([]byte{12, 34, 56})
	b := bytes.NewReader([]byte{65, 43, 21})

	r := cc.NewXORReader(a, b)

	var xor [5]byte
	n, err := r.Read(xor[:])
	check.Eq(t, err, nil)
	check.Eq(t, n, 3)
	check.Eq(t, xor[:3], []byte{12 ^ 65, 34 ^ 43, 56 ^ 21})
}

func TestFirstReaderStopsEarly(t *testing.T) {
	a := bytes.NewReader([]byte{12})
	b := bytes.NewReader([]byte{65, 43, 21})

	r := cc.NewXORReader(a, b)

	var xor [5]byte
	n, err := r.Read(xor[:])
	check.Eq(t, err, nil)
	check.Eq(t, n, 1)
	check.Eq(t, xor[:1], []byte{12 ^ 65})
}

func TestSecondReaderStopsEarly(t *testing.T) {
	a := bytes.NewReader([]byte{12, 34, 56})
	b := bytes.NewReader([]byte{65})

	r := cc.NewXORReader(a, b)

	var xor [5]byte
	n, err := r.Read(xor[:])
	check.Eq(t, err, nil)
	check.Eq(t, n, 1)
	check.Eq(t, xor[:1], []byte{12 ^ 65})
}

func TestErrorOnFirstReaderBubblesUp(t *testing.T) {
	a := failAfterRead(bytes.NewReader([]byte{12, 34, 56}))
	b := bytes.NewReader([]byte{65, 43, 21})

	r := cc.NewXORReader(a, b)

	_, err := r.Read(make([]byte, 5))
	check.Eq(t, err.Error(), "fail")
}

func TestErrorOnSecondReaderBubblesUp(t *testing.T) {
	a := bytes.NewReader([]byte{12, 34, 56})
	b := failAfterRead(bytes.NewReader([]byte{65, 43, 21}))

	r := cc.NewXORReader(a, b)

	_, err := r.Read(make([]byte, 5))
	check.Eq(t, err.Error(), "fail")
}

func failAfterRead(r io.Reader) io.Reader {
	return &failer{r: r}
}

type failer struct {
	r io.Reader
}

func (r *failer) Read(p []byte) (n int, err error) {
	n, _ = r.r.Read(p)
	err = errors.New("fail")
	return
}

func TestFuncReader(t *testing.T) {
	f := func() byte { return 123 }
	var buf [3]byte
	n, err := cc.NewFuncReader(f).Read(buf[:])
	check.Eq(t, err, nil)
	check.Eq(t, n, 3)
	check.Eq(t, buf, [3]byte{123, 123, 123})
}
