package ccc_test

import (
	"bytes"
	"errors"
	"io"
	"testing"

	"github.com/gonutz/ccc"
	"github.com/gonutz/check"
)

func TestXORReader(t *testing.T) {
	a := bytes.NewReader([]byte{12, 34, 56})
	b := bytes.NewReader([]byte{65, 43, 21})

	r := ccc.NewXORReader(a, b)

	var xor [5]byte
	n, err := r.Read(xor[:])
	check.Eq(t, err, nil)
	check.Eq(t, n, 3)
	check.Eq(t, xor[:3], []byte{12 ^ 65, 34 ^ 43, 56 ^ 21})
}

func TestFirstReaderStopsEarly(t *testing.T) {
	a := bytes.NewReader([]byte{12})
	b := bytes.NewReader([]byte{65, 43, 21})

	r := ccc.NewXORReader(a, b)

	var xor [5]byte
	n, err := r.Read(xor[:])
	check.Eq(t, err, nil)
	check.Eq(t, n, 1)
	check.Eq(t, xor[:1], []byte{12 ^ 65})
}

func TestSecondReaderStopsEarly(t *testing.T) {
	a := bytes.NewReader([]byte{12, 34, 56})
	b := bytes.NewReader([]byte{65})

	r := ccc.NewXORReader(a, b)

	var xor [5]byte
	n, err := r.Read(xor[:])
	check.Eq(t, err, nil)
	check.Eq(t, n, 1)
	check.Eq(t, xor[:1], []byte{12 ^ 65})
}

func TestErrorOnFirstReaderBubblesUp(t *testing.T) {
	a := failAfterRead(bytes.NewReader([]byte{12, 34, 56}))
	b := bytes.NewReader([]byte{65, 43, 21})

	r := ccc.NewXORReader(a, b)

	_, err := r.Read(make([]byte, 5))
	check.Eq(t, err.Error(), "fail")
}

func TestErrorOnSecondReaderBubblesUp(t *testing.T) {
	a := bytes.NewReader([]byte{12, 34, 56})
	b := failAfterRead(bytes.NewReader([]byte{65, 43, 21}))

	r := ccc.NewXORReader(a, b)

	_, err := r.Read(make([]byte, 5))
	check.Eq(t, err.Error(), "fail")
}

func TestFuncReader(t *testing.T) {
	f := func() byte { return 123 }
	buf := make([]byte, 3)
	n, err := ccc.FuncReader(f).Read(buf)
	check.Eq(t, err, nil)
	check.Eq(t, n, 3)
	check.Eq(t, buf, []byte{123, 123, 123})
}

func TestLoopReader(t *testing.T) {
	r := bytes.NewReader([]byte{1, 2, 3})
	buf := make([]byte, 7)
	n, err := io.ReadFull(ccc.NewLoopReader(r), buf)
	check.Eq(t, err, nil)
	check.Eq(t, n, 7)
	check.Eq(t, buf, []byte{1, 2, 3, 1, 2, 3, 1})
}

func TestLoopReaderCanFail(t *testing.T) {
	r := failAfterRead(bytes.NewReader([]byte{1, 2, 3}))
	buf := make([]byte, 7)
	n, err := io.ReadFull(ccc.NewLoopReader(r), buf)
	check.Eq(t, err.Error(), "fail")
	check.Eq(t, n, 3)
	check.Eq(t, buf[:3], []byte{1, 2, 3})
}

func failAfterRead(r *bytes.Reader) io.ReadSeeker {
	return &failer{r}
}

type failer struct {
	*bytes.Reader
}

func (r *failer) Read(p []byte) (n int, err error) {
	n, _ = r.Reader.Read(p)
	err = errors.New("fail")
	return
}
