package zshhist

import (
	"io"
	"strconv"
	"strings"
)

type Writer struct {
	out io.Writer
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{out: w}
}

func (w *Writer) Write(h History) (err error) {
	_, err = w.out.Write([]byte{':', ' '})
	if err != nil {
		return
	}
	_, err = w.out.Write([]byte(strconv.Itoa(h.Time)))
	if err != nil {
		return
	}
	_, err = w.out.Write([]byte{':'})
	if err != nil {
		return
	}
	_, err = w.out.Write([]byte(strconv.Itoa(h.Elapsed)))
	if err != nil {
		return
	}
	_, err = w.out.Write([]byte{';'})
	if err != nil {
		return
	}
	idx := strings.IndexRune(h.Command, '\n')
	if idx < 0 {
		_, err = w.out.Write([]byte(Metafy(h.Command)))
		if err != nil {
			return
		}
	} else {
		start := 0
		for {
			if idx < 0 {
				_, err = w.out.Write([]byte(Metafy(h.Command[start:])))
				if err != nil {
					return
				}
				break
			}
			_, err = w.out.Write([]byte(Metafy(h.Command[start : start+idx])))
			if err != nil {
				return
			}
			_, err = w.out.Write([]byte{'\\', '\n'})
			if err != nil {
				return
			}
			start, idx = start+idx+1, strings.IndexRune(h.Command[start+idx+1:], '\n')
		}
	}
	if err != nil {
		return
	}
	_, err = w.out.Write([]byte{'\n'})
	if err != nil {
		return
	}
	return nil
}
