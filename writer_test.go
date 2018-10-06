package zshhist

import (
	"bytes"
	"testing"
)

func TestWriter(t *testing.T) {
	out := new(bytes.Buffer)
	writer := NewWriter(out)
	for _, h := range histories {
		err := writer.Write(h)
		if err != nil {
			t.Fatal(err)
		}
	}
	if histString != out.String() {
		t.Errorf("expected: %+v, got: %+v\n%+v\n%+v", histString, out.String(), []byte(histString), []byte(out.String()))
	}
}
