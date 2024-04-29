package zshhist

import (
	"io"
	"strconv"
	"strings"
)

// A Writer writes history to a writer.
type Writer struct {
	out io.Writer
}

// NewWriter creates a new writer.
func NewWriter(w io.Writer) *Writer {
	return &Writer{out: w}
}

// Write a history to the writer.
func (w *Writer) Write(h History) (err error) {
	if h.Time != 0 {
		if err = w.extended(h); err != nil {
			return err
		}
	}

	start, idx := 0, strings.IndexRune(h.Command, '\n')
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
	_, err = w.out.Write([]byte{'\n'})
	return err
}

func (w *Writer) extended(h History) (err error) {
	_, err = w.out.Write([]byte{':', ' '})
	if err != nil {
		return err
	}
	_, err = w.out.Write([]byte(strconv.Itoa(h.Time)))
	if err != nil {
		return err
	}
	_, err = w.out.Write([]byte{':'})
	if err != nil {
		return err
	}
	_, err = w.out.Write([]byte(strconv.Itoa(h.Elapsed)))
	if err != nil {
		return err
	}
	_, err = w.out.Write([]byte{';'})
	if err != nil {
		return err
	}

	return nil
}
