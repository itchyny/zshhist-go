package zshhist

import (
	"bytes"
	"os"
	"path"
	"testing"
)

func TestWriter(t *testing.T) {
	testCases := map[string]struct {
		file      string
		histories []History
	}{
		"extended": {
			file: "extended.histfile",
			histories: []History{
				{1537447916, 0, "echo hello, world!"},
				{1537448167, 1, "echo hello\necho world\n"},
				{1537448170, 0, "go version"},
				{1537448316, 0, "echo こんにちは、世界"},
				{1537448503, 0, "echo a\nb\nc\n"},
				{1537448453, 0, "echo \"あいうえお\nかきくけこ\nさしすせそ\""},
			},
		},
	}

	for name, tt := range testCases {
		t.Run(name, func(t *testing.T) {
			got := new(bytes.Buffer)
			w := NewWriter(got)

			for _, h := range tt.histories {
				err := w.Write(h)
				if err != nil {
					t.Fatalf("unable to write: %v", err.Error())
				}
			}

			want, err := os.ReadFile(path.Join("testdata", tt.file))
			if err != nil {
				t.Fatalf("unable to read testdata file: %v: %v", tt.file, err.Error())
			}

			if string(want) != got.String() {
				t.Errorf("\nwant:\n%v\ngot:\n%v\n", string(want), got.String())
			}
		})
	}
}
