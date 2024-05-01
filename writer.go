package zshhist

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

// A Writer writes history to a writer.
type Writer struct {
	w   *bufio.Writer
	buf [24]byte
}

// NewWriter creates a new writer.
func NewWriter(w io.Writer) *Writer {
	return &Writer{w: bufio.NewWriter(w)}
}

// Write a history to the writer.
func (w *Writer) Write(h History) (err error) {
	if h.Time != 0 {
		if err = w.writeExtended(h); err != nil {
			return
		}
	}

	for cmd := h.Command; ; {
		i := strings.IndexByte(cmd, '\n')
		if i < 0 {
			_, err = w.w.WriteString(Metafy(cmd))
			if err != nil {
				return
			}
			break
		}
		_, err = w.w.WriteString(Metafy(cmd[:i]))
		if err != nil {
			return
		}
		_, err = w.w.WriteString("\\\n")
		if err != nil {
			return
		}
		cmd = cmd[i+1:]
	}

	err = w.w.WriteByte('\n')
	if err != nil {
		return
	}
	err = w.w.Flush()
	return
}

func (w *Writer) writeExtended(h History) (err error) {
	buf := append(w.buf[:0], ':', ' ')
	buf = append(strconv.AppendInt(buf, int64(h.Time), 10), ':')
	buf = append(strconv.AppendInt(buf, int64(h.Elapsed), 10), ';')
	_, err = w.w.Write(buf)
	return
}
