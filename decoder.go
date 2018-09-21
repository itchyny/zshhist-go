package zshhist

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// A Decoder scans histfile and decodes each history.
type Decoder struct {
	s       *bufio.Scanner
	history History
	err     error
}

// NewDecoder creates a new decoder.
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{s: bufio.NewScanner(r)}
}

// Scan the reader and reports whether it successfully parses one history or not.
func (d *Decoder) Scan() bool {
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
		if !strings.HasPrefix(line, ": ") {
			d.err = fmt.Errorf("invalid histfile line: %q", line)
			return false
		}
		i := strings.IndexRune(line[2:], ':')
		if i < 0 {
			d.err = fmt.Errorf("invalid histfile line: %q", line)
			return false
		}
		time, err = strconv.Atoi(line[2 : i+2])
		if err != nil {
			d.err = fmt.Errorf("invalid histfile line: %q", line)
			return false
		}
		j := strings.IndexRune(line[2:], ';')
		if j < 0 {
			d.err = fmt.Errorf("invalid histfile line: %q", line)
			return false
		}
		elapsed, err = strconv.Atoi(line[i+3 : j+2])
		if err != nil {
			d.err = fmt.Errorf("invalid histfile line: %q", line)
			return false
		}
		cmd = line[j+3:]
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
func (d *Decoder) History() History {
	return d.history
}

// Err returns the parse error.
func (d *Decoder) Err() error {
	return d.err
}
