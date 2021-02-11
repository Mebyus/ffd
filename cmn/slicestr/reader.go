package slicestr

import "io"

// A Reader implements the io.Reader
// The zero value for Reader will return EOF
// on the first call of Read
type Reader struct {
	ss []string // underlying slice
	i  int64    // reading index of current string
	j  int      // index of current string
}

// NewReader creates an instance of Reader which reads
// sequentially from underlying slice of strings.
// Pertains original slice
func NewReader(ss []string) *Reader {
	return &Reader{
		ss: ss,
	}
}

func (r *Reader) Read(b []byte) (n int, err error) {
	if r.j >= len(r.ss) {
		return 0, io.EOF
	}
	n = copy(b, r.ss[r.j][r.i:])
	r.i += int64(n)
	if r.i >= int64(len(r.ss[r.j])) {
		r.j++
		r.i = 0
	}
	return
}
