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

// Scan the reader and reports whether it successfully parses one history or not.
func (d *Reader) Scan() bool {
	var time, elapsed int
	var cmd string
	var cont bool
	var err error
	for d.s.Scan() {
		line := Unmetafy(d.s.Text())
		if cont {
			if strings.HasSuffix(line, "\\") {
				line = line[:len(line)-1]
			} else {
				cont = false
			}
			cmd += "\n" + line
			if !cont {
				d.history, d.err = History{time, elapsed, cmd}, nil
				return true
			}
			continue
		}

		if strings.HasPrefix(line, ": ") {
			time, elapsed, cmd, err = extended(line)
			if err != nil {
				d.err = err
				return false
			}
		} else {
			cmd = line
		}

		if strings.HasSuffix(cmd, "\\") {
			cont = true
			cmd = cmd[:len(cmd)-1]
		} else {
			d.history, d.err = History{time, elapsed, cmd}, nil
			return true
		}
	}
	return false
}

// History returns the lastly parsed history.
func (d *Reader) History() History {
	return d.history
}

// Err returns the parse error.
func (d *Reader) Err() error {
	return d.err
}

func extended(line string) (time, elapsed int, cmd string, err error) {
	i := strings.IndexRune(line[2:], ':')
	if i < 0 {
		return time, elapsed, cmd, fmt.Errorf("invalid histfile line: %q", line)
	}

	time, err = strconv.Atoi(line[2 : i+2])
	if err != nil {
		return time, elapsed, cmd, fmt.Errorf("invalid histfile line: %q", line)
	}

	j := strings.IndexRune(line[2:], ';')
	if j < 0 {
		return time, elapsed, cmd, fmt.Errorf("invalid histfile line: %q", line)
	}

	elapsed, err = strconv.Atoi(line[i+3 : j+2])
	if err != nil {
		return time, elapsed, cmd, fmt.Errorf("invalid histfile line: %q", line)
	}

	cmd = line[j+3:]

	return time, elapsed, cmd, nil
}
