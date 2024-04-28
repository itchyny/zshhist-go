package zshhist

import (
	"os"
	"path"
	"reflect"
	"testing"
)

func TestReader(t *testing.T) {
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
			f, err := os.Open(path.Join("testdata", tt.file))
			if err != nil {
				t.Fatalf("unable to open testdata file: %v: %v", tt.file, err.Error())
			}
			defer f.Close()

			var got []History

			r := NewReader(f)
			for r.Scan() {
				got = append(got, r.History())
			}
			if r.Err() != nil {
				t.Fatalf("unable to scan testdata: %v", r.Err())
			}

			if !reflect.DeepEqual(tt.histories, got) {
				t.Errorf("\nwant:\n%+v\ngot:\n%+v\n", tt.histories, got)
			}
		})
	}
}
