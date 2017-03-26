package flv

import (
	"io"
)

// Writer writes FLV header and tags to an output stream.
type Writer struct {
	*fileWriter
}

// NewWriter returns a new writer that writes to w.
func NewWriter(w io.Writer) *Writer {
	return &Writer{newFileWriter(w)}
}

// WriteHeader writes FLV header.
func (w *Writer) WriteHeader(h *Header) error {
	b := w.next(13)
	putUint24(b, signature)
	b[3] = 1
	b[4] = h.Flags
	putUint32(b[5:], 9)
	putUint32(b[9:], 0)
	return w.flush()
}

func (w *Writer) WriteTag(tag *Tag, r io.Reader) error {
	p := len(w.buf)
	b := w.next(11)
	b[0] = tag.Type
	putTime(b[4:], tag.Time)
	putUint24(b[8:], tag.Stream)
	n, err := w.fill(r)
	if err != nil {
		return err
	}
	putUint24(w.buf[p+1:], uint32(n))
	putUint32(w.next(4), uint32(n+11))
	return w.flush()
}

var bufferSize = 4096

type fileWriter struct {
	w   io.Writer
	buf []byte
}

func newFileWriter(w io.Writer) *fileWriter {
	return &fileWriter{w, nil}
}

func (w *fileWriter) next(n int) (v []byte) {
	v, w.buf = grow(w.buf, n)
	return
}

func (w *fileWriter) fill(r io.Reader) (int, error) {
	total := 0
	for {
		b := w.buf
		p := len(b)
		v := b[p:cap(b)]
		if len(v) < bufferSize {
			v, b = grow(b, bufferSize)
		}
		n, err := r.Read(v)
		total += n
		w.buf = b[:p+n]
		if err != nil {
			if err == io.EOF {
				break
			}
			return total, err
		}
	}
	return total, nil
}

func (w *fileWriter) flush() (err error) {
	if w.buf == nil {
		return nil
	}
	_, err = w.w.Write(w.buf)
	w.buf = w.buf[:0]
	return
}

func grow(b []byte, n int) (v, next []byte) {
	l := len(b)
	r := l + n
	if r > cap(b) {
		next := make([]byte, (1+((r-1)>>10))<<10)
		if l > 0 {
			copy(next, b[:l])
		}
		b = next
	}
	return b[l:r], b[:r]
}
