package ccc

import "io"

func NewXORReader(r1, r2 io.Reader) *XORReader {
	return &XORReader{r1: r1, r2: r2}
}

type XORReader struct {
	r1, r2 io.Reader
	buf    []byte
}

func (r *XORReader) Read(p []byte) (n int, err error) {
	if len(r.buf) < len(p) {
		r.buf = make([]byte, len(p))
	}
	q := r.buf[:len(p)]

	n1, err1 := r.r1.Read(p)
	n2, err2 := r.r2.Read(q)

	n = n1
	if n2 < n1 {
		n = n2
	}

	for i := 0; i < n; i++ {
		p[i] ^= q[i]
	}

	err = err1
	if err == nil {
		err = err2
	}

	return
}

type FuncReader func() byte

func (f FuncReader) Read(p []byte) (n int, err error) {
	n = len(p)
	for i := range p {
		p[i] = f()
	}
	return
}

func NewLoopReader(r io.ReadSeeker) *LoopReader {
	return &LoopReader{r}
}

type LoopReader struct {
	r io.ReadSeeker
}

func (r *LoopReader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	if err == io.EOF {
		_, err = r.r.Seek(0, io.SeekStart)
	}
	return
}
