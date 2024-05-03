package zshhist

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// A Reader scans histfile and decodes each history.
type Reader struct {
	s       *bufio.Scanner
	history History
	err     error
}

// NewReader creates a new reader.
func NewReader(r io.Reader) *Reader {
	return &Reader{s: bufio.NewScanner(r)}
}

// ErrInvalidLine is the error on invalid histfile line.
type ErrInvalidLine string

func (e ErrInvalidLine) Error() string {
	return fmt.Sprintf("zshhist: invalid histfile line: %q", string(e))
}

// Scan the reader and reports whether it successfully parses one history or not.
func (r *Reader) Scan() bool {
	var time, elapsed int64
	var cmd string
	var cont, extended bool
	var err error

	for r.s.Scan() {
		line := Unmetafy(r.s.Text())
		if cont {
			cmd += "\n" + line
		} else if cmd, extended = strings.CutPrefix(line, ": "); extended {
			time, elapsed, cmd, err = parseExtended(cmd)
			if err != nil {
				r.err = ErrInvalidLine(line)
				return false
			}
		}
		if cmd, cont = strings.CutSuffix(cmd, "\\"); !cont {
			r.history, r.err = History{time, elapsed, cmd}, nil
			return true
		}
	}

	return false
}

// History returns the lastly parsed history.
func (r *Reader) History() History {
	return r.history
}

// Err returns the parse error.
func (r *Reader) Err() error {
	return r.err
}

func parseExtended(line string) (time, elapsed int64, cmd string, err error) {
	var i, j int
	if i = strings.IndexByte(line, ':'); i < 0 {
		err = ErrInvalidLine(line)
		return
	}
	if time, err = strconv.ParseInt(line[:i], 10, 64); err != nil {
		err = ErrInvalidLine(line)
		return
	}
	if j = strings.IndexByte(line[i:], ';'); j < 0 {
		err = ErrInvalidLine(line)
		return
	}
	if elapsed, err = strconv.ParseInt(line[i+1:i+j], 10, 64); err != nil {
		err = ErrInvalidLine(line)
		return
	}
	cmd = line[i+j+1:]
	return
}
