package zshhist

import (
	"reflect"
	"strings"
	"testing"
)

func TestReadAll(t *testing.T) {
	r := strings.NewReader(`: 1537447916:0;echo hello, world!
: 1537448167:1;echo hello\
echo world\

: 1537448170:0;go version
: 1537448316:0;echo ` + "\xe3\x81\x83\xb3\xe3\x82\x83\xb3\xe3\x81\xab\xe3\x81\x83\x81\xe3\x81\xaf\xe3\x80\x81\xe4\xb8\x83\xb6\xe7\x83\xb5\x83\xac" + `
`)
	xs, err := ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
	expected := []History{
		{1537447916, 0, "echo hello, world!"},
		{1537448167, 1, "echo hello\necho world\n"},
		{1537448170, 0, "go version"},
		{1537448316, 0, "echo こんにちは、世界"},
	}
	if !reflect.DeepEqual(xs, expected) {
		t.Errorf("expected: %+v, got: %+v\n", expected, xs)
	}
}
