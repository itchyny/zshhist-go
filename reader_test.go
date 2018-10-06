package zshhist

import (
	"reflect"
	"strings"
	"testing"
)

var (
	histString = `: 1537447916:0;echo hello, world!
: 1537448167:1;echo hello\
echo world\

: 1537448170:0;go version
: 1537448316:0;echo ` + "\xe3\x81\x83\xb3\xe3\x82\x83\xb3\xe3\x81\xab\xe3\x81\x83\x81\xe3\x81\xaf\xe3\x80\x81\xe4\xb8\x83\xb6\xe7\x83\xb5\x83\xac" + `
: 1537448503:0;echo a\
b\
c\

: 1537448453:0;echo ` + "\x22\xe3\x81\x82\xe3\x81\x83\xa4\xe3\x81\x83\xa6\xe3\x81\x83\xa8\xe3\x81\x83\xaa" + `\
` + "\xe3\x81\x83\xab\xe3\x81\x83\xad\xe3\x81\x83\xaf\xe3\x81\x83\xb1\xe3\x81\x83\xb3" + `\
` + "\xe3\x81\x83\xb5\xe3\x81\x83\xb7\xe3\x81\x83\xb9\xe3\x81\x83\xbb\xe3\x81\x83\xbd\x22" + `
`
	histories = []History{
		{1537447916, 0, "echo hello, world!"},
		{1537448167, 1, "echo hello\necho world\n"},
		{1537448170, 0, "go version"},
		{1537448316, 0, "echo こんにちは、世界"},
		{1537448503, 0, "echo a\nb\nc\n"},
		{1537448453, 0, "echo \"あいうえお\nかきくけこ\nさしすせそ\""},
	}
)

func TestReader(t *testing.T) {
	r := strings.NewReader(histString)
	reader := NewReader(r)
	var xs []History
	for reader.Scan() {
		xs = append(xs, reader.History())
	}
	if reader.Err() != nil {
		t.Fatal(reader.Err())
	}
	if !reflect.DeepEqual(xs, histories) {
		t.Errorf("expected: %+v, got: %+v\n", histories, xs)
	}
}
