package zshhist

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// ReadAll reads the entire zsh histfile
func ReadAll(r io.Reader) (histories []History, err error) {
	var time, elapsed int
	var cmd string
	var cont bool
	s := bufio.NewScanner(r)
	for s.Scan() {
		line := Unmetafy(s.Text())
		if cont {
			if strings.HasSuffix(line, "\\") {
				line = line[:len(line)-1]
			} else {
				cont = false
			}
			cmd += "\n" + line
			if !cont {
				histories = append(histories, History{time, elapsed, cmd})
				cmd = ""
			}
			continue
		}
		if !strings.HasPrefix(line, ": ") {
			return nil, fmt.Errorf("invalid histfile line: %q", line)
		}
		i := strings.IndexRune(line[2:], ':')
		if i < 0 {
			return nil, fmt.Errorf("invalid histfile line: %q", line)
		}
		time, err = strconv.Atoi(line[2 : i+2])
		if err != nil {
			return nil, fmt.Errorf("invalid histfile line: %q", line)
		}
		j := strings.IndexRune(line[2:], ';')
		if j < 0 {
			return nil, fmt.Errorf("invalid histfile line: %q", line)
		}
		elapsed, err = strconv.Atoi(line[i+3 : j+2])
		if err != nil {
			return nil, fmt.Errorf("invalid histfile line: %q", line)
		}
		cmd = line[j+3:]
		if strings.HasSuffix(cmd, "\\") {
			cont = true
			cmd = cmd[:len(cmd)-1]
		} else {
			histories = append(histories, History{time, elapsed, cmd})
			cmd = ""
		}
	}
	return histories, nil
}
